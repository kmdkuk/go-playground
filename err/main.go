package main

import "fmt"

func hoge() error {
	return nil
}

func main() {
	err := hoge()
	fmt.Println(err.Error())
}
