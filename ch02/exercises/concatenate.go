package main

import (
	"fmt"
)

func concat[T any](arr1, arr2 []T) []T {
	result := make([]T, len(arr1)+len(arr2))
	copy(result, arr1)
	copy(result[len(arr1):], arr2)
	return result
}

func main() {
	arr1 := []float64{1.2, 2.3, 3.4}
	arr2 := []float64{4.5, 5.6, 6.7, 7.8, 8.9}

	result := concat(arr1, arr2)
	fmt.Printf("Concat from s1: %v, s2: %v to slice: %v", arr1, arr2, result)
}
