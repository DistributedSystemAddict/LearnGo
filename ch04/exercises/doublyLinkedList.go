package main

import (
	"fmt"
)

// Doubly-linked list node
type dNode[T comparable] struct {
	Data T
	next *dNode[T]
	prev *dNode[T]
}

// Doubly-linked list
type dList[T comparable] struct {
	start *dNode[T]
	end   *dNode[T]
}

// add adds a new node to the end of the doubly-linked list
func (l *dList[T]) add(data T) {
	n := &dNode[T]{
		Data: data,
		next: nil,
		prev: nil,
	}

	if l.start == nil {
		// First node
		l.start = n
		l.end = n
		return
	}

	// Add to the end
	n.prev = l.end
	l.end.next = n
	l.end = n
}

// addFront adds a new node to the beginning of the list
func (l *dList[T]) addFront(data T) {
	n := &dNode[T]{
		Data: data,
		next: nil,
		prev: nil,
	}

	if l.start == nil {
		l.start = n
		l.end = n
		return
	}

	n.next = l.start
	l.start.prev = n
	l.start = n
}

// search finds and returns the first node with the given value
func (l *dList[T]) search(data T) *dNode[T] {
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
func (l *dList[T]) delete(data T) bool {
	if l.start == nil {
		return false
	}

	// Find the node to delete
	cur := l.start
	for cur != nil {
		if cur.Data == data {
			// Found the node to delete
			if cur.prev == nil {
				// Deleting the first node
				l.start = cur.next
				if l.start != nil {
					l.start.prev = nil
				} else {
					// List is now empty
					l.end = nil
				}
			} else if cur.next == nil {
				// Deleting the last node
				l.end = cur.prev
				l.end.next = nil
			} else {
				// Deleting a middle node
				cur.prev.next = cur.next
				cur.next.prev = cur.prev
			}
			return true
		}
		cur = cur.next
	}

	return false
}

// printForward prints the list from start to end
func (l *dList[T]) printForward() {
	fmt.Print("Forward: ")
	cur := l.start
	for cur != nil {
		fmt.Printf("%v", cur.Data)
		if cur.next != nil {
			fmt.Print(" <-> ")
		}
		cur = cur.next
	}
	fmt.Println()
}

// printBackward prints the list from end to start
func (l *dList[T]) printBackward() {
	fmt.Print("Backward: ")
	cur := l.end
	for cur != nil {
		fmt.Printf("%v", cur.Data)
		if cur.prev != nil {
			fmt.Print(" <-> ")
		}
		cur = cur.prev
	}
	fmt.Println()
}

// To test the doubly-linked list, rename this function to main()
// and comment out main() in structures.go
func demoDoublyLinkedList() {
	var myDList dList[int]

	fmt.Println("=== Doubly-Linked List Demo ===\n")

	// Add elements
	fmt.Println("Adding elements: 10, 20, 30, 40")
	myDList.add(10)
	myDList.add(20)
	myDList.add(30)
	myDList.add(40)
	myDList.printForward()
	myDList.printBackward()

	// Add to front
	fmt.Println("\nAdding 5 to the front:")
	myDList.addFront(5)
	myDList.printForward()
	myDList.printBackward()

	// Search
	fmt.Println("\n--- Testing search ---")
	found := myDList.search(30)
	if found != nil {
		fmt.Printf("Found node with value 30: %v\n", found.Data)
		var prevVal, nextVal interface{} = nil, nil
		if found.prev != nil {
			prevVal = found.prev.Data
		}
		if found.next != nil {
			nextVal = found.next.Data
		}
		fmt.Printf("Previous: %v, Next: %v\n", prevVal, nextVal)
	}

	// Delete middle node
	fmt.Println("\n--- Testing delete (middle node: 20) ---")
	if myDList.delete(20) {
		fmt.Println("Successfully deleted 20")
		myDList.printForward()
		myDList.printBackward()
	}

	// Delete first node
	fmt.Println("\n--- Testing delete (first node: 5) ---")
	if myDList.delete(5) {
		fmt.Println("Successfully deleted 5")
		myDList.printForward()
		myDList.printBackward()
	}

	// Delete last node
	fmt.Println("\n--- Testing delete (last node: 40) ---")
	if myDList.delete(40) {
		fmt.Println("Successfully deleted 40")
		myDList.printForward()
		myDList.printBackward()
	}

	// Test with strings
	fmt.Println("\n=== Testing with strings ===")
	var strList dList[string]
	strList.add("Hello")
	strList.add("World")
	strList.add("Go")
	strList.printForward()
	strList.printBackward()

	fmt.Println("\nDeleting 'World':")
	strList.delete("World")
	strList.printForward()
	strList.printBackward()
}

// Uncomment to test doubly-linked list separately
func main() {
	demoDoublyLinkedList()
}
