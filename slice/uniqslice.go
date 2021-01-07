package main

import "fmt"

//type info struct {
//	count int
//	value string
//}

func main() {
	a := []int{1, 2, 3, 4, 1, 2, 3, 4}
	b := make(map[int]int)
	for i := 0; i < len(a); i++ {
		c := 1
		if _, ok := b[a[i]]; ok {
			continue
		}
		for j := i + 1; j < len(a); j++ {
			if a[i] == a[j] {
				c++
			}
		}
		b[a[i]] = c
	}
	fmt.Printf("b value: %+v", b)
}
