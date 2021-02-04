package main

import (
	"encoding/json"
	"fmt"
)

type People struct {
	Name string `json:"name"`
	Age  int32  `json:"age"`
}

func main() {
	//map2Struct()
	updateStructField()
}

func map2Struct() {
	var m = make(map[string]interface{})
	m["Name"] = "A"
	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println("marshal failed:", err)
		return
	}
	var s = People{Age: 32}
	if err = json.Unmarshal(b, &s); err != nil {
		fmt.Println("unmarshal failed:", err)
		return
	}
	fmt.Printf("s:%+v\n", s)
}

type data struct {
	name string
}

func updateStructField() {
	var value = map[string]data{
		"X": {name: "Tom"},
	}
	//value["X"].name = "Jerry" // map中元素不可寻址 报错

	x := value["X"]
	x.name = "Jerry"
	value["X"] = x
	fmt.Println("value:", value)

	m := map[string]*data{
		"X": {name: "Tom"},
	}
	m["X"].name = "Jerry"
	fmt.Println("X:", m["X"])
}
