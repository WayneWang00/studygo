package main

import "fmt"

type people struct {
	name string
	age  int
}

func main() {
	a := [3]int{1, 2, 3}
	b := []int{1, 2, 3}
	c := map[string]int{"first": 1, "second": 2, "third": 3}
	d := people{"aaa", 10}

	fmt.Println("a change before: ", a)
	a1 := changeArr(a)
	fmt.Println("a change: ", a1)
	fmt.Println("a change after: ", a)

	fmt.Println("b change before: ", b)
	b1 := changeSlice(b)
	fmt.Println("b change: ", b1)
	fmt.Println("b change after: ", b)

	fmt.Println("c change before: ", c)
	c1 := changeMap(c)
	fmt.Println("c change: ", c1)
	fmt.Println("c change after: ", c)

	fmt.Println("d change before: ", d)
	d1 := changeStruct(d)
	fmt.Println("d change: ", d1)
	fmt.Println("d change after: ", d)
}

func changeArr(data [3]int) [3]int {
	data[0] = 4
	return data
}

func changeSlice(data []int) []int {
	data[0] = 4
	return data
}

func changeMap(data map[string]int) map[string]int {
	data["first"] = 4
	return data
}

func changeStruct(data people) people {
	data.name = "bbb"
	return data
}
