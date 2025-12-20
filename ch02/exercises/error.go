package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

var customError = errors.New("this is a custom error message")

// Custom error message with errors.New()
func check(a, b int) error {
	if a == 0 && b == 0 {
		return customError
	}
	return nil
}

// Custom error message with fmt.Errorf()
func formattedError(a, b int) error {
	if a == 0 && b == 0 {
		return fmt.Errorf("a %d and b %d. UserID: %d", a, b, os.Getuid())
	}
	return nil
}

func main() {
	err := check(0, 10)
	if !errors.Is(err, customError) {
		fmt.Println("check() executed normally!")
	} else {
		fmt.Println(err)
	}

	err = check(0, 0)
	if errors.Is(err, customError) {
		fmt.Println("Custom error detected!")
	}

	err = formattedError(0, 0)
	if err != nil {
		fmt.Println(err)
	}

	i, err := strconv.Atoi("-123")
	if err == nil {
		fmt.Println("Int value is", i)
	}

	i, err = strconv.Atoi("Y123")
	if err != nil {
		fmt.Println(err)
	}
}