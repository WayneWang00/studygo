package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	testRegexp()
}

func testRegexp() {
	s := "i LoVe You , xiaorui.cc 13684390231 ,i like tornado lang, golang 正则 . Shell zsh golang bash c+ java jsp php"
	fmt.Println("string:", s)

	// 转化为小写
	re, err := regexp.Compile(`[0-9a-zA-Z]*`)
	if err != nil {
		fmt.Println("compile failed:", err)
		return
	}
	src := re.ReplaceAllStringFunc(s, strings.ToLower)
	fmt.Println("to lower:", src)

	// 替换所有数字
	re, err = regexp.Compile(`[0-9]*`)
	if err != nil {
		fmt.Println("compile failed:", err)
		return
	}
	src = re.ReplaceAllString(s, "")
	fmt.Println("to number:", src)

	// 判断是否为电话号码
	re, err = regexp.Compile(`1[0-9]{10,}`)
	if err != nil {
		fmt.Println("compile failed:", err)
		return
	}
	b := re.MatchString(s)
	fmt.Println("is phone number:", b)

	// 返回匹配的位置
	re, err = regexp.Compile(`xiao`)
	if err != nil {
		fmt.Println("compile failed:", err)
		return
	}
	ins := re.FindStringIndex(s)
	fmt.Printf("str index:%+v\n", ins)

	in := strings.Index(s, "xiao")
	fmt.Println("string index:", in)

	re, err = regexp.Compile(`正则`)
	if err != nil {
		fmt.Println("chance compile failed:", err)
		return
	}
	ins = re.FindIndex([]byte(s))
	fmt.Printf("index:%+v\n", ins)

	// 返回所有匹配
	re, err = regexp.Compile(`zsh\s(\w+)\sbash`)
	if err != nil {
		fmt.Println("compile failed:", err)
		return
	}
	ss := re.FindSubmatch([]byte(s))
	for i, v := range ss {
		fmt.Println(i, ":", v)
	}
}
