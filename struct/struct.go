package main

import "fmt"

func main() {
	setStu()
	ad := admin{user{name: "张三", email: "zhangsan@flysnow.org"}, "管理员"}
	fmt.Println(ad)
	sayHello(ad.user)
	sayHello(ad)
	//a := user{"name", "email", 20}
	//b := admin{a, "a"}
	//fmt.Println("a.age: ", a.age)
	//fmt.Printf("a: %+v\n", a)
	//fmt.Println("a.hello:")
	//a.hello()
	//fmt.Println("b.hello:")
	//b.hello()
	//fmt.Println("b.name:", b.name)
}

type user struct {
	name  string
	email string
	age   int
}

type admin struct {
	user
	level string
}

type Hello interface {
	hello()
}

func (u user) hello() {
	fmt.Println("hello, i am a user")
}

func (a admin) hello() {
	fmt.Println("hello, i am a admin")
}

func sayHello(h Hello) {
	h.hello()
}

type stu1 struct {
	a string
	b int
}

type stu2 struct {
	stu1
	b int
}

func (s *stu1) add() {
	fmt.Println("stu1.add")
	s.set() // 只会调用stu1的set方法
}

func (s *stu1) set() {
	s.b = 20
}

func (s *stu2) set() {
	s.b = 15
}

func setStu() {
	stu := stu2{stu1{"a", 5}, 10}
	stu.set() // 当stu2没有set方法时，修改的还是stu1中b的值
	fmt.Println("stu:", stu)
	stu.stu1.set()
	fmt.Println("stu:", stu)

	stu.add()
	fmt.Println("stu:", stu)
}
