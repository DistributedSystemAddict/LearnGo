package main

import "fmt"

func main2() {
	mapp := map[int]string{}
	mapp[0] = "0"
	mapp[1] = "1"
	mapp[2] = "2"

	slice1 := []int{}
	slice2 := []string{}

	for i := 0; i < 3; i++ {
		slice1 = append(slice1, i)
		slice2 = append(slice2, mapp[i])
	}

	fmt.Println(slice1)
	fmt.Println(slice2)
}
