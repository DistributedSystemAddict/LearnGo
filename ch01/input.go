package main
import (
	"fmt"
)


func main() {
	fmt.Println("Please tell me your name")
	var name string
	fmt.Scanln(&name)
	fmt.Println("hello", name)
}