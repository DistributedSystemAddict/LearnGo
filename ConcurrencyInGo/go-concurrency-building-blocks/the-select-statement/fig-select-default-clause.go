package main

import (
	"fmt"
	"time"
)

func main2() {
	start := time.Now()
	var c1, c2 <-chan int
	select {
	case <-c1:
	case <-c2:
	default:
		fmt.Printf("In default after %c\n\n", time.Since(start))
	}
}
