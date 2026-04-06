package main

import "fmt"

var curString string
var n int

func genString(pos int) {
	if pos == n {
		fmt.Println(curString)
		return
	}

	for i := '0'; i <= '1'; i++ {
		curString += string(i)
		genString(pos + 1)
		curString = curString[:len(curString)-1]
	}
}

func main13() {
	n = 5
	genString(0)
}
