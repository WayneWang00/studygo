package main

import "fmt"

func main() {
	////使用make创建切片
	//var slice1 = make([]int, 3, 6)
	////使用append添加元素，并且没有超过cap
	//slice2 := append(slice1, 1, 2, 3)
	////使用append添加元素，并且超出cap。这个时候底层数组会发生变化，新添加的元素自会添加到新的底层数组，不会覆盖旧的底层数组
	//slice3 := append(slice1, 4, 5, 6, 7)
	//slice1[0] = 10
	//fmt.Printf("slice1 = %p, len = %d, cap = %d %v \n", slice1, len(slice1), cap(slice1), slice1)
	//fmt.Printf("slice2 = %p, len = %d, cap = %d %v \n", slice2, len(slice2), cap(slice2), slice2)
	//fmt.Printf("slice3 = %p, len = %d, cap = %d %v \n", slice3, len(slice3), cap(slice3), slice3)
	testSlice()
}

func testSlice() {
	s := []int{1, 2}
	fmt.Printf("len=%d, cap=%d\n", len(s), cap(s))
	s1 := append(s, 4, 5, 6)
	fmt.Printf("len=%d, cap=%d\n", len(s1), cap(s1))
	s2 := append(s, 4, 5, 6, 7)
	fmt.Printf("len=%d, cap=%d", len(s2), cap(s2))
}
