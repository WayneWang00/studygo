package main

import "fmt"
import "math"

func modify(array [5]int) {
	array[0] = 10
	fmt.Println("In modify(), array values:", array)
}
func Smodify(slice []int) {
	slice = append(slice, 10)
	fmt.Println("In Smodify(), slice values:", slice)
}

func main() {
	slice := make([]int, 3, 5)
	Smodify(slice)
	fmt.Println("In main(), slice values:", slice)
	array := [5]int{1, 2, 3, 4, 5}
	modify(array)
	fmt.Println("In main(), array values:", array)
	a, b, c := 1, 2.5, "Hello World,"
	const (
		d, e = 2, "你好，世界！"
		g    = 1 << 3
	)
	var f, h = "字符串", (1 != 0)
	var i float32 = 12
	j := 12.0
	b = float64(i)
	//func IsEqual(b,j,0.00001 float64) bool {
	//	return math.Abs(b-j) <0.00001
	//}
	//const g = 1 << 3
	fmt.Println("测试:")
	fmt.Println(a, b, c, d, e, f, g, h, i, j)
	fmt.Printf("c+e=%s\n", c+e)
	fmt.Println(math.Abs(b-j) < 0.00001)
	fmt.Println(math.Abs(j-d) < 0.00001)
	fmt.Println("换行\n")
	var k = 2
	fmt.Printf("a: %d, k: %d\n", a, k)
	a, k = k, a
	fmt.Printf("a: %d, k: %d\n", a, k)
}
