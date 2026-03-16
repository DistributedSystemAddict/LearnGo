package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func createFile(file string) {
	_, err := os.Stat(file)
	if err == nil {
		fmt.Printf("%s already exists!\n", file)
		return
	}

	// Can the file be created?
	f, err := os.Create(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	lines := random(10, 30)

	for i := 0; i < lines; i++ {
		data := random(0, 20)
		fmt.Fprintf(f, "%d\n", data)
	}
	fmt.Printf("%s created!\n", file)
}

func main6() {
	arguments := os.Args
	if len(arguments) != 5 {
		fmt.Println("Usage: randomFiles firstInt lastInt filename sirectory")
		return
	}

	start, err := strconv.Atoi(arguments[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	end, err := strconv.Atoi(arguments[2])
	if err != nil {
		fmt.Println(err)
		return
	}

	if end < start {
		fmt.Println(end, "<", start)
		return
	}

	filename := arguments[3]
	path := arguments[4]

	_, err = os.Open(path)
	if err != nil {
		fmt.Println(path, "does not exist!")
		return
	}

	var waitGroup sync.WaitGroup
	for i := start; i <= end; i++ {
		waitGroup.Add(1)

		go func(n int) {
			filepath := filepath.Join(path, fmt.Sprintf("%s%d", filename, n))
			defer waitGroup.Done()
			createFile(filepath)
		}(i)
	}
	waitGroup.Wait()
}
