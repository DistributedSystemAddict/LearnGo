//////////////////////////////////////////////////////////////////////
//
// Given is a mock process which runs indefinitely and blocks the
// program. Right now the only way to stop the program is to send a
// SIGINT (Ctrl-C). Killing a process like that is not graceful, so we
// want to try to gracefully stop the process first.
//
// Change the program to do the following:
//   1. On SIGINT try to gracefully stop the process using
//          `proc.Stop()`
//   2. If SIGINT is called again, just kill the program (last resort)
//

package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var wg sync.WaitGroup

func main() {
	// Create a process
	proc := MockProcess{}

	// Run the process (blocking)
	go proc.Run()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT)
	graceful := false
	wg.Add(2)
	go func() {
		for {
			sig := <-sigs
			fmt.Println("\nreceived signal:", sig)

			if !graceful {
				wg.Done()
				graceful = true
				go proc.Stop() // gọi graceful stop
			} else {
				wg.Done()
				fmt.Println("force exit!")
				os.Exit(1)
			}
		}
	}()

	wg.Wait()
}
