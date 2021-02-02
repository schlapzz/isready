package pkg

import (
	"context"
	"fmt"
	"time"
)

type Connecter interface {
	Connect() error
}

var Timeout string

func ConnectWithRetries(c Connecter, retries int, interval, timeout time.Duration) error {
	t := time.Now()
	deadline := t.Add(timeout)
	ctx, _ := context.WithDeadline(context.TODO(), deadline)
	result := make(chan error)

	go func() {

		n := 0
		t = time.Now()
		var lastErr error
		timer := time.After(0)
		for {

			select {
			case <-timer:
				fmt.Println("...")
				if n == retries && lastErr != nil {
					result <- fmt.Errorf("maximum retries reached: " + lastErr.Error())
					return
				}

				lastErr = c.Connect()

				if lastErr == nil {
					result <- nil
					return
				} else {
					fmt.Println("error: " + lastErr.Error())
				}

				n++
				timer = time.After(interval)
			case <-ctx.Done():
				result <- fmt.Errorf("context deadline exceeded")
				return
			}
		}
	}()

	return <-result

}
