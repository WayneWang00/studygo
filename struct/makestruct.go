package main

import "fmt"

type testMake struct {
	a int
	b string
	c bool
	d float64
}

func main() {
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
