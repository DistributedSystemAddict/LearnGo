package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main2() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide an arguments!")
		os.Exit(1)
	}
	filename := arguments[1]
	fileinfo, err := os.Lstat(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if fileinfo.Mode()&os.ModeSymlink != 0 {
		fmt.Print(filename, "is a symbolic link")
		realpath, err := filepath.EvalSymlinks(filename)
		if err == nil {
			fmt.Println("Path:", realpath)
		}
	}

}
