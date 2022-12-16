package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	skipDirs := []string{
		//".git",
	}
	_ = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			for _, skipDir := range skipDirs {
				if path == skipDir {
					return filepath.SkipDir
				}
			}
			return nil
		}

		fmt.Println(path)
		return nil
	})
}
