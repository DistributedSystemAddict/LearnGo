package main

import "fmt"

// Concat two arrays into a new slice
func concatArrayToSlice[T any](a, b []T) []T {
	// Tạo slice mới với capacity = len(a) + len(b)
	result := make([]T, 0, len(a)+len(b))
	result = append(result, a...)
	result = append(result, b...)
	return result
}

func main22() {
	arr1 := [3]int{1, 2, 3}
	arr2 := [2]int{4, 5}

	// Chuyển array → slice bằng arr1[:]
	result := concatArrayToSlice(arr1[:], arr2[:])

	fmt.Println(result) // [1 2 3 4 5]
}
