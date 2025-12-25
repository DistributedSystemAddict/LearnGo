package main

import (
	"fmt"
)

type node[T comparable] struct {
	Data T
	next *node[T]
}

type list[T comparable] struct {
	start *node[T]
}

func (l *list[T]) add(data T) {
	n := node[T]{
		Data: data,
		next: nil,
	}

	if l.start == nil {
		l.start = &n
		return
	}

	if l.start.next == nil {
		l.start.next = &n
		return
	}

	temp := l.start
	l.start = l.start.next
	l.add(data)
	l.start = temp
}

// search finds and returns the first node with the given value
func (l *list[T]) search(data T) *node[T] {
	cur := l.start
	for cur != nil {
		if cur.Data == data {
			return cur
		}
		cur = cur.next
	}
	return nil
}

// delete removes the first node with the given value
func (l *list[T]) delete(data T) bool {
	if l.start == nil {
		return false
	}

	// If deleting the first node
	if l.start.Data == data {
		l.start = l.start.next
		return true
	}

	// Search for the node to delete
	cur := l.start
	for cur.next != nil {
		if cur.next.Data == data {
			cur.next = cur.next.next
			return true
		}
		cur = cur.next
	}

	return false
}

func (l list[T]) printMe() {
	fmt.Printf("%p %p\n", &l, l)
	cur := l.start
	for {
		fmt.Println("*", cur)
		if cur == nil {
			break
		}
		cur = cur.next
	}
}

func main() {
	var myList list[int]
	fmt.Println("Initial list:", myList)
	myList.add(12)
	myList.add(9)
	myList.add(3)
	myList.add(9)

	fmt.Println("\n--- After adding elements ---")
	myList.printMe()

	// Test search
	fmt.Println("\n--- Testing search ---")
	found := myList.search(9)
	if found != nil {
		fmt.Printf("Found node with value 9: %+v\n", found.Data)
	} else {
		fmt.Println("Value 9 not found")
	}

	notFound := myList.search(99)
	if notFound != nil {
		fmt.Printf("Found node with value 99: %+v\n", notFound.Data)
	} else {
		fmt.Println("Value 99 not found")
	}

	// Test delete
	fmt.Println("\n--- Testing delete ---")
	fmt.Println("Deleting value 3...")
	if myList.delete(3) {
		fmt.Println("Successfully deleted 3")
	} else {
		fmt.Println("Failed to delete 3")
	}
	myList.printMe()

	fmt.Println("\nDeleting value 9 (first occurrence)...")
	if myList.delete(9) {
		fmt.Println("Successfully deleted 9")
	} else {
		fmt.Println("Failed to delete 9")
	}
	myList.printMe()
}
