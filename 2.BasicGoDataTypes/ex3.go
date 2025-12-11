package main

import "fmt"

// Cannot

// func concatArrays[T any, N1, N2 int](a [N1]T, b [N2]T) [N1 + N2]T {
// 	var result [N1 + N2]T

// 	for i := 0; i < N1; i++ {
// 		result[i] = a[i]
// 	}
// 	for i := 0; i < N2; i++ {
// 		result[N1+i] = b[i]
// 	}
// 	return result
// }

// Concatenate two arrays of FIXED LENGTHS (not generic)
func concat3and2[T any](a [3]T, b [2]T) [5]T {
	var result [5]T
	copy(result[:], a[:])
	copy(result[3:], b[:])
	return result
}

func main23() {
	a := [3]int{1, 2, 3}
	b := [2]int{4, 5}

	//r := concatArrays(a, b)
	r1 := concat3and2(a, b)

	// fmt.Println(r) // [1 2 3 4 5]
	fmt.Println(r1)
}
