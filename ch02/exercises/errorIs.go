package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

var ErrNumber = errors.New("number error")

func handleError() {
	if r := recover(); r != nil {
		err, ok := r.(error)
		if !ok {
			err = fmt.Errorf("%v", r)
		}

		fmt.Println("Bắt được:", err)

		if errors.Is(err, ErrNumber) {
			fmt.Println("hahaaha")
		}
	}
}

func isNumber(a string) bool {
	_, err := strconv.ParseFloat(a, 64)
	if err != nil {
		panic(ErrNumber)
	}
	return true
}

func main() {
	input := os.Args[1]
	defer handleError()
	isNumber(input)

	fmt.Println("Can convert to number")
}
