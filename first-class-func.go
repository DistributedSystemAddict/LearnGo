package main

import "fmt"

func chao() {
	fmt.Println("xin chao")
}

func chayHam(f func()) {
	f() // gọi hàm được truyền vào
}

func main() {
	chayHam(chao) // truyền hàm như một giá trị
}
