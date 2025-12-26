package main

import "fmt"

type node[T any] struct {
	Data T
	prev *node[T]
	next *node[T]
}

type list[T any] struct {
	start *node[T]
}

func (l *list[T]) add(data T) {
	n := node[T]{
		Data: data,
		prev: l.start,
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
	myList.add(10)
	myList.add(10)
	myList.add(10)
	// myList.start = temp
	cur := myList.start
	for {
		fmt.Println("*", cur)
		if cur == nil {
			break
		}
		if cur.prev != nil {
			fmt.Println(cur.prev.Data)
		}
		cur = cur.next
	}

}
