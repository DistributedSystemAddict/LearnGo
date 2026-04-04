package main

import (
	"fmt"
	"sync"
)

func main1() {
	var wg sync.WaitGroup
	sayHello := func() {
		defer wg.Done()
		fmt.Println("hello")
	}
	wg.Add(1)
	go sayHello()
	wg.Wait() // <1>
}
