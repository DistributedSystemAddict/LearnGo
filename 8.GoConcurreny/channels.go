package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func writeToChannel(c chan int, x int) {
	c <- x
	close(c)
}

func printer(ch chan bool) {
	ch <- true
}

func main7() {
	c := make(chan int, 1)

	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	go func(c chan int) {
		defer waitGroup.Done()
		writeToChannel(c, 10)
		fmt.Println("Exit.")
	}(c)

	//time.Sleep(2 * time.Second) // when used it channel will be closed before reading

	fmt.Println("Read:", <-c)
	_, ok := <-c
	if ok {
		fmt.Println("Channel is open!")
	} else {
		fmt.Println("Channel is closed!")
	}

	waitGroup.Wait()

	var ch chan bool = make(chan bool)

	for i := 0; i < 5; i++ {
		go printer(ch)
	}

	n := 0
	for i := range ch {
		fmt.Println(i)
		if i == true {
			n++
		}
		if n > 2 {
			fmt.Println("n:", n)
			close(ch)
			break
		}
	}

	for i := 0; i < 5; i++ {
		fmt.Println(<-ch)
	}

	var ch1 chan int = make(chan int)
	for i := 0; i < 5; i++ {
		go get(ch1)
	}

	idx := 0
	for i := range ch1 {
		fmt.Println(i)
		idx++
		if idx == 5 {
			close(ch1)
		}
	}

	time.Sleep(5 * time.Second)
}

var id int = 0

func get(ch chan int) {
	i := rand.Intn(5)
	ch <- i
	fmt.Println("get: ", i)
}
