package main

import (
	"fmt"
	"os"
	"regexp"
)

func matchNameSur(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`^[A-Z][a-z]*$`)
	return re.Match(t)
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Usage: <utility> one or more strings.")
		return
	}

	for _, s := range arguments[1:] {
		ret := matchNameSur(s)
		fmt.Println(ret)
	}
}
