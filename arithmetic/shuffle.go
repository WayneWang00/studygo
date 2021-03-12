package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//FisherYates()
	//Perm()
	KnuthDurstenfeld()
}

func FisherYates() {
	var value = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	rand.Seed(time.Now().UnixNano()) // 没有Seed(),每次运行程序得到的都为相同的伪随机序列
	rand.Shuffle(len(value), func(i, j int) {
		value[i], value[j] = value[j], value[i]
	})

	fmt.Printf("value:%+v\n", value)
}

func Perm() {
	var value = []int{1, 2, 3, 4, 5, 6}
	n := len(value)
	newVal := make([]int, n)
	rand.Seed(time.Now().UnixNano())
	rr := rand.Perm(n)

	for i := 0; i < n; i++ {
		newVal[i] = value[rr[i]]
	}

	fmt.Printf("newVal:%+v\n", newVal)
}

func KnuthDurstenfeld() {
	var value = []int{1, 2, 3, 4, 5, 6, 7, 8}
	n := len(value)
	rand.Seed(time.Now().UnixNano())
	for i := n; i > 0; i-- {
		last := i - 1
		l := rand.Intn(i)
		value[last], value[l] = value[l], value[last]
	}

	fmt.Printf("value:%+v\n", value)
}
