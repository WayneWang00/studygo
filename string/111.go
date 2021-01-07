package main

import (
	"fmt"
	"unsafe"
)

func main(){
		x := 2
		y := 4

		table := make([][]int, x)//
		for a  := range table {
			table[a] = make([]int, y)
		}
		fmt.Println(table)
		str := "string"
		fmt.Println("改变前：", str, &str)
		str = "String"
		fmt.Println("改变后：", str, &str)
	i := [][]int {{1,2,3,4},{2,3,4},{5},{0},{3}}
	k := make (map[int]int)
	k[0] = 0
	k[1] = 1
	k[2] = 2
	k[3] = 3
	k[4] = 4
	k[10] = 10
	var j =make ([][]int, 5,5)
	fmt.Println("二维切片为:", i)
	fmt.Println("二维切片的字节数：",unsafe.Sizeof(i))
	fmt.Println("二维切片:", j)
	fmt.Println("map: ",k)
	fmt.Println("map的长度为：", len(k))
	fmt.Println(len(j),cap(j))
	fmt.Println(len(i[0]), cap(i[1]))
}
