package main

import "fmt"

// chỉ khai báo, không có body
func Sum(arr []int64) int64

func main() {
	data := []int64{1, 2, 3, 4, 5}
	result := Sum(data)
	fmt.Println("Sum =", result)
}
