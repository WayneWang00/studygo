package main

import "fmt"

func main() {
	var f = [20]int{1, 1}
	for i := 2; i < 20; i++ {
		f[i] = f[i-1] + f[i-2]
	}

	for i := 0; i < 20; i++ { //采用下标进行遍历
		if i%5 == 0 {
			fmt.Printf("\n")
		}
		fmt.Printf("f[%2d] = %4d", i, f[i])
	}
	fmt.Println()

	for k, v := range f { //采用range关键字进行遍历
		fmt.Printf("f[%2d] = %4d", k, v)
	}
}
