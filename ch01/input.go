package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Get User Input
	fmt.Printf("Please give me your name: ")

	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n') // Đọc đến khi gặp '\n'

	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	fmt.Println("Your name is", name)
}
