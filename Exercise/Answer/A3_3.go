package main

import "fmt"

func main() {
	s := "aaSASA ddd dsjkdsjs dk"
	r := []rune(s)
	copy(r[4:4+3], []rune("abc"))
	fmt.Printf("Before: %s\n", s)
	fmt.Printf("After: %s", string(r))
}
