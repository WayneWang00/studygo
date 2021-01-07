package main

import "fmt"

func sprintln(a ...int) {
	for _, v := range a {
		fmt.Printf("print: %d\n", v)
	}
}

func main() {
	sprintln(1, 2, 3, 4)
	sprintln(1, 5, 5, 6)
}
