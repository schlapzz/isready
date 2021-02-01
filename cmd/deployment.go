/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"isready/pkg"
	"os"

	"github.com/spf13/cobra"
)

// deploymentCmd represents the deployment command
var deploymentCmd = &cobra.Command{
	Use:   "deployment",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("deployment called")

		name, _ := cmd.Flags().GetString("name")
		namespace, _ := cmd.Flags().GetString("namespace")
		kubeconfig, _ := cmd.Flags().GetString("kubeconfig")

		err := pkg.IsReady(
			pkg.KubernetesConfig{
				KubeConfig: kubeconfig,
				Name:       name,
				Namespace:  namespace,
				Timeout:    timeout,
			})

		if err != nil {
			os.Stderr.WriteString("deployment error: " + err.Error())
			os.Exit(6)
		}

	},
}

func init() {
	rootCmd.AddCommand(deploymentCmd)

	deploymentCmd.Flags().StringP("namespace", "s", "default", "namespace of deployment resource")
	deploymentCmd.Flags().StringP("name", "n", "", "name of deployment resource")
	deploymentCmd.Flags().StringP("kubeconfig", "c", "~/.kube/config", "path to kubeconfig file")

}
