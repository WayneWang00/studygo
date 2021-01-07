package main

import "fmt"

func main() {
	a := `"string" , "字符串文字"\n`//``中\没有特殊含义
	b := "string\n"
	c := "str"
	fmt.Print(a)
	fmt.Print(b)
	fmt.Print(c)
}
