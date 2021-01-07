package main

import "fmt"

func main() {
	var a *int
	b := 0
	a = &b
	fmt.Println("a: ", *a)
}
