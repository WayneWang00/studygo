package main

import "fmt"

type stack struct {
	i    int
	data [10]int
}

func (s *stack) push(k int) {
	s.data[s.i] = k
	s.i++
}

func (s *stack) pop() int {
	s.i--
	fmt.Printf("i: %d\n", s.i)
	return s.data[s.i]
}

func main() {
	var s stack
	s.push(25)
	s.push(14)
	fmt.Printf("My stack: %v\n", s)
	fmt.Printf("My stack: %v\n", s.pop())
}
