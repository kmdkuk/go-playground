package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/cybozu-go/log"
	"github.com/cybozu-go/well"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 21*time.Second)
	defer cancel()
	env := well.NewEnvironment(ctx)
	env.Go(func(ctx context.Context) error {
		command := well.CommandContext(ctx, "./bin/welltry")
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr
		fmt.Println("exec go run ./welltry/main.go")
		return command.Run()
	})
	env.Stop()
	err := env.Wait()
	if err != nil {
		log.ErrorExit(err)
	}
}
