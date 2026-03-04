package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func Same[T comparable](a, b T) bool {
	// or
	// return a == b
	if a == b {
		return true
	}
	return false
}

func main() {
	fmt.Println("4 = 4 is", Same(4, 4))
	fmt.Println("aa = aa is", Same("aa", "aa"))
	fmt.Println("4.1 = 4.15 is", Same(4.1, 4.15))

	s1 := "aa"
	s2 := "aa"
	s3 := fmt.Sprintf("%s", "aa")

	// Lấy địa chỉ Data trực tiếp từ StringHeader
	p1 := (unsafe.StringData(s1))
	p2 := (*reflect.StringHeader)(unsafe.Pointer(&s2)).Data
	p3 := (*reflect.StringHeader)(unsafe.Pointer(&s3)).Data

	fmt.Printf("Địa chỉ thực sự s1 trỏ đến: 0x%x\n", p1)
	fmt.Printf("Địa chỉ thực sự s2 trỏ đến: 0x%x\n", p2)
	fmt.Printf("Địa chỉ thực sự s3 trỏ đến: 0x%x\n", p3)

	// This is not going to work
	// _ = Same([]int{1, 2}, []int{1, 3})
}
