package main

import (
	"bufio"
	"fmt"
	"os"
)

func lineByLine(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}

	defer f.Close()

	//r := bufio.NewReader(f)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
	// for {
	// 	line, err := r.ReadString('\n')
	// 	if err == io.EOF {
	// 		if len(line) != 0 {
	// 			fmt.Println(line)
	// 		}
	// 		break
	// 	}

	// 	if err != nil {
	// 		fmt.Printf("error reading file %s", err)
	// 		return err
	// 	}
	// 	fmt.Print(line)
	// }
	return nil
}

func main2() {
	args := os.Args
	if len(args) == 1 {
		fmt.Printf("usage: byLine <file1> [<file2> ...]\n")
		return
	}

	for _, file := range args[1:] {
		err := lineByLine(file)
		if err != nil {
			fmt.Println(err)
		}
	}
}
