package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	cmdPath, _ := filepath.Abs(filepath.Join(".", "env.sh"))
	cmd := exec.Command(cmdPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = append(os.Environ(), "HOGE=")
	if err := cmd.Run(); err != nil {
		fmt.Printf("error: %+v", err)
		os.Exit(1)
	}
}
