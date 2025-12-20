package main

import (
	"fmt"
	"os"
	"regexp"
)

func matchInt(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`^[-+]?\d+$`)
	return re.Match(t)
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Usage: <utility> one or more strings.")
		return
	}

	var totalTrue, totalFalse int

	for _, s := range arguments[1:] {
		ret := matchInt(s)
		if ret {
			totalTrue++
		} else {
			totalFalse++
		}
		fmt.Println(ret)
	}
	fmt.Println("Total true:", totalTrue)
	fmt.Println("Total false:", totalFalse)
}
