package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/cybozu-go/log"
	"github.com/cybozu-go/well"
)

func main() {
	well.Go(func(ctx context.Context) (err error) {
		fmt.Println("go sleep")
		defer func() {
			fmt.Println("defer")
			time.Sleep(3 * time.Second)
			// Fails if it has already been Canceled using the same ctx in defer.
			c := exec.CommandContext(ctx, "curl", "google.com")
			// c := exec.CommandContext(context.Background(), "curl", "google.com")
			c.Stdout = os.Stdout
			c.Stderr = os.Stderr
			err2 := c.Run()
			if err2 != nil {
				err = err2
				return
			}
			fmt.Println("defer done")
		}()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("done")
				return nil
			case <-time.After(100 * time.Second):
				return fmt.Errorf("time over")
			}
		}
	})

	well.Stop()
	err := well.Wait()
	if err != nil {
		if well.IsSignaled(err) {
			fmt.Printf("catch signal %v\n", err)
		}
		log.ErrorExit(err)
	}
}
