package main

import "fmt"

func main() {
	//dosmt1(nil)
	//dosmt2(0)
	//testAppend()
	//arr2Slice()
	s := make([]int32, 0)
	for _, v := range s {
		fmt.Println(v)
	}
}

func dosmt1(is []int) {
	is = append(is, 2)
	fmt.Println(is)
}

func dosmt2(it int) {
	fmt.Println(it)
}

func testAppend() {
	arr := [4]int{10, 20, 30, 40}
	slice := arr[1:3]
	testSlice1 := slice
	testSlice2 := append(append(append(slice, 1), 2), 3)
	slice[0] = 11
	fmt.Println("len:", len(testSlice1), "cap: ", cap(testSlice1))
	fmt.Println("testSlice1[0]:", testSlice1[0])
	fmt.Println("len: ", len(testSlice2), "cap: ", cap(testSlice2))
	fmt.Println("testSlice2[0]:", testSlice2[0])
}

func arr2Slice() {
	arr := [5]int{1, 2, 3, 4, 5}
	s := arr[1:4]
	fmt.Println("len:", len(s), "cap:", cap(s))
	a1 := [2]int{1, 2}
	a2 := [2]int{1, 3}
	fmt.Println("a1==a2:", a1 == a2)
}
