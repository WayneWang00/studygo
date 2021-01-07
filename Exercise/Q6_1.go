package main

import "fmt"

func f(a, b int) (c, d int) {
	if a > b {
		c = b
		d = a
	} else {
		c = a
		d = b
	}
	return
}

func main() {
	a, b := f(2, 7)
	fmt.Printf("a: %d\nb: %d\n", a, b)
	fmt.Println(f(7, 2))
	fmt.Print(f(7, 2))
	//fmt.Println("f(2, 7): ", f(2, 7))
	//fmt.Println("f(7, 2): ", f(7, 2))
}
