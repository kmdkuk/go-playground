package main

import (
	"fmt"
	"os"
)

func main() {
	dirs, err := os.ReadDir(".")
	if err != nil {
		panic(err)
	}
	for _, d := range dirs {
		if d.IsDir() {
			fmt.Printf("%+v\n", d.Name())
		}
	}
}
