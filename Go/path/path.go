package main

import (
	"fmt"
	"path"
)

func main() {
	//match()
	//clean()
	//split()
	//join()
	//ext()
	//base()
	//isAbs()
	dir()
}

//模式匹配字符串，返回真
func match() {
	fmt.Println(path.Match("abc", "abc"))
	fmt.Println(path.Match("a*", "abc"))
	fmt.Println(path.Match("a*/b", "a/c/b"))
}

//通过词法操作返回和原始路径代表同一地址的最短路径
func clean() {
	paths := []string{
		"a/c",
		"a//c",
		"a/c/.",
		"a/c/b/..",
		"/../a/c",
		"/../a/b/../././/c",
		"/",
		"",
	}

	for _, p := range paths {
		fmt.Printf("Clean(%q) = %q\n", p, path.Clean(p))
	}
}

//将路径从最后一个斜杠后面分隔成两部分，即目录和文件名
func split() {
	fmt.Println(path.Split("static/myfile.css"))
	fmt.Println(path.Split("myfile.css"))
	fmt.Println(path.Split(""))
}

//将任意数量的路径元素放入到一个路径中，且会添加'/'
func join() {
	fmt.Println(path.Join("a", "b", "c"))
	fmt.Println(path.Join("a", "b/c"))
	fmt.Println(path.Join("a/b", "c"))
	fmt.Println(path.Join("", ""))
	fmt.Println(path.Join("a", ""))
	fmt.Println(path.Join("", "a"))
}

//返回路径最后一个元素的文件扩展名（包含'.'）
func ext() {
	fmt.Println(path.Ext("/a/b/c/bar.css"))
	fmt.Println(path.Ext("/"))
	fmt.Println(path.Ext(""))
}

//返回路径的最后一个元素
func base() {
	fmt.Println(path.Base("/a/b"))
	fmt.Println(path.Base("/"))
	fmt.Println(path.Base(""))
}

//判断是否为一个绝对路径
func isAbs() {
	fmt.Println(path.IsAbs("/dev/null"))
}

//返回除去最后一个路径元素部分的路径
func dir() {
	fmt.Println(path.Dir("/a/b/c"))
	fmt.Println(path.Dir("a/b/c"))
	fmt.Println(path.Dir("/a/"))
	fmt.Println(path.Dir("a/"))
	fmt.Println(path.Dir("/"))
	fmt.Println(path.Dir("///abc"))
	fmt.Println(path.Dir(""))
}
