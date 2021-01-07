package main

import "fmt"

func main() {
	array := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	i := 0
LOOP:
	if i < len(array) {
		fmt.Printf("arry[%d]的值为: %d\n", i, array[i])
		i++
		goto LOOP
	}
	fmt.Println("数组的长度为: ", len(array))
	fmt.Println("i的值为: ", i)
}
