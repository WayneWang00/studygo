package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	a := 1
	var p1 person     //分配了person的值给p1，数据类型为person。
	p2 := new(person) //分配了内存并将指针赋值给p2，数据类型为*person。
	fmt.Printf("p1: %v\n", p1)
	fmt.Printf("p2: %v\n", p2)
	Set(&a)
	Set1(a)
	fmt.Printf("a: %v\n", &a)
}

func Set(t *int) {
	x := t //x指向了t指向的内容，就是实际上的参数指向的内容
	fmt.Printf("Set: %v\n", x)
}

func Set1(t int) {
	x := &t //x指向了一个新的变量t，其中包含了实参的副本
	fmt.Printf("Set1: %v\n", x)
}
