package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main19() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run byCharacter.go <filename>")
		return
	}

	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	for {
		char, size, err := reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
			return
		}
		if size > 0 {
			fmt.Printf("Char: %c  | Unicode: %U | Size: %d bytes\n", char, char, size)
		}
	}
}
