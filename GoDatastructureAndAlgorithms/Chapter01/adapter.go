package main

import "fmt"

type IProcess interface {
	process()
}

type Adapter struct {
	adaptee Adaptee
}

func (adapter Adapter) process() {
	fmt.Println("Adapter process")
	adapter.adaptee.convert()
}

type Adaptee struct {
	adapterType int
}

// Adaptee class method convert
func (adaptee Adaptee) convert() {
	fmt.Println("Adaptee convert method")
}

// main method
func main4() {

	var processor IProcess = Adapter{}

	processor.process()

}
