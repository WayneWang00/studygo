package main

import (
	"errors"
	"fmt"
)

func main() {
	//fmt.Println("a")
	//defer fmt.Println("b") //先进后出：a c e d b
	//f := "a" + "b"
	//fmt.Println(f)
	//fmt.Println("c")
	//defer fmt.Println("d")
	//fmt.Println("e")

	fmt.Println("f1:", f1())
	n := f2()
	fmt.Println("f2:", n, &n)
	//fmt.Println("f3:", f3())
	//fmt.Println("f4:", f4())
	//fmt.Println("testDefer1:", testDefer1())
	//fmt.Println("testDefer2:", testDefer2())

	//deferRecover()
	//deferCall()
	//deferFunc()
	fmt.Println(deferPanic())
}

// 1.r=0 2.r++ 3.return 	r=1
func f1() (r int) {
	defer func() {
		r++
	}()
	return
}

// 1.r=t 2.t=t+5 3.return 	r=5
func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
		fmt.Println("defer t:", t, &t)
	}()
	fmt.Println("defer before t:", t, &t)
	return t
}

// 1.r=1 2.r1=r1+5 3.return 	r=1
func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

func f4() (r int) {
	defer func(r *int) {
		*r = *r + 5
	}(&r)
	return 1
}

// 1.r=i 2.i++ 3.return 	r=0
func testDefer1() int {
	var i int
	defer func() {
		i++
	}()
	return i
}

// 1.r=0 2.r++ 3.return 	r=1
func testDefer2() (r int) {
	defer func() {
		r++
	}()
	return r
}

func deferRecover() {
	defer func() {
		//fmt.Println("recover:", recover()) // recover: test panic
		doRecover() // recover: <nil>
	}()

	panic("test panic")
}

func doRecover() {
	fmt.Println("recover:", recover())
}

func deferCall() {
	defer func() { fmt.Println(111) }()
	defer func() { fmt.Println(222) }()
	defer func() { fmt.Println(333) }()
	fmt.Println(444)
	fmt.Println(555)
	fmt.Println(666)
}

func deferFunc() {
	defer func() {
		fmt.Println("defer end")
	}()

	fmt.Println("defer start")
	deferCall()
}

func deferPanic() (err error) { // 按照下面数字依次输出 nil，b，c，a，c 	return有两步操作，defer在return第一步（即将err赋值为c）后执行，且defer里的命令在编译的时候就计算好了，因为err这是为nil所以1也为nil；而3中的err是和外层函数中的err是同一个地址，所以为c。
	defer func() {
		fmt.Println(err) // 3
		fmt.Println(&err)
		fmt.Println("a") // 4
	}()

	defer func(err error) {
		fmt.Println(err) // 1
		fmt.Println(&err)
		fmt.Println("b") // 2
	}(err)

	return errors.New("c") // 5
}
