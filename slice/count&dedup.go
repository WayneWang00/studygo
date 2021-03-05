package main

import (
	"fmt"
)

func main() {
	count()
	//deDuplication()
}

// 计数
func count() {
	a := []int{1, 2, 3, 4, 5, 4, 3, 2, 1}
	b := make(map[int]int)

	for _, item := range a {
		if v, ok := b[item]; ok {
			b[item] = v + 1
		} else {
			b[item] = 1
		}
	}

	fmt.Printf("b value:%+v\n", b)
}

// 去重
func deDuplication() {
	var a = []int{1, 2, 3, 4, 5, 4, 3, 2, 1}
	b := make(map[int]struct{})
	var c []int

	for i := 0; i < len(a); i++ {
		if _, ok := b[a[i]]; !ok {
			c = append(c, a[i])
			b[a[i]] = struct{}{}
		}
	}

	fmt.Printf("a:%+v, b:%+v, c:%+v\n", a, b, c)
}
