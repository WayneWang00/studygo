package main

import (
	"fmt"
	"unsafe"
)

type testMake struct {
	a int
	b string
	c bool
	d float64
}

func main() {
	makeStu()
	//sizeStu()
}

func makeStu() {
	//a := new(testMake)
	//b := testMake{
	//	a: 1,
	//	b: "test",
	//	c: true,
	//	d: 1.000,
	//}
	//a = &b
	//if (testMake{}) == *a {
	//	fmt.Println("a is null")
	//}
	//fmt.Println(a)
	//var c *testMake
	//c = &b
	//fmt.Println("c value: ", c)
	b := testMake{
		a: 1,
		b: "test",
		c: true,
		d: 1.000,
	}
	a := make([][]testMake, 2)
	a[0] = append(a[0], b)
	fmt.Println("a value: ", a)
}

type stu struct { // 所有字段对齐后再进行struct对齐，整个struct的对齐值为系统默认对齐值和所包含的字段的最大对齐值中取最小值
	b  bool   // 占一个字节，对齐值为1
	i  int32  // 占四个字节，对齐值为4
	s  string // 占十六个字节，对齐值为系统默认对齐值和16中取最小值
	bt byte   // 占一个字节，对齐值为1
}

func sizeStu() {
	s := stu{}

	fmt.Println("size:", unsafe.Sizeof(s)) // 1+3+4+16+1+7 = 32
}
