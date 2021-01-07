package main

import (
	"fmt"
	"strings"
)

func isSlash1(r rune) bool {
	return r == '\\' || r == '/'
}

func main() {
	s := "C:\\Windows\\System32\\FileName"
	//s := "abcd"
	ss := strings.FieldsFunc(s, isSlash1)
	fmt.Printf("%q\n", ss)
}
