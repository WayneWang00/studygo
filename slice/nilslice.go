package main

import (
	"fmt"
	"strconv"
)

func main() {
	i := 4
	ret := make([]string, i)
	fmt.Println("befor: ", ret)
	ret = append(ret, "test")
	fmt.Println("after: ", ret)
	fmt.Println("len: ", len(ret))
	fmt.Println("ret[4]", ret[4])

	ret1 := make([]string, i)
	for j := 0; j < i; j++ {
		ret1[j] = "test" + strconv.Itoa(j)
	}
	fmt.Println("ret1 value: ", ret1)

	s := make([]int32, 0)
	for _, v := range s {
		fmt.Println(v)
	}
}
