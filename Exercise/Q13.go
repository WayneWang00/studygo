package main

import "fmt"

//可以不用返回排序后的切片，因为切片是引用型的
func bubble(a []int) []int {
	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[i] {
				a[j], a[i] = a[i], a[j]
			}
		}
	}
	fmt.Println(&a[0])
	return a
}

func main() {
	a := []int{2, 4, 6, 1, 3, 5}
	fmt.Printf("bubbule: %v\n", bubble(a))
	fmt.Println(&a[0])
}
