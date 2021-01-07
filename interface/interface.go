package main

import "fmt"

func main() {
	var a interface{}
	a = int64(16)
	a = "test"
	fmt.Println(a)
}
