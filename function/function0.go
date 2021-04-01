package main

import "fmt"

func init() { // 同一个包不同文件中的init按照文件名的字典顺序初始化，同一个文件中的多个init按照前后顺序初始化
	fmt.Println(1)
}

func init() {
	fmt.Println(2)
}

func main() {
	//stuFunc()
	fmt.Println("hello,playground")
	a = 1
	fmt.Println(f) // 输出：f1 hello,playground f2
}

var a = 0

var f = f2() // 初始化顺序 1、引入包 2、变量 3、init 4、main

func f2() string {
	fmt.Println("f1")
	if a == 1 {
		return "f3"
	}
	return "f2"
}

type stu struct {
	f1 func()
}

func (s1 *stu) f2() {
	fmt.Println("f2")
}

var f1 = func() {
	fmt.Println("f1")
}

func stuFunc() {
	s := stu{}
	s.f1 = s.f2
	fmt.Println(s.f1)
	fmt.Println(s.f2)
	s.f1()
	s.f2()
}
