package main

import (
	"fmt"
	"sync"
)

func main2() {
	var wg sync.WaitGroup
	for _, salatation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(salatation)
		}()
	}
}
