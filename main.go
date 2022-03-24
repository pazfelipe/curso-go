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

		if !strings.HasPrefix(path, "main.go") &&
			!strings.HasPrefix(path, "exercicios") &&
			strings.Contains(path, ".go") {
			folders = append(folders, path)
		}

		return nil
	})

	if err != nil {
		panic(err)
	}

	for _, folder := range folders {
		originalPath := folder

		f := strings.Split(folder, "/")
		newPath := f[0] + "/main.go"

		err := os.Rename(originalPath, newPath)

		if err != nil {
			panic(err)
		}
	}

	fmt.Println("Concluído!!!")
}
