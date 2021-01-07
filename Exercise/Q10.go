package main

import (
	"fmt"
)

func fibonacci(a int) []int {
	x := make([]int, a)
	x[0], x[1] = 1, 1
	for i := 2; i < a; i++ {
		x[i] = x[i-1] + x[i-2]
	}
	return x
}

func main() {
	for _, v := range fibonacci(10) {
		fmt.Printf("fibonacci数: %d\n", v)
	}
}
