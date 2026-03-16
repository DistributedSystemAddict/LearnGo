package main

import (
	"fmt"
	"sync"
	"time"
)

type value struct {
	mu    sync.Mutex
	value int
}

var wg sync.WaitGroup
var id int

func main() {
	printSum := func(v1, v2 *value) {
		id++
		defer wg.Done()
		v1.mu.Lock()
		defer v1.mu.Unlock()

		time.Sleep(2 * time.Second)
		v2.mu.Lock()
		defer v2.mu.Unlock()

		fmt.Printf("sum=%v, id=%v\n", v1.value+v2.value, id)
	}

	var a, b value
	wg.Add(2)
	go printSum(&a, &b)
	time.Sleep(2 * time.Second)
	go printSum(&b, &a)
	wg.Wait()
}
