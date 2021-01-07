package main

import "fmt"

func main() {
	var i = 0
LOOP:
	if i < 10 {
		fmt.Println("i的值为: ", i)
		i++
		goto LOOP
	}
}
