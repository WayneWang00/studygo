package main

import "fmt"

func main() {
	//appendSlice()
	copySlice()
}

func appendElement() {
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

	s := []int{1, 2}
	fmt.Printf("len=%d, cap=%d\n", len(s), cap(s))
	s1 := append(s, 4, 5, 6)
	fmt.Printf("len=%d, cap=%d\n", len(s1), cap(s1))
	s2 := append(s, 4, 5, 6, 7)
	fmt.Printf("len=%d, cap=%d", len(s2), cap(s2))
}

func appendSlice() {
	var s1 []int
	s2 := []int{1, 2, 3}
	s3 := []int{4, 5, 6}
	s1 = append(s2, s3...)
	fmt.Printf("s1:%+v, len=%d, cap=%d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2:%+v, len=%d, cap=%d\n", s2, len(s2), cap(s2))
	fmt.Printf("s3:%+v, len=%d, cap=%d\n", s3, len(s3), cap(s3))
}

func copySlice() {
	var slice1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var slice2 = make([]int, 3, 5)
	var n int
	n = copy(slice2, slice1) // 只复制了3个元素
	fmt.Println(n, slice2, len(slice2), cap(slice2))

	slice3 := slice1[3:6]         //二者引用同一个底层数组
	n = copy(slice3, slice1[1:5]) //所以，copy的时候发生元素重叠
	fmt.Println(n, slice1, slice3)
}
