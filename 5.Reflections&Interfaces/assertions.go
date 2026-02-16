package main

import (
	"fmt"
)

func returnNumber() interface{} {
	return 12
}
func main6() {
	anInt := returnNumber()
	number, ok := anInt.(int)
	if ok {
		fmt.Println("Type assertion successful: ", number)
	} else {
		fmt.Println("Type assertion failed!")
	}
	number++
	fmt.Println(number)

	value, ok := anInt.(int64)
	if ok {
		fmt.Println("Type assertion successful: ", value)
	} else {
		fmt.Println("Type assertion failed!")
	}

	i := anInt.(int)
	fmt.Println("i:", i)
	// The following will PANIC because anInt is not bool
	_ = anInt.(bool)
}
