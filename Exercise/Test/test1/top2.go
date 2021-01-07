package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	cmd := exec.Command("top", "-bn 1")
	var out bytes.Buffer
	cmd.Stdout = &out //输出
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
	var a []string
	var b []string
	for i := 0; i < 5; i++ {
		line, _ := out.ReadString('\n')
		if i == 3 || i == 4 {
			fmt.Println(line)
			a = strings.Split(line, ",")
			for i = 0; i < len(a); i++ {
				fmt.Println(a[i])
				b = strings.Split(a[i], " ")
				fmt.Println(b[0])
			}
		}
	}
}
