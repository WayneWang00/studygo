package main

import "fmt"

func Map(f func(int) int, a []int) []int {
	b := make([]int, len(a))
	for k, v := range a {
		b[k] = f(v)
	}
	return b
}

func Map2(f func(string) string, a []string) []string {
	b := make([]string, len(a))
	for k, v := range a {
		b[k] = f(v)
	}
	return b
}

func F(a int) (b int) {
	b = a * a
	return
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	f := func(i int) int {
		return i * i
	}
	fmt.Printf("Map: %v\n", Map(f, a))
	b := []string{"a", "b", "c", "d", "e"}
	g := func(i string) string {
		return i + i
	}
	fmt.Printf("Map2: %v\n", Map2(g, b))
	fmt.Printf("FMap: %v\n", Map(F, a))
}
