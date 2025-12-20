package main

import "fmt"

func main() {
	aString := "Hello World! €"
	fmt.Printf("Type of astring: %T\n", aString)
	fmt.Println("First byte ", aString[0])

	r := '€'
	fmt.Println(`As an int32 value: `, r)
	// Convert Runes to text
	fmt.Printf("As a string: %s and as a character: %c\n", r, r)
	// Print an existing string as runes
	for _, v := range aString {
		fmt.Printf("%x ", v)
	}

	fmt.Println()

	for _, v := range aString {
		fmt.Printf("%c", v)
	}
	fmt.Println()
}
