package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, e := os.Open("src/Wayne/bufio/l.txt")
	if e != nil {
		fmt.Println("读取文件错误")
	}
	r := bufio.NewReader(f)
	b := make([]byte, 1024)
	n, er := r.Read(b)
	if er != nil {
		fmt.Println("read err")
	}
	fmt.Println(string(b), n)
}
