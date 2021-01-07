package main

import "fmt"

type e interface{}

func main() {
	a := []e{1, 2, 3, 4, 5, 6}
	str := []e{"a", "b", "c", "d"}
	A := Map(mult, a)
	Str := Map(mult, str)
	fmt.Printf("int: %v\n", A)
	fmt.Printf("string: %v\n", Str)
}

func mult(f e) e {
	switch f.(type) {
	case int:
		return f.(int) * 2
	case string:
		return f.(string) + f.(string)
	}
	return f
}

func Map(f func(e) e, a []e) []e {
	b := make([]e, len(a))
	for k, v := range a {
		b[k] = f(v)
	}
	return b
}
