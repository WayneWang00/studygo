package main

import "fmt"

func push(s []int, a int) (s1 []int) {
	s1 = append(s, a)
	fmt.Println(s1)
	return
}

func pop(s []int) (s2 []int) {
	s2 = append(s2, s[:len(s)-1]...)
	return s2
}

func main() {
	s := []int{1, 2, 3}
	s1 := push(s, 1)
	fmt.Printf("push: %v\n", s1)
	s2 := pop(s1)
	fmt.Printf("pop: %v\n", s2)
}
