package main

import (
	"fmt"
)

func main0() {
	aMap := map[string]int{}
	aMap["test"] = 1
	aMap["test1"] = 2
	fmt.Println("aMap:", aMap)
	aMap = nil
	fmt.Println("aMap", aMap)
	if aMap == nil {
		fmt.Println("nil map!")
		aMap = map[string]int{}
	}
	aMap["test"] = 1

	// Cái này sẽ crash bởi vì map bị khai báo là nil thì
	// phải khởi tạo lại
	// aMap = nil
	// aMap["test"] = 1
}
