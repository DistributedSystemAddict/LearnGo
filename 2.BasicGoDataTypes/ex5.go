package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// define reusable sentinel error
var ErrCustom = errors.New("this is a custom error message")

func check(a, b int) error {
	if a == 0 && b == 0 {
		// wrap ErrCustom using %w for errors.Is()
		return fmt.Errorf("%w", ErrCustom)
	}
	return nil
}

func formattedError(a, b int) error {
	if a == 0 && b == 0 {
		// keep formatting, also wrap the sentinel
		return fmt.Errorf("a %d and b %d. UserID: %d: %w", a, b, os.Getuid(), ErrCustom)
	}
	return nil
}

func main() {
	err := check(0, 10)
	if err == nil {
		fmt.Println("check() executed normally!")
	} else {
		fmt.Println(err)
	}

	// using errors.Is() instead of err.Error()
	err = check(0, 0)
	if errors.Is(err, ErrCustom) {
		fmt.Println("Custom error detected!")
	}

	err = formattedError(0, 0)
	if err != nil {
		// still prints full message
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
