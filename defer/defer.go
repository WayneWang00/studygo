package main

import "fmt"

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
	deferCall()
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
