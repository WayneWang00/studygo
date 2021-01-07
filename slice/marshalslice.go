package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type test struct {
	Id     int32    `json:"id"`
	People []people `json:"people"`
}
type people struct {
	Name string `json:"name"`
	Age  int32  `json:"age"`
}

func main() {
	ts := make([]test, 0)
	for i := 0; i < 3; i++ {
		t := newTest(i)
		ts = append(ts, t)
	}
	fmt.Printf("ts:%+v\n", ts)
	b, err := json.Marshal(ts)
	if err != nil {
		fmt.Println("marshal error:", err)
	}
	fmt.Println("marshal:", string(b))

	c := string(b)
	uts := make([]test, 0)
	err = json.Unmarshal([]byte(c), &uts)
	if err != nil {
		fmt.Println("unmarshal error:", err)
	}
	fmt.Printf("unmarshal:%+v\n", uts)
}

func newTest(i int) test {
	ps := make([]people, 0)
	for j := 0; j < 3; j++ {
		p := newPeople(j)
		ps = append(ps, p)
	}
	return test{
		Id:     int32(i + 1),
		People: ps,
	}
}

func newPeople(i int) people {
	return people{
		Name: "name" + strconv.Itoa(i+1),
		Age:  int32(i + 1),
	}
}
