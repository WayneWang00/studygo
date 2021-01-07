package main

import (
	"fmt"
	"strconv"
)

func String(s []int) {
	var str string
	for _, v := range s {
		str = str + strconv.Itoa(v)
	}
	fmt.Printf("My stack: %v\n", str)
}

func main() {
	s := []int{1, 2, 3}
	String(s)
	smap := make(map[int]int)
	for i, v := range s {
		smap[i] = v
	}
	fmt.Printf("smap: %v\n", smap)
}
