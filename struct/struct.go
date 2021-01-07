package main

import "fmt"

func main() {
	//ad := admin{user{"张三", "zhangsan@flysnow.org"}, "管理员"}
	//fmt.Println(ad)
	//sayHello(ad.user)
	//sayHello(ad)
	a := user{}
	fmt.Println("a.age: ", a.age)
	fmt.Printf("a: %+v\n", a)
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
