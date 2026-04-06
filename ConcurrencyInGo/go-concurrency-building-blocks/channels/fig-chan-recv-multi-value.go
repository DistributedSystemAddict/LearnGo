package main

import "fmt"

func main4() {
	stringStream := make(chan string)
	go func() {
		stringStream <- "Hello channels!"
	}()

	salatation, ok := <-stringStream
	fmt.Printf("(%v): %v", ok, salatation)
}
