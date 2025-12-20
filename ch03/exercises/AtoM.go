package main

import "fmt"

func main() {
	arrays := [6]int{1, 2, 3, 4, 5, 6}

	mapp := make(map[int]int)
	for _, value := range arrays {
		mapp[value] = value
	}

	for key, value := range mapp {
		fmt.Println(key, value)
	}
}
