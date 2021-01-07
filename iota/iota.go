package main

import "fmt"

func main() {
	const (
		a = iota
		b
		c = "iota"
		d
		e = 10
		f
		g = iota
	)
	fmt.Println(a, b, c, d, e, f, g)

	testIota()
}

// 只有出现const，才会刷新iota
func testIota() {
	const (
		a = 1 + iota
		b
		c = 1 + iota
		d
	)
	fmt.Println(a, b, c, d)

	const (
		A = 1 + iota
		B
	)
	const (
		C = 1 + iota
		D
	)
	fmt.Println(A, B, C, D)
}
