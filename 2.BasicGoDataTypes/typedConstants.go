package main

import "fmt"

const (
	typedConstant   = int16(100)
	untypedConstant = 100
)

func main8() {
	i := int(1)
	fmt.Println("untyped:", i*untypedConstant)
	fmt.Println("Typed:", i*int(typedConstant))
}
