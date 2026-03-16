package main

import (
	"fmt"
	"os"
	"sync"
)

func main5() {
	count := 20
	fmt.Printf("Going to create %d goroutines.\n", count)
	flag := true
	if len(os.Args) == 1 {
		flag = false
	}

	var waitGroup sync.WaitGroup

	fmt.Printf("%#v\n", waitGroup)
	for i := 0; i < count; i++ {
		waitGroup.Add(1)
		go func(x int) {
			defer waitGroup.Done()
			fmt.Printf("%d ", x)
		}(i)
	}

	if flag {
		waitGroup.Add(1)
	} else {
		waitGroup.Done()
	}

	fmt.Printf("%#v\n", waitGroup)
	waitGroup.Wait()
	fmt.Println("\nExiting...")
}
