package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func charByChar(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}

	defer f.Close()
	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			for _, c := range line {
				fmt.Println(string(c))
			}
			break
		}
		if err != nil {
			return err
		}
		for _, c := range line {
			fmt.Println(string(c))
		}
	}

	return nil
}

func main3() {
	args := os.Args
	if len(args) == 1 {
		fmt.Printf("usage: byChar <file1> [<file2> ...]\n")
		return
	}

	for _, file := range args[1:] {
		err := charByChar(file)
		if err != nil {
			fmt.Println(err)
		}
	}
}
