package main

import "fmt"

func max(a []int) (b int) {
	for _, v := range a {
		if v > b {
			v, b = b, v
		}
	}
	return
}

func min(a []int) (b int) {
	b = a[0]
	for _, v := range a {
		if v < b {
			b = v
		}
	}
	return
}

func main() {
	slice := []int{1, 2, 3, 5, 8, 4, 6}
	fmt.Printf("Max: %d\n", max(slice))
	fmt.Printf("Min: %d\n", min(slice))
}
