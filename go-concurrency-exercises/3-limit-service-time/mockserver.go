//////////////////////////////////////////////////////////////////////
//
// DO NOT EDIT THIS PART
// Your task is to edit `main.go`
//

package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// RunMockServer pretends to be a video processing service. It
// simulates user interacting with the Server.
func RunMockServer() {
	u1 := User{ID: 0, IsPremium: false}
	u2 := User{ID: 1, IsPremium: true}

	wg.Add(5)

	go createMockRequest(1, shortProcess, &u1)
	time.Sleep(1 * time.Second)

	go createMockRequest(2, longProcess, &u2)
	time.Sleep(2 * time.Second)

	go createMockRequest(3, shortProcess, &u1)
	time.Sleep(1 * time.Second)

	go createMockRequest(4, longProcess, &u1)
	go createMockRequest(5, shortProcess, &u2)

	wg.Wait()
}

var inJob = make(map[int]chan struct{})
var m sync.Mutex

func check(id int, ch chan struct{}) {
	time.Sleep(10 * time.Second)
	close(ch)
	m.Lock()
	delete(inJob, id)
	m.Unlock()
}

func createMockRequest(pid int, fn func(), u *User) {
	var ch chan struct{}

	m.Lock()
	if _, ok := inJob[u.ID]; !ok {
		ch = make(chan struct{})
		inJob[u.ID] = ch
		go check(u.ID, ch)
	}
	ch = inJob[u.ID]
	m.Unlock()

	fmt.Printf("UserID: %d \tProcess %d started.\n", u.ID, pid)

	res := HandleRequest(fn, u, ch)

	if res {
		fmt.Printf("UserID: %d \tProcess %d done.\n", u.ID, pid)
	} else {
		fmt.Printf("UserID: %d \tProcess %d killed. (No quota left)\n", u.ID, pid)
	}

	wg.Done()
}

func shortProcess() {
	time.Sleep(6 * time.Second)
}

func longProcess() {
	time.Sleep(11 * time.Second)
}
