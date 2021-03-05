package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//dosmt1(nil)
	//dosmt2(0)
	//testAppend()
	arr2Slice()
	//rangeSlice()
	//towDimensionSlice()
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

func rangeSlice() {
	var slice1 = []int{1, 2, 3, 4, 5}

	////使用下标访问slice
	//for i := 0; i < 5; i++ {
	//	fmt.Printf("slice1[%d] = %d", i, slice1[i])
	//}
	//fmt.Println()
	//
	////使用range进行遍历
	//for i, v := range slice1 {
	//	fmt.Printf("slice1[%d] = %d", i, v)
	//}

	//a := 6
	//var count int
	//for _, v := range slice1 {
	//	if a == v {
	//		break
	//	}
	//	count++
	//}
	//if count == len(slice1) {
	//	fmt.Println("count:", count)
	//}

	for k := range slice1 {
		fmt.Println(k, ":", slice1[k])
	}
}

func towDimensionSlice() {
	x, y := 2, 4
	table := make([][]int, x)
	for k := range table {
		table[k] = make([]int, y)
	}
	fmt.Printf("%v\n", table)

	i := [][]int{{1, 2, 3, 4}, {2, 3, 4}, {4, 5}, {6}, {0}, {5}}
	var j = make([][]int, 5, 5)
	fmt.Println("二维切片i：", i)
	fmt.Println("二维切片i字节数：", unsafe.Sizeof(i))
	fmt.Println("二维切片j：", j)
	fmt.Println(len(j), cap(j))
	fmt.Println(len(i[0]), cap(i[1]))
}
