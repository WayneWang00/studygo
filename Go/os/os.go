package main

import (
	"fmt"
	"os"
)

func main() {
	//exit()
	args()
}

// 让当前的程序以给定的状态码code退出，0 成功，非0 出错
func exit() {
	//os.Exit(0)
	//os.Exit(1)
	os.Exit(-1)
}

// 第一个参数为执行程序路径
func args() {
	for i, v := range os.Args {
		fmt.Printf("参数：%d 值：%s\n", i, v)
	}
}
