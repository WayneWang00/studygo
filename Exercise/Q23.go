package main

import "fmt"

/*
Q23:
    在第 72 页的代码 5.3 编译正常——就像文中开始描述的那样。但是当运行的时
    候，会得到运行时错误，因此有些东西有错误。为什么代码编译没有问题呢？

代码能够编译是因为整数类型实现了空接口，这是在编译时检查的。
正确的方法时测试这个空接口是否可以被转换，如果可以就调用相应的方法。如下方法g2()。
*/

type I interface {
	Get() int
	Put(int)
}

func g(something interface{}) int {
	return something.(I).Get()
}

func g2(something interface{}) int {
	if v, ok := something.(I); ok {
		return v.Get()
	}
	return -1
}

func main() {
	i := 5
	//fmt.Println(g(i))
	fmt.Println(g2(i))
}
