package main

import "fmt"

var cur []int
var n1 int
var k int

func getSubset() {
	var lastNum int
	if len(cur) == 0 {
		lastNum = 0
	} else {
		lastNum = cur[len(cur)-1]
	}

	for i := lastNum + 1; i <= n1; i++ {
		cur = append(cur, i)
		if len(cur) == k {
			for j := 0; j < k; j++ {
				fmt.Print(cur[j])
			}
			fmt.Println()
		} else {
			getSubset()
		}
		cur = cur[:len(cur)-1]
	}
}

func main() {
	n1 = 5
	k = 4
	getSubset()
}
