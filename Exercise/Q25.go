package main

import "fmt"

func less(l, r interface{}) bool {
	switch l.(type) {
	case int:
		if _, ok := r.(int); ok {
			return l.(int) < r.(int)
		}
	case float32:
		if _, ok := r.(float32); ok {
			return l.(float32) < r.(float32)
		}
	}
	return false
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := []float32{1.2, 3.5, 5.2, 3.1, 4.2}
	maxa := 0
	var maxb float32 = 0.0
	for i := 0; i < len(a); i++ {
		if less(maxa, a[i]) {
			maxa = a[i]
		}
	}
	fmt.Printf("a中的最大值为: %d\n", maxa)
	for i := 0; i < len(b); i++ {
		if less(maxb, b[i]) {
			maxb = b[i]
		}
	}
	fmt.Printf("b中的最大值为:%f\n", maxb)
}
