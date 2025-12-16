package main

import "fmt"

func main1() {
	aMap := make(map[string]string)
	aMap["123"] = "456"
	aMap["key"] = "A value"
	for key, v := range aMap {
		fmt.Println("key:", key, "value", v)
	}

	for _, v := range aMap {
		fmt.Println(" # ", v)
	}
	fmt.Println()
}
