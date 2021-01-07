package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "aaSASA ddd dsjkdsjs dk"
	fmt.Println("替换前: ", s)
	s1 := strings.Replace(s, "ASA", "abc", 1)
	fmt.Println("替换后: ", s1)
}
