package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	LOGFILE := "myLog.log"
	fmt.Println(LOGFILE)
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// The call to os.OpenFile() creates the log file for writing,
	// if it does not already exist, or opens it for writing
	// by appending new data at the end of it (os.O_APPEND)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	iLog := log.New(f, "iLog ", log.LstdFlags)
	iLog.Println("Hello there!")
	iLog.Println("Mastering Go 4th edition!")
}
