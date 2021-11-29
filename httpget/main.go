package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	c := http.Client{}
	var data []byte
	resp, err := c.Get("https://github.com/cybozu-go/neco/releases/download/release-2021.08.30-15651/neco-operation-cli-linux_2021.08.30-15651_amd64.deb")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	data, err = io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	f, err := os.Create("neco-operation-cli-linux_2021.08.30-15651_amd64.deb")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.Write(data)
	if err != nil {
		panic(err)
	}
}
