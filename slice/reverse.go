package main

import "fmt"

func main() {
	reverse1()
	reverse2()
}

func reverse1() {
	var a = []int{1, 2, 3, 4, 5, 6, 7, 1, 2, 3, 4, 5, 6}

	for i := len(a)/2 - 1; i >= 0; i-- {
		j := len(a) - 1 - i
		a[i], a[j] = a[j], a[i]
	}

	fmt.Println("reverse1:", a)
}

func reverse2() {
	var a = []int{7, 6, 5, 4, 3, 2, 1, 7, 6, 5, 4, 3, 2}

	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}

	fmt.Println("reverse2:", a)
}
