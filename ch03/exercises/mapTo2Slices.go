package main

import "fmt"

func main() {
	mapp := map[int]string{1: "one", 2: "two", 3: "three", 4: "four", 5: "five", 6: "six"}

	var keys []int
	var values []string

	for key, value := range mapp {
		keys = append(keys, key)
		values = append(values, value)
	}

	for i := 0; i < len(keys); i++ {
		fmt.Println(keys[i], values[i])
	}
}
