package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	e1 := l.PushFront(1)
	e2 := l.InsertAfter(2, e1)
	l.InsertAfter(4, e2)
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("link list: %v", e.Value)
	}
}
