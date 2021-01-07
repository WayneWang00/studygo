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
