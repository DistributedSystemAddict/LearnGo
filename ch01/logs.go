package main

import (
	"log"
	"os"
)

func myLogger() {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered from panic: %v\n", r)
		}
	}()
	log.Panic("Panic: Hello World!")

}

func main() {
	if len(os.Args) != 1 {
		log.Fatal("Fatal: Hello World!")
	}
	
	defer log.Println("Deferred log: Hello World!")
	
	// Recover panic


	myLogger()
	log.Println("Log: Hello World!")
}