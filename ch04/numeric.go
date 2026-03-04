package main

import (
	"fmt"
)

type Numeric interface {
	int | int8 | int16 | int32 | int64 | float64
}

func Add[U int, T Numeric](a, b T, c U) T {
	fmt.Printf("Type of a: %T\n", a)
	fmt.Printf("Type of b: %T\n", b)
	fmt.Printf("Type of c: %T\n", c)
	return a + b + T(c)
}

func main() {
	fmt.Println("4 + 3 + 2 =", Add(4, 3, 2))
	// fmt.Println("4.1 + 3.2 + 2.3 =", Add(int(4), 3.5, 2)) // sai

	// This is not going to work
	fmt.Println("4.1 + 3 + 2 =", Add(4.0, 2.0, int(3))) // Dm nó work mà???? nó tự ép kiểu của các biến đầu dựa vào biến cuối
}
