package main

import (
	"fmt"
	"strconv"
)

//var (
//	size =1024
	//max_size =size*2
//)
func main() {
	var i int64
	i = 36
	b := strconv.Itoa(i)
	a := FormatInt(int64(i),26)
	j := fmt.Sprintf ("%d", i)
	fmt.Println(a, b, j)
	//s1 := []int {1, 2, 3}
	//s2 := []int {4, 5}
	//s1 = append(s1, s2...)
	//fmt.Println(size, max_size)
}
