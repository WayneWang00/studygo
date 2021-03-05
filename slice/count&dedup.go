package main

import (
	"fmt"
)

func main() {
	//count()
	deduplication()
}

func count() {
	a := []int{1, 2, 3, 4, 4, 3, 2, 1}
	b := make(map[int]int)

	for i := 0; i < len(a); i++ {
		if v, ok := b[a[i]]; ok {
			b[a[i]] = v + 1
		} else {
			b[a[i]] = 1
		}
	}

	fmt.Printf("b value:%+v\n", b)
}

func deduplication() {
	var a = []int{1, 2, 3, 4, 5, 4, 3, 2, 1}
	b := make(map[int]struct{})
	var c []int

	for i := 0; i < len(a); i++ {
		if _, ok := b[a[i]]; ok {
			continue
		} else {
			c = append(c, a[i])
			b[a[i]] = struct{}{}
		}
	}

	fmt.Printf("a:%+v, b:%+v, c:%+v\n", a, b, c)
}
