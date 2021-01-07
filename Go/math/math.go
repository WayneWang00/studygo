package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	// 需要设置随机种子，如果不设置，每次的运行结果都一样
	rand.Seed(time.Now().UTC().UnixNano())
	f := rand.Float64()
	n := 4

	nFloat(f, n)
}

// 获取指定精度的小数
func nFloat(f float64, n int) {
	/*
		# 方法一：通过math函数实现
	*/
	pow10 := math.Pow10(n)
	// +0.5是为了四舍五入
	fmt.Println("pow10:", math.Trunc(f*pow10+0.5)/pow10)

	/*
		# 方法二：通过转字符串实现
	*/
	nf, err := strconv.ParseFloat(fmt.Sprintf("%.4f", f), 64)
	if err != nil {
		fmt.Println("parseFloat failed:", err)
		return
	}
	fmt.Println(nf)
}
