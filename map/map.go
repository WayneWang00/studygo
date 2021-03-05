package main

import (
	"fmt"
	"unsafe"
)

func main() {
	nilMap()
	//nilKey()
	//nilValue()
	//intKey()
	//lenMap()
}

// 清空map
func nilMap() {
	m := map[int]string{
		123: "name1",
		456: "name2",
	}
	fmt.Printf("m:%v, &m:%p\n", m, &m)
	fmt.Printf("m before type:%T\n", m)

	m = nil
	fmt.Println("m len:", len(m), unsafe.Pointer(&m))
	fmt.Printf("m after type:%T\n", m)

	m = make(map[int]string) // 没有这条语句，下面会报错
	if _, ok := m[123]; !ok {
		m[123] = "name3"
	}
	fmt.Printf("m:%v, &m:%p\n", m, &m)
	fmt.Printf("m end type:%T\n", m)
}

// map中为nil的值
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

// map中不存在的key
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

func lenMap() {
	m := map[int]int{
		0:  0,
		1:  1,
		2:  2,
		3:  3,
		4:  4,
		10: 10,
	}

	fmt.Println("map m:", m)
	fmt.Println("map len:", len(m))
}
