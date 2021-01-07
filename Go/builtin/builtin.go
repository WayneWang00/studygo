package main

import (
	"errors"
	"fmt"
)

func main() {
	//testRecover()
	//testComplex()
	testCopy()
}

func testRecover() {
	var err error
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover:", err)
		}
	}()

	err = errors.New("test err")
	fmt.Println("print:", err)
	panic(err)
}

func testComplex() {
	com := complex(1, 2)
	i := imag(com)
	r := real(com)
	fmt.Println(com, i, r)
}

func testCopy() {
	a := make([]int32, 1)
	b := []int32{11, 2}
	c := copy(b, a)
	fmt.Println(a, b, c)
}
