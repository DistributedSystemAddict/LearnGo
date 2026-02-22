package main

import "fmt"

type animal interface {
	breathe()
	walk()
}

type mammal interface {
	feed()
}

type lion struct {
	age int
}

func (l lion) breathe() {
	fmt.Println("Lion breathes")
}
func (l lion) walk() {
	fmt.Println("Lion walk")
}
func (l lion) feed() {
	fmt.Println("Lion feeds young")
}
func main() {
	var a animal
	l := lion{}
	a = l
	a.breathe()
	a.walk()
	var m mammal
	m = l
	m.feed()
}
