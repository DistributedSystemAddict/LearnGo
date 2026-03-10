package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

func wc(file string) (int, int, int, error) { //char, line, word, err
	f, err := os.Open(file)
	if err != nil {
		return 0, 0, 0, err
	}
	char := 0
	line := 0
	word := 0
	defer f.Close()
	r := bufio.NewReader(f)
	re := regexp.MustCompile("[^\\s]+")
	for {
		l, err := r.ReadString('\n')
		if err == io.EOF {
			line += 1
			char += len(l)
			if len(l) != 0 {
				words := re.FindAllString(l, -1)
				word += len(words)
			}
			break
		}
		if len(l) != 0 {
			line += 1
			char += len(l)
			words := re.FindAllString(l, -1)
			word += len(words)

		}
	}
	return char, line, word, nil
}

func main() {
	args := os.Args
	file := args[1]

	char, line, word, err := wc(file)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%7d %7d %7d %s\n", char, line, word, file)
}
