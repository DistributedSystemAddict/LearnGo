package main

import (
	"fmt"
)

func change(s []string) {
	s[0] = "Change_function"
}

func main() {
	a := [5]string{"Zero", "One", "Two", "Three"}
	fmt.Println("a:", a)

	var S0 = a[0:1]
	fmt.Printf("S0: %v, len: %d, cap: %d\n", S0, len(S0), cap(S0))
	S0[0] = "S0"

	var S12 = a[1:3:4]
	fmt.Printf("S12: %v, len: %d, cap: %d\n", S12, len(S12), cap(S12))
	S12[0] = "S12_0"
	S12[4] = "S12_1"

	fmt.Printf("a: %v, len: %d, cap: %d\n", a, len(a), cap(a))

	// Changes to slice -> changes to array
	change(S12)
	fmt.Println("a:", a)

	// capacity of S0
	fmt.Println("Capacity of S0:", cap(S0), "Length of S0:", len(S0))
	fmt.Println("S0:", S0)
	fmt.Println("a:", a)
	// Adding 4 elements to S0
	S0 = append(S0, "N1")
	S0 = append(S0, "N2")
	S0 = append(S0, "N3")
	fmt.Printf("S0[3]: %v\n", S0[3])
	a[0] = "-N1"

	// Changing the capacity of S0
	// Not the same underlying array any more!
	S0 = append(S0, "N4")

	// This change does not go to S0
	a[0] = "-N1-"
	
	// This change does go to S12
	// Because slice S12 is still connected to array a.
	a[1] = "-N2-"
	S0 = append(S0, "N5")
	fmt.Println("Capacity of S0:", cap(S0), "Length of S0:", len(S0))

	fmt.Println("S0:", S0)
	fmt.Println("a: ", a)
	fmt.Println("S12:", S12)
}
