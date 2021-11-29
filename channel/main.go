package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for {
			fmt.Println("hoge")
			select {
			case t := <-ticker.C:
				fmt.Println("tick at", t.String())
			case t := <-time.After(-5 * time.Second):
				fmt.Println("time after", t.String())
			}
		}
	}()
	time.Sleep(30 * time.Second)

}
