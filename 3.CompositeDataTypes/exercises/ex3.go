package main

import (
	"fmt"
	"os"
)

type Arg struct {
	Index int
	Value string
}

func argsToStructSlice() []Arg {
	args := os.Args
	result := make([]Arg, 0, len(args))

	for i, v := range args {
		result = append(result, Arg{
			Index: i,
			Value: v,
		})
	}
	return result
}

func main() {
	argStructs := argsToStructSlice()

	for _, arg := range argStructs {
		fmt.Printf("Index: %d, Value: %s\n", arg.Index, arg.Value)
	}
}
