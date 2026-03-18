package main

import (
	"fmt"
	"sync"
)

var m1 sync.Mutex
var w sync.WaitGroup

func function() {
	m1.Lock()
	fmt.Println("Locked!")
}

func main() {
	w.Add(1)
	go func() {
		defer w.Done()
		function()
	}()

	w.Add(1)
	go func() {
		defer w.Done()
		function()
	}()

	w.Wait()
}
