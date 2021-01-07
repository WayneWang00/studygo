package main

import "fmt"

func main() {
	ch := make(chan int)
	go show(ch)
	for i := 0; i < 10; i++ {
		ch <- i
	}
}

func show(ch chan int) {
	for {
		j := <-ch
		fmt.Printf("%d\n", j)
	}
}
