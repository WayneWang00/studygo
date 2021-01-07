package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	q := make(chan bool)
	go show(ch, q)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	q <- false
}

func show(ch chan int, q chan bool) {
	for {
		select {
		case n := <-ch:
			fmt.Printf("%d\n", n)
		case <-q:
			break
		}
	}
}
