package main

import "fmt"

func main() {
	var i = 0
	j := i
	for ; i < j+10; i++ {
		fmt.Println("i的值为: ", i)
	}
}
