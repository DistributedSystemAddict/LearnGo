package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name  string
	Grade float64
}

type Product struct {
	Name  string
	Price float64
}

type StudentSlice []Student

func (s StudentSlice) Len() int {
	return len(s)
}

func (s StudentSlice) Less(i, j int) bool {
	return s[i].Grade < s[j].Grade
}

func (s StudentSlice) Swap(i, j int) {
	s[j], s[j] = s[j], s[i]
}

func PrintInfo(x interface{}) {
	switch v := x.(type) {
	case Student:
		fmt.Println("Student:", v.Name, "Grade:", v.Grade)

	case Product:
		fmt.Println("Product:", v.Name, "Price:", v.Price)

	default:
		fmt.Println("Unknown type")
	}
}

func main() {

	// Create slice of Students
	students := StudentSlice{
		{"Alice", 8.5},
		{"Bob", 6.2},
		{"Charlie", 9.1},
	}

	// Sort by Grade
	sort.Sort(students)

	fmt.Println("Sorted Students:")
	for _, s := range students {
		fmt.Println(s.Name, s.Grade)
	}

	fmt.Println("\nType Switch Demo:")
	PrintInfo(students[0])
	PrintInfo(Product{"Laptop", 1500})
	PrintInfo(123) // Unknown
}
