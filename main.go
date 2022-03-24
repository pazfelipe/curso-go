package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var folders []string

	err := filepath.Walk("./", func(path string, info os.FileInfo, err error) error {

		if !strings.HasPrefix(path, ".") && !strings.HasPrefix(path, "exercicio") && !strings.HasPrefix(path, "main.go") {
			folders = append(folders, path)
		}

		return nil
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(folders)
}
