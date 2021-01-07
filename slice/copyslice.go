package main

import "fmt"

func main() {
	//var slice1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//var slice2 = make([]int, 3, 5)
	//var n int
	//n = copy(slice2, slice1) // 只复制了3个元素
	//fmt.Println(n, slice2, len(slice2), cap(slice2))
	//
	//slice3 := slice1[3:6]         //二者引用同一个底层数组
	//n = copy(slice3, slice1[1:5]) //所以，copy的时候发生元素重叠
	//fmt.Println(n, slice1, slice3)

	copySlice()
}

func copySlice() {
	var s1 []int
	s2 := []int{1, 2, 3}
	s3 := []int{4, 5, 6}
	s1 = append(s2, s3...)
	fmt.Printf("s1:%+v\n", s1)
	fmt.Printf("s2:%+v\n", s2)
	fmt.Printf("s3:%+v\n", s3)
}
