package main

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/cybozu-go/well"
)

func main() {
	ctx := context.Background()
	cpu, err := DetectCPUNumber(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("CPU")
	fmt.Print(cpu)
	if cpu == 1 {
		fmt.Println("OK")
	}

	memory, err := DetectMemoryNumber(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("Memory")
	fmt.Print(memory)
}

func DetectCPUNumber(ctx context.Context) (int, error) {
	checkCmd := well.CommandContext(ctx, "sh", "-c", "cat /proc/cpuinfo | grep physical.id | sort -u | wc -l")
	out, err := checkCmd.Output()
	if err != nil {
		return -1, err
	}
	return strconv.Atoi(strings.TrimSpace(string(out)))
}

func DetectMemoryNumber(ctx context.Context) (int, error) {
	checkCmd := well.CommandContext(ctx, "sh", "-c", "dmidecode -t memory | grep '^\\s*Size:' | grep -v \"Size: No Module Installed\" | wc -l")
	out, err := checkCmd.Output()
	if err != nil {
		return -1, err
	}
	return strconv.Atoi(strings.TrimSpace(string(out)))
}
