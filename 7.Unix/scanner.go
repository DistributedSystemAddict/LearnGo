package main

import (
	"bufio"
	"fmt"
	"os"
)

func main18() {
	if len(os.Args) < 2 {
		return
	}

	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	if err := scanner.Err(); err != nil {
		return
	}
}
