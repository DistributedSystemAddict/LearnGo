package main

import "fmt"

type TreeLast1[T any] []T

func (t TreeLast1[T]) replaceLastValue(x T, val T) {
	fmt.Printf("\n--- Inside Value Receiver ---\n")
	t[len(t)-1] = x
	t = append(t, val)
	fmt.Printf("Bản copy trong hàm: %v, Pointer mảng: %p, Len: %d\n", t, t, len(t))
}

func (t *TreeLast1[T]) replaceLastPointer(x T, val T) {
	fmt.Printf("\n--- Inside Pointer Receiver ---\n")
	if len(*t) > 0 {
		(*t)[len(*t)-1] = x
	}

	*t = append(*t, val)
	fmt.Printf("Gốc (thông qua con trỏ): %v, Pointer mảng: %p, Len: %d\n", *t, *t, len(*t))
}

func main7() {
	// Demo với kiểu int
	slice1 := TreeLast1[int]{1, 2, 3}
	fmt.Printf("Ban đầu slice1: %v, Pointer mảng: %p\n", slice1, slice1)

	// Truyền 9 để sửa và 100 để append
	slice1.replaceLastValue(9, 100)
	fmt.Printf("Sau Value Receiver: %v, Len: %d\n", slice1, len(slice1))

	// fmt.Println("\n" + "=".repeat(50))

	slice2 := TreeLast1[int]{1, 2, 3}
	fmt.Printf("Ban đầu slice2: %v, Pointer mảng: %p\n", slice2, slice2)

	slice2.replaceLastPointer(9, 100)
	fmt.Printf("Sau Pointer Receiver: %v, Len: %d\n", slice2, len(slice2))
}
