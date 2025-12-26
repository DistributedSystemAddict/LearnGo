package main

import "fmt"

type node[T comparable] struct {
	Data T
	next *node[T]
}

type list[T comparable] struct {
	start *node[T]
}

func (l *list[T]) isExist(data T) bool {
	curr := l.start
	for curr != nil {
		if curr.Data == data {
			return true
		}
		curr = curr.next
	}
	return false
}

func (l *list[T]) delete(data T) {
	if l.start == nil {
		return
	}

	if l.start.Data == data {
		l.start = l.start.next
		return
	}
	curr := l.start
	for curr.next != nil {
		if curr.next.Data == data {
			curr.next = curr.next.next
			return
		}
		curr = curr.next
	}
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

func main() {
	var myList list[int]

	fmt.Println(myList)
	myList.add(12)
	// temp := myList.start
	myList.add(9)
	myList.add(3)
	myList.add(9)
	myList.delete(9)
	// myList.start = temp
	cur := myList.start
	check := myList.isExist(3)
	fmt.Println(check)
	for {
		fmt.Println("*", cur)
		if cur == nil {
			break
		}
		cur = cur.next
	}

}
