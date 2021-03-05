package main

import "fmt"

func main() {
	fmt.Println(ANDIsOdd(11))
	fmt.Println(ANDCheckBit(15))
	XORSwap()
}

// 判断是否为奇数
func ANDIsOdd(i int) bool {
	return (i & 1) == 1
}

// 检查二级制位有多少个1
func ANDCheckBit(i int) int {
	var count int

	for i > 0 {
		count = count + (i & 1)
		i >>= 1
	}

	return count
}

// 两个数互换
func XORSwap() {
	var a, b = 5, 12
	fmt.Println("a:", a, "b:", b)

	a ^= b
	b ^= a
	a ^= b

	fmt.Println("a:", a, "b:", b)
}
