package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	name string
	age  int
}

type Person2 struct {
	Name string
	age  int
}

func set(i interface{}) {
	switch i.(type) {
	case *Person:
		r := reflect.ValueOf(i)
		r.Elem().Field(0).SetString("Albert Einstein")
	case *Person2:
		r := reflect.ValueOf(i)
		r.Elem().Field(0).SetString("Albert Einstein")
	}
}

func main() {
	p3 := new(Person2)
	p2 := new(Person)
	p := Person{name: "Name", age: 12}
	p1 := Person2{Name: "Name", age: 12}
	set(p) //当调用一个非指针参数时，变量的复制的。所以，反射是在副本上。这样不能改变原来的值，仅仅改变副本的。
	fmt.Printf("Person: %v\n", p)
	set(p1)
	fmt.Printf("Person1: %v\n", p1)
	set(p2) //运行时会报错
	fmt.Printf("Person2:%v\n", p2)
	set(p3)
	fmt.Printf("Person3: %v\n", p3)
}
