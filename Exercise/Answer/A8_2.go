package main

import (
	"fmt"
	"strconv"
)

type stack struct {
	i    int
	data [10]int
}

func (s *stack) push(k int) {
	s.data[s.i] = k
	s.i++
	fmt.Printf("i: %d\n", s.i)
	fmt.Printf("stack: %v\n", s)
}

func (s stack) String() {
	var str string
	for i := 0; i < s.i; i++ {
		fmt.Println(str + "[" + strconv.Itoa(i) + ":" + strconv.Itoa(s.data[i]) + "]")
	}
}

func (s stack) PrintStr() string {
	var str string
	for i := 0; i < s.i; i++ {
		str = str + "[" + strconv.Itoa(i) + ":" + strconv.Itoa(s.data[i]) + "]"
	}
	return str
}

func main() {
	var s stack
	s.push(25)
	s.push(14)
	fmt.Printf("push: %v\n", s)
	//fmt.Printf("String: %s\n", s.String())
	s.String()
	fmt.Printf("PrintStr :%v\n", s.PrintStr())
}
