package main

import (
	"fmt"
)

func main6() {
	fmt.Printf("Please give me your name: ")
	var name string
	fmt.Scanln(&name)
	fmt.Println("Your name is", name)
}
