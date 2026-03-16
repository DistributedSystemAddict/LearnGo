package main

import (
	"fmt"
	"sync"
	"time"
)

type Foo struct {
	firstDone  chan struct{}
	secondDone chan struct{}
}

func NewFoo() *Foo {
	return &Foo{
		firstDone:  make(chan struct{}),
		secondDone: make(chan struct{}),
	}
}

func (f *Foo) first() {
	fmt.Print("first")
	time.Sleep(2 * time.Second)
	fmt.Println()
	close(f.firstDone)
}

func (f *Foo) second() {
	<-f.firstDone
	fmt.Print("second")
	close(f.secondDone)
}

func (f *Foo) third() {
	<-f.secondDone
	fmt.Print("third")
}

func main() {
	foo := NewFoo()
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		foo.first()
	}()

	go func() {
		defer wg.Done()
		foo.second()
	}()

	go func() {
		defer wg.Done()
		foo.third()
	}()

	wg.Wait()
}
