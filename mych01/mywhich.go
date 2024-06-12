package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatal("please provide an argument")
	}

	file := os.Args[1]
	path := os.Getenv("PATH")
	dirs := filepath.SplitList(path)
	for _, dir := range dirs {

		fullPath := filepath.Join(dir, file)
		fi, err := os.Stat(fullPath)
		if err != nil {
			continue
		}

		if !fi.Mode().IsRegular() {
			continue
		}

		if isExecutable(fi.Mode()) {
			fmt.Println(fullPath)
			return
		}
	}
}

func isExecutable(mode os.FileMode) bool {
	return mode&0111 != 0
}
