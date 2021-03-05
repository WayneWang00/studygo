package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	usuallyMath()
	//nFloat()
}

// 获取指定精度的小数
func nFloat() {
	// 需要设置随机种子，如果不设置，每次的运行结果都一样
	rand.Seed(time.Now().UTC().UnixNano())
	f := rand.Float64()
	n := 4

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

// 常用数学方法
func usuallyMath() {
	var n = -145

	fmt.Println("绝对值：", math.Abs(float64(n)))          // 取绝对值
	fmt.Println("向上取整：", math.Ceil(5.12))              // 向上取整
	fmt.Println("向下取整：", math.Floor(5.78))             // 向下取整
	fmt.Println("取最大值：", math.Max(-5, 6.78))           // 取最大值
	fmt.Println("取最小值：", math.Min(-3, 1.23))           // 取最小值
	fmt.Println("取x-y和0中的最大值：", math.Dim(-3.45, 1.23)) // 取x-y和0中的最大值
	fmt.Println("取余数：", math.Mod(13, 4))               // 取余数
	fmt.Println("取整数：", math.Trunc(-3.45))             // 取整数
	fmt.Print("取整数、小数：")
	fmt.Println(math.Modf(-5.678))                                 // 取整数、小数
	fmt.Println("float32的IEE 754二进制表示：", math.Float32bits(1.2345)) // float32的IEE 754二进制表示对应的4字节uint
	fmt.Println("uint32的IEE 754二进制表示：", math.Float32frombits(2))   // uint32的IEE 754二进制表示对应的4字节float
	fmt.Println("x的y次方：", math.Pow(4, 3))                          // x的y次方
	fmt.Println("10的n次方：", math.Pow10(4))                          // 10的n次方
	fmt.Println("指数：", math.Exp2(3))                               // 指数
	fmt.Println("开平方：", math.Sqrt(16))                             // 开平方
	fmt.Println("开立方：", math.Cbrt(8))                              // 开立方
	fmt.Println("圆周率：", math.Pi)                                   // 圆周率
}
