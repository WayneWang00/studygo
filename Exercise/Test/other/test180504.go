package main

import "fmt"

var a = 0

func main() {
	f()
	fmt.Println(a)
	fmt.Println(&a)
}
func f() {
	a = a + 1
	fmt.Println(a)
	fmt.Println(&a)
}
