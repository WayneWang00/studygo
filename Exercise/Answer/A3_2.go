package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "aaSASA ddd dsjkdsjs dk"
	fmt.Printf("str= %s\nLength= %d\nRune= %d", str, len([]byte(str)), utf8.RuneCount([]byte(str)))
}
