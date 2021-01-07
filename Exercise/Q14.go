package main

import "fmt"

func plustwo() func(int) int {
	f := func(b int) int {
		return b + 2
	}
	return f
}

func plusx(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}

func main() {
	p := plustwo()
	fmt.Printf("plustwo: %d\n", p(2))
	px := plusx(3)
	fmt.Printf("plusx: %d\n", px(2))
}
