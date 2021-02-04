package main

import "fmt"

func main() {
	nilKey()
	nilValue()
	//intKey()
}

func nilValue() {
	x := map[string]string{"one": "a", "two": "", "three": "c"}
	if v, ok := x["two"]; !ok {
		fmt.Println("v:", v)
		fmt.Println("no entry")
	} else {
		fmt.Println("v:", v)
		fmt.Println("exist")
	}
}

func nilKey() {
	var map1 = map[string]int{"key1": 100, "key2": 200}

	v, OK := map1["key1"]
	if OK {
		fmt.Println(v, OK)
	} else {
		fmt.Println(v)
	}
	// 这里 不是 :=，是 = ，因为这些变量已经定义过了。
	v, OK = map1["key3"]
	if OK {
		fmt.Println(v, OK)
	} else {
		fmt.Println(v)
	}
}

func intKey() {
	var a = map[int]string{1: "test", 2: "string"}
	fmt.Printf("a:%+v\n", a)
	fmt.Println("a[1]:", a[1])
	fmt.Println("a[2]:", a[2])
}
