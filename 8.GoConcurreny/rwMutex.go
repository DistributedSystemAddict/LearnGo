package main

import (
	"fmt"
	"sync"
	"time"
)

var Password *secret
var wg2 sync.WaitGroup

type secret struct {
	RWM      sync.RWMutex
	password string
}

func Change(pass string) {
	if Password == nil {
		fmt.Println("Password is nil!")
		return
	}

	fmt.Println("Change() function")
	Password.RWM.Lock()
	fmt.Println("Change() Locked")
	time.Sleep(4 * time.Second)
	Password.password = pass
	Password.RWM.Unlock()
	fmt.Println("Change() Unlocked")
}

func show() {
	defer wg2.Done()
	defer Password.RWM.RUnlock()
	Password.RWM.RLock()
	fmt.Println("Show function locked!")
	time.Sleep(2 * time.Second)
	fmt.Println("Pass value:", Password.password)
}

func main22() {
	Password = &secret{password: "myPass"}
	for i := 0; i < 3; i++ {
		wg2.Add(1)
		go show()
	}

	wg2.Add(1)
	go func() {
		defer wg2.Done()
		Change("123456")
	}()

	wg2.Add(1)
	go func() {
		defer wg2.Done()
		Change("654321")
	}()

	wg2.Wait()

	// Direct access to Password.password
	fmt.Println("Current password value:", Password.password)
}
