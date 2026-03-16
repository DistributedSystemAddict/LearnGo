package main

import "fmt"

func main8() {
	willClose := make(chan complex64, 10)

	willClose <- -1
	willClose <- 1i

	<-willClose
	<-willClose
	close(willClose)

	_, ok := <-willClose
	if ok {
		fmt.Println("Channel is open!")
	} else {
		fmt.Println("Channel is closed!")
	}

}
