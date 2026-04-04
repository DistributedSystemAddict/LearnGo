package main

import (
	"fmt"
	"sync"
)

func main3() {
	var wg sync.WaitGroup

	for _, salatation := range []string{"hello", "greetings", "good day"} {
		go func(salatation string) {
			defer wg.Done()
			fmt.Println(salatation)
		}(salatation)
	}
	wg.Wait()
}
