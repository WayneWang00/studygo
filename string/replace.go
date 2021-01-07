package main

import (
	"fmt"
	"strconv"
	"strings"
)

var str = "aa,bb,cc,dd,ee,ff,ee,dd,cc,bb,aa"

func main() {
	//replace1()
	replace2()
}

func replace1() {
	arr := strings.Split(str, ",")
	var newStr string
	for k, v := range arr {
		if k == len(arr)-1 {
			newStr += v
			continue
		}
		n := k + 1
		newStr += v + strconv.Itoa(n)
	}
	fmt.Println(newStr)
}

func replace2() {
	n := strings.Count(str, ",")
	for i := 1; i <= n; i++ {
		if strings.Contains(str, ",") {
			str = strings.Replace(str, ",", strconv.Itoa(i), 1)
		}
	}
	fmt.Println(str)
}
