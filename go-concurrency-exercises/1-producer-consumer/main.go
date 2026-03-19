//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer scenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func producer(stream Stream, c chan Tweet) {
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			close(c)
			wg.Done()
			return
		}
		c <- *tweet
	}
}

func consumer(c chan Tweet) {
	defer wg.Done()
	for t := range c {
		if t.IsTalkingAboutGo() {
			fmt.Println(t.Username, "\ttweets about golang")
		} else {
			fmt.Println(t.Username, "\tdoes not tweet about golang")
		}
	}
}

func main() {
	start := time.Now()
	stream := GetMockStream()
	c := make(chan Tweet)

	// Producer
	wg.Add(1)
	go producer(stream, c)

	// Consumer
	wg.Add(1)
	go consumer(c)

	wg.Wait()
	fmt.Printf("Process took %s\n", time.Since(start))
}
