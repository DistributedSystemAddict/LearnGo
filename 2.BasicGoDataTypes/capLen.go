package main

import "fmt"

func mai10() {
	a := make([]int, 4)
	fmt.Println("L:", len(a), "C:", cap(a))

	b := []int{0, 1, 2, 3, 4}
	fmt.Println("L:", len(b), "C:", cap(b))

	aSlice := make([]int, 4, 4)
	fmt.Println(aSlice)

	aSlice = append(aSlice, 5)
	fmt.Println(aSlice)

	fmt.Println("L:", len(aSlice), "C:", cap(aSlice))

	aSlice = append(aSlice, []int{-1, -2, -3, -4}...)
	// ... operator expands []int{-1, -2, -3, -4} into multiple arguments and append() appends each argument one by one to aSlice
	fmt.Println(aSlice)

	fmt.Println("L:", len(aSlice), "C:", cap(aSlice))
}
