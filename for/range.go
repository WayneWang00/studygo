package main

import (
	"fmt"
	"time"
)

func main() {
	//pointerRange()
	rangeGo()
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
