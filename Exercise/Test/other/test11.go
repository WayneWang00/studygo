package main

import (
	"fmt"
)

type inter interface{}
type str2 struct {
	b int
}
type str1 struct {
	a *int
}

func main() {
	var a int = nil
	var b float32 = nil
	var c string = nil
	var d bool = nil
	var e *int = nil
	var f func() = nil
	var g inter = nil
	var h []int = nil
	var i map[int]int = nil
	var j chan int = nil
	k := str1{nil}
	var l str1 = nil

	fmt.Println(a, b, c, d, e, f, g, h, i, j, k, l)
}
