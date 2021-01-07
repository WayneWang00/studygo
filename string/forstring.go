package main

import "fmt"

func main() {
	var str = "hello 世界"
	fmt.Println("for:")
	for i := 0; i < len(str); i++ {
		fmt.Println(i, ":", str[i])
	}
	fmt.Println("for range:")
	for i, v := range str {
		fmt.Println(i, ":", v)
	}
	fmt.Println("for range i:")
	for i := range str {
		fmt.Println(i, ":", str[i])
	}
	p := []rune(str)
	fmt.Println("rune:", p)
	for i := 0; i < len(p); i++ {
		fmt.Println(i, ":", p[i])
	}
}
