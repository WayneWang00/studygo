package main

import (
	"errors"
	"fmt"
)

type value int

type Node struct {
	value
	prev, next *Node
}

type List struct {
	head, tail *Node
}

func (l *List) Front() *Node {
	return l.head
}

func (n *Node) Next() *Node {
	return n.next
}

func (l *List) Push(v value) *List {
	n := &Node{value: v}
	if l.head == nil {
		l.head = n
	} else {
		l.tail.next = n
		n.prev = l.tail
	}
	l.tail = n
	return l
}

var errEmpty = errors.New("List is empty!")

func (l *List) Pop() (v value, err error) {
	if l.tail == nil {
		err = errEmpty
	} else {
		v = l.tail.value
		l.tail = l.tail.prev
		//if l.tail == nil {
		//	l.head = nil
		//}
	}
	return v, err
}

func main() {
	l := new(List)
	l.Push(1)
	l.Push(2)
	l.Push(4)
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("List: %v\n", e.value)
	}
	fmt.Println()
	for v, err := l.Pop(); err == nil; v, err = l.Pop() {
		fmt.Printf("v: %v\n", v)
	}
}
