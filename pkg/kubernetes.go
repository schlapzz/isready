package pkg

import (
	"context"
	"fmt"
	appsV1 "k8s.io/api/apps/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"time"
)

type KubernetesConfig struct {
	KubeConfig string
	Name       string
	Namespace  string
	Timeout    time.Duration
}

const (
	RevisionAnnotation = "deployment.kubernetes.io/revision"
	// TimedOutReason is added in a deployment when its newest replica set fails to show any progress
	// within the given deadline (progressDeadlineSeconds).
	TimedOutReason = "ProgressDeadlineExceeded"
)

func IsReady(config KubernetesConfig) error {
	t := time.Now()
	ctx, _ := context.WithDeadline(context.TODO(), t.Add(config.Timeout))

	// use the current context in kubeconfig
	rest, err := clientcmd.BuildConfigFromFlags("", config.KubeConfig)
	if err != nil {
		return err
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(rest)
	if err != nil {
		return err
	}

	watch, err := clientset.AppsV1().Deployments(config.Namespace).Watch(ctx, v1.ListOptions{})
	if err != nil {
		return err
	}

	for {
		select {
		case event := <-watch.ResultChan():
			p, ok := event.Object.(*appsV1.Deployment)
			if !ok {
				continue
			}

			if p.Name != config.Name {
				continue
			}

			if p.Status.ReadyReplicas == p.Status.Replicas {
				fmt.Println("deployment ready")
				return nil
			}

			fmt.Println()
		case <-ctx.Done():
			return fmt.Errorf("deadline exceeded")
		}
	}

}

func isDeploymentReady(deployment *appsV1.Deployment) (bool, error) {

	if deployment.Generation <= deployment.Status.ObservedGeneration {
		cond := GetDeploymentCondition(deployment.Status, appsV1.DeploymentProgressing)
		if cond != nil && cond.Reason == TimedOutReason {
			return false, fmt.Errorf("deployment %q exceeded its progress deadline", deployment.Name)
		}
		if deployment.Spec.Replicas != nil && deployment.Status.UpdatedReplicas < *deployment.Spec.Replicas {
			return false, fmt.Errorf("Waiting for deployment %q rollout to finish: %d out of %d new replicas have been updated...\n", deployment.Name, deployment.Status.UpdatedReplicas, *deployment.Spec.Replicas)
		}
		if deployment.Status.Replicas > deployment.Status.UpdatedReplicas {
			return false, fmt.Errorf("Waiting for deployment %q rollout to finish: %d old replicas are pending termination...\n", deployment.Name, deployment.Status.Replicas-deployment.Status.UpdatedReplicas)
		}
		if deployment.Status.AvailableReplicas < deployment.Status.UpdatedReplicas {
			return false, fmt.Errorf("Waiting for deployment %q rollout to finish: %d of %d updated replicas are available...\n", deployment.Name, deployment.Status.AvailableReplicas, deployment.Status.UpdatedReplicas)
		}
		return true, nil
	}

	return false, fmt.Errorf("Waiting for deployment spec update to be observed...\n")
}

// GetDeploymentCondition returns the condition with the provided type.
func GetDeploymentCondition(status appsV1.DeploymentStatus, condType appsV1.DeploymentConditionType) *appsV1.DeploymentCondition {
	for i := range status.Conditions {
		c := status.Conditions[i]
		if c.Type == condType {
			return &c
		}
	}
	return nil
}
