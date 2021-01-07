package main

import (
	"encoding/json"
	"fmt"
)

type str struct {
	V  string `json:"v"`
	V2 int    `json:"v2"`
	V3 string `json:"v3"`
	V4 []test
}

type test struct {
	T string
}

func main() {
	v := `{"v":"123456","v3":"2345","v2":1}`
	//v := `{"v":"123456","v3":"2345","v2":1}`
	r := str{}
	json.Unmarshal([]byte(v), &r)
	fmt.Println("r[V][0] value: ", r.V[0])

	data := str{
		V:  "v",
		V2: 2,
		V3: "v3",
		V4: make([]test, 0),
	}
	databyte, err := json.Marshal(data)
	if err != nil {
		fmt.Println("err: ", err)
	}
	fmt.Println("test length: ", len(data.V4))
	fmt.Println("data: ", string(databyte))
}
