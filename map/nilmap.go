package main

import "fmt"

func main() {
	//m := make(map[string]int)
	//m["name"] = 1
	//m["a"] = 2
	//fmt.Println("map value: ", m)
	//fmt.Println("nil befor: ", &m)
	//m = nil
	//fmt.Println("nil after: ", &m)
	//m = make(map[string]int)
	//fmt.Println("make: ", &m)
	//if m["a"] == 0 {
	//	fmt.Println("m[a]: ", m["a"])
	//	m["a"]++
	//	fmt.Println("m[a]: ", m["a"])
	//}
	//fmt.Println("map value: ", m)
	testNilVar()
}

func testNilVar() {
	m := make(map[int32]string)
	m[123] = "name1"
	m[456] = "name2"
	fmt.Printf("m:%v\n", m)
	fmt.Printf("m before Type:%T\n", m)
	m = nil
	fmt.Println("m len:", len(m))
	fmt.Printf("m after Type:%T\n", m)
	m = make(map[int32]string)
	m[123] = "name1"
	fmt.Printf("m:%v\n", m)
	fmt.Printf("m end Type:%T\n", m)
}
