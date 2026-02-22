package main

import (
	_ "embed"
	"fmt"
)

//go:embed printSource.go
var src string

func main11() {
	fmt.Println(src)
}
