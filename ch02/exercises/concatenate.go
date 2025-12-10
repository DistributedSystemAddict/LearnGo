package main

import "fmt"

func concat[T any](a1 []T, a2 []T) []T {
	result := make([]T, 0, len(a1)+len(a2))
	result = append(result, a1[:]...)
	result = append(result, a2[:]...)
	return result
}

func main() {
	arr1 := [3]int{1, 2, 3}
	arr2 := [5]int{4, 5, 6, 7, 8}

	result := concat[int](arr1, arr2)
	fmt.Printf("Concat from s1: %v, s2: %v to slice: %v", arr1, arr2, result)
}
