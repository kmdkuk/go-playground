package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("now.After(now.Add(1*time.Hour))=", now.After(now.Add(1*time.Hour)))     // false
	fmt.Println("now.After(now.Add(-1*time.Hour))=", now.After(now.Add(-1*time.Hour)))   // true
	fmt.Println("now.Before(now.Add(1*time.Hour))=", now.Before(now.Add(1*time.Hour)))   // true
	fmt.Println("now.Before(now.Add(-1*time.Hour))=", now.Before(now.Add(-1*time.Hour))) // false

	disable := false
	rebootable := false
	if !disable && rebootable {
		fmt.Println("reboot")
	} else {
		fmt.Println("stop reboot")
	}
}
