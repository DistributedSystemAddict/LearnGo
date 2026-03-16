package main

import (
	"fmt"
	"time"
)

func printme(x int) {
	fmt.Println("*", x)
}

func main2() {
	go func(x int) {
		fmt.Printf("%d ", x)
	}(10)
	go printme(15)

	time.Sleep(time.Second)
	fmt.Println("Exiting...")
}
