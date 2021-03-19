package main

import "fmt"

var _ in = (*person)(nil) // 用于约束结构体必须实现该接口。若没有，则在编译时报错
var _ in = (*base)(nil)

func main() {
	testPerson()
}

type base struct {
	name string
	age  int32
}

func (u base) getAge() int32 {
	return u.age
}

type person struct {
	rank int32
	base
}

func (u person) getAge() int32 {
	return u.age + 10
}

func (u *base) set(pram base) {
	*u = pram
}

func (u *base) getName() string {
	return u.name
}

//func (u *base) getAge() int32 {
//	return u.age
//}

type in interface {
	set(pram base)
	getName() string
	getAge() int32
}

func testPerson() {
	u := new(base)
	//p := &person{base: *u}

	//testInterface(p)
	testInterface(u)
}

func testInterface(i in) {
	i.set(base{"name", 20})
	fmt.Println("i:", i)
	fmt.Println("i.name:", i.getName())
	fmt.Println("i.age:", i.getAge())
}
