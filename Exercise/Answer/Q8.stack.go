package main

import (
	"fmt"
	"strconv"
)

type STACK struct {
	i    int
	data [10]int
}

func main() {
	/*
		1. 创建一个固定大小保存整数的栈。它无须超出限制的增长。定义 push 函数——
		将数据放入栈，和 pop 函数——从栈中取得内容。栈应当是后进先出（LIFO）的。
	*/
	stack := new(STACK)
	fmt.Println("初始化:", stack)
	stack.push(11)
	fmt.Println("写入元素:11", stack)
	stack.push(12)
	fmt.Println("写入元素:12", stack)
	one := stack.pop()
	fmt.Println("取出一个元素:", one, stack)

	fmt.Println("-------------------------------------------------------------")

	/*
		2. 更进一步。编写一个 String 方法将栈转化为字符串形式的表达。
		可以这样的方式打印整个栈：fmt.Printf("My stack %v\n", stack)栈可以被输出成这样的形式：[0:m] [1:l] [2:k]
	*/
	stack.pop()
	stack.pop()
	stack.pop()
	stack.push(50)
	fmt.Println("打印堆栈:", stack.print())
	stack.push(51)
	fmt.Println("打印堆栈:", stack.print())
	stack.push(52)
	fmt.Println("打印堆栈:", stack.print())
	stack.push(53)
	fmt.Println("打印堆栈:", stack.print())
	stack.push(54)
	fmt.Println("打印堆栈:", stack.print())
}

func (s *STACK) push(k int) {
	if s.i+1 > 9 {
		return
	}
	s.data[s.i] = k
	s.i++
	return
}

func (s *STACK) pop() int {
	if s.i < 1 {
		s.i = 0
		return 0
	}
	s.i--
	return s.data[s.i]
}

func (s *STACK) print() (f string) {
	for i := 0; i < s.i; i++ {
		f += "[" + strconv.FormatInt(int64(i), 10) + ":" + strconv.FormatInt(int64(s.data[i]), 10) + "] "
	}
	return
}
