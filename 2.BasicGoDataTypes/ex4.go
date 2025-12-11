package main

import "fmt"

func Concat2[T any](a, b []T) []T {
	total := len(a) + len(b)
	result := make([]T, total)
	copy(result, a)
	copy(result[len(a):], b)
	return result
}

func main24() {
	a := []int{1, 2, 3}
	b := []int{4, 5}

	s := Concat2(a, b)
	arr := (*[5]int)(s)
	fmt.Println(*arr)
}
