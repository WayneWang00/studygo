package main

import "fmt"

func main() {
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
