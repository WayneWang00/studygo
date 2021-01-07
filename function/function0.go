package main

import "fmt"

type stu struct {
	f1  func()
}

func (s1 *stu) f2() {
	fmt.Println("f2")
}

var f1 =func() {
	fmt.Println("f1")
}

func main() {
	s := stu{}
	s.f1 = s.f2
	fmt.Println(s.f1)
	fmt.Println(s.f2)
}
