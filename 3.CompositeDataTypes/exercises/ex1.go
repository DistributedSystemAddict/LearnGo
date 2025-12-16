package main

import "fmt"

func main1() {
	arr := [5]string{"apple", "banana", "cherry", "date", "elderberry"}
	res := make(map[int]string)

	for i := 0; i < len(arr); i++ {
		res[i] = arr[i]
	}

	fmt.Println(res)
}
