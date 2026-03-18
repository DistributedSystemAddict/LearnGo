package main

import (
	"fmt"
	"time"
)

func printer2(ch chan<- bool, times int) {
	for i := 0; i < times; i++ {
		ch <- true
	}
	close(ch)
}

func main10() {
	var ch chan bool = make(chan bool)
	go printer2(ch, 5)

	time.Sleep(10 * time.Second)
	for val := range ch {
		fmt.Print(val, " ")
	}
	fmt.Println()

	for i := 0; i < 15; i++ {
		fmt.Print(<-ch, " ")
	}
	fmt.Println()
}
