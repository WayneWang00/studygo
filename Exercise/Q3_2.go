package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	a := "aaSASA ddd dsjkdsjs dk"
	i := 0
	for _, b := range a {
		if b >= 'a' && b <= 'z' || b >= 'A' && b <= 'Z' {
			i++
		}
	}
	fmt.Println("字符串a中的字符个数为: ", i)
	fmt.Println("字符串a的字节数为：", len([]rune(a)))
	fmt.Println("字符串a的字节数为: ", utf8.RuneCountInString(a))
	num := strings.Count(a, "") - 1
	num1 := bytes.Count([]byte(a), nil) - 1
	fmt.Println("strings.Count(): ", num)
	fmt.Println("bytes.Count(): ", num1)
}
