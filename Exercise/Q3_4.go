package main

import "fmt"

func main() {
	s := "foobar"
	fmt.Println("s转换前: ", s)
	sbyte := []byte(s)
	fmt.Println("sbyte转换前: ", sbyte)
	for i, _ := range sbyte {
		if i < len(sbyte)-1-i {
			sb := sbyte[len(sbyte)-1-i]
			sbyte[len(sbyte)-1-i] = sbyte[i]
			sbyte[i] = sb
		}
	}
	fmt.Println("sbyte转换后: ", sbyte)
	s = string(sbyte)
	fmt.Println("s转换后: ", s)
}
