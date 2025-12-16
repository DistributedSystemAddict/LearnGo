package main

import (
	"fmt"
	"os"
	"regexp"
)

func matchNameSur(s string) bool {
	t := []byte(s) //Chuyển một chuỗi string thành 1 slice các byte
	fmt.Println(t)
	re := regexp.MustCompile(`^[A-Z][a-z]*$`)
	return re.Match(t)
}

func main5() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Usage: <utility> string.")
		return
	}

	s := arguments[1]
	ret := matchNameSur(s)
	fmt.Println(ret)

	for _, s := range arguments[1:] {
		ret := matchNameSur(s)
		fmt.Println(ret)
	}
}
