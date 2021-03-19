package main

import (
	"fmt"
	"time"
)

func main() {
	//pointerRange()
	//rangeGo()
	appendSlice()
}

func pointerRange() {
	var value = []*struct{ number int }{{1}, {2}, {3}}
	for _, v := range value {
		v.number *= 10
	}

	fmt.Println(value[0], value[1], value[2])
}

func rangeGo() {
	var value = []string{"one", "two", "three"}
	for _, v := range value {
		fmt.Println("v:", v)
		go func(s string) {
			fmt.Println("go s:", s)
		}(v)

		go func() {
			fmt.Println("go v:", v)
		}()
	}

	time.Sleep(3 * time.Second)
}

func appendSlice() {
	var value = []int{1, 2, 3, 4}
	for k := range value { // 在range时对value进行了拷贝，所以range过程中对value的修改不会反映到range中
		value = append(value, k)
	}

	fmt.Println("value:", value)
}
