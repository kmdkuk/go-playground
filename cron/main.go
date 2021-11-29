package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func print() {
	fmt.Println("hoge")
}

func main() {
	schedule := "@every 2s"
	c := cron.New()
	if _, err := c.AddFunc(schedule, print); err != nil {
		panic(err)
	}
	c.Start()
	time.Sleep(1000 * time.Second)
}
