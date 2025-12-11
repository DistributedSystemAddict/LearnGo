package main

import (
	"fmt"
	"os"
	"strconv"
)

func main4() {
	if len(os.Args) == 1 {
		fmt.Println("Print provide an integer.")
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	input := strconv.Itoa(n) // Chuyển số thành chuỗi
	fmt.Printf("strconv.Itoa() %s of type %T\n", input, input)

	input = strconv.FormatInt(int64(n), 10) // Chuyển số dạng thập phân thành chuỗi
	fmt.Printf("strconv.FormatInt() %s of type %T\n", input, input)

	input = string(n) // Chuyển ký tự thành Unicode
	fmt.Printf("string() %s of type %T\n", input, input)
}
