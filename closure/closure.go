package main

import "fmt"

func main() {
	a := 1
	fmt.Println("new a:", &a)
	f := func() int {
		a++
		fmt.Println("&a:", &a)
		return a
	}
	fmt.Println("f1:", f())
	fmt.Println("f2:", f())
	i := incr()
	fmt.Println("i1:", i())
	fmt.Println("i2:", i())
	fmt.Println("i3:", i())
}

func incr() func() int {
	var x int
	return func() int {
		x++
		fmt.Println("&x:", &x)
		return x
	}
}
