package main

import (
	"fmt"
	"unsafe"
)

func main() {
	type St struct {
		A byte
		B int32
		C int64
	}
	st1 := St{}
	fmt.Println(unsafe.Sizeof(st1.A))
	fmt.Println(unsafe.Sizeof(st1.B))
	fmt.Println(unsafe.Sizeof(st1.C))
	fmt.Println(unsafe.Sizeof(st1))

	type Ft struct {
		A byte
		C int64
		B int32
	}
	ft1 := Ft{}
	fmt.Println(unsafe.Sizeof(ft1))

}
