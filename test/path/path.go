package main

import (
	"fmt"
	"os"
)

func main() {
	dirPath, err := os.Getwd()
	if err != nil {
		fmt.Println("获取当前目录失败")
		return
	}
	fmt.Println(dirPath)
}
