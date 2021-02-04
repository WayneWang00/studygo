package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	rr = rand.New(rand.NewSource(time.Now().UnixNano()))
	a1 = [2]int{}
	a2 = [10]int{}
)

func main() {
	now := time.Now()
	a := []int{0, 1}
	for i := 0; i < 10000; i++ {
		randslice()
		//fmt.Println(rr.Intn(10))
		sliceOutOfOrder(a)
	}
	fmt.Println(a1)
	fmt.Println(a2)
	shuffle1()
	val := []int{10, 12, 14, 16, 18, 20}
	fmt.Println(shuffle(val))
	fmt.Println("time:", time.Now().Sub(now))
}
func sliceOutOfOrder(in []int) []int {
	l := len(in)
	for i := l - 1; i > 0; i-- {
		r := rr.Intn(i)
		in[r], in[i] = in[i], in[r]
	}
	a1[in[0]] += 1
	return in
}

func randslice() {
	in := rr.Perm(10)
	a2[in[0]] += 1
}

//不知道切片的大小和容量，且不修改值传递的值
func shuffle(vals []int) []int {
	l := len(vals)
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for i := l - 1; i > 0; i-- {
		r := r.Intn(i)
		vals[r], vals[i] = vals[i], vals[r]
	}
	return vals
}

//无需创建新的数组或者切片来实现洗牌算法
func shuffle1() {
	vals := []int{10, 12, 14, 16, 18, 20}
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for _, i := range r.Perm(len(vals)) {
		val := vals[i]
		fmt.Println(val)
	}
}
