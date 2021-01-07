package main

import "fmt"

func main() {
	var a = []int{0, 1, 2, 3, 4, 5, 6, 7}
	var s = make([]int, 6)

	//源长度为8，目标长度为6，只会复制前6个
	n1 := copy(s, a)
	fmt.Println("s: ", s)
	fmt.Println("n1: ", n1)

	//源长度为7，目标长度为6，复制索引1到6
	n2 := copy(s, a[1:])
	fmt.Println("s: ", s)
	fmt.Println("n2: ", n2)

	//源长度为8-5=3，只会复制3个值，目标中的后三个元素不会改变
	n3 := copy(s, a[5:])
	fmt.Println("s: ", s)
	fmt.Println("n3: ", n3)

	//将源中索引为5，6，7的元素复制到目标中的3，4，5中
	n4 := copy(s[3:], a[5:])
	fmt.Println("s: ", s)
	fmt.Println("n4: ", n4)
}
