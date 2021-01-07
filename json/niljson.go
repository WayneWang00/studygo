package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type test struct {
	Name string `json:"Name"`
	Age  int32
}

type timeStruct struct {
	Id int32     `json:"id"`
	T  time.Time `json:"t"`
}

type notTag struct {
	A string
	B int32
}

func main() {
	//// 空slice字符串unmarshal不会报错
	//nilSlice()
	//
	//// time.Time类型通过json的marshal和unmarshal后类型变为string
	//timeJson()
	//
	//// struct类型数据转成map类型
	//struct2map()

	// map通过unmarshal将struct类型中的一个元素更新到struct中
	one2struct()

	//// struct进行json时字段首字母要大写，json里的字段名不区分大小写
	//notTagJson()
}

func nilSlice() {
	s := "[]"
	var p = make([]test, 0)
	err := json.Unmarshal([]byte(s), &p)
	fmt.Println("err:", err)
	fmt.Println("p:", p)
}

func timeJson() {
	t := time.Now()
	b, err := json.Marshal(t)
	if err != nil {
		fmt.Println("marshal failed:", err)
		return
	}
	fmt.Println("marshal data:", string(b))
	var end interface{}
	err = json.Unmarshal(b, &end)
	if err != nil {
		fmt.Println("unmarshal failed:", err)
		return
	}
	fmt.Printf("%+v\n", end)
}

func struct2map() {
	newTime := timeStruct{
		Id: 1,
		T:  time.Now(),
	}
	fmt.Printf("struct:%+v\n", newTime)
	fmt.Printf("%T\n", newTime.T)
	byt, err := json.Marshal(newTime)
	if err != nil {
		fmt.Println("marshal failed:", err)
		return
	}
	fmt.Println("marshal data:", string(byt))
	var endTime map[string]interface{}
	err = json.Unmarshal(byt, &endTime)
	if err != nil {
		fmt.Println("unmarshal failed:", err)
		return
	}
	fmt.Printf("map:%+v\n", endTime)
	fmt.Printf("%T\n", endTime["t"])
}

func one2struct() {
	var name = "new name"
	var newTest = test{
		Name: "name",
		Age:  10,
	}
	var value = map[string]interface{}{
		"name": name,
	}
	b, err := json.Marshal(value)
	if err != nil {
		fmt.Println("marshal failed:", err)
		return
	}
	fmt.Println("b:", string(b))
	err = json.Unmarshal(b, &newTest)
	if err != nil {
		fmt.Println("unmarshal failed:", err)
		return
	}
	fmt.Printf("newTest:%+v\n", newTest)
}

func notTagJson() {
	var a = map[string]interface{}{
		"A": "A",
		"B": int32(1),
	}
	b, err := json.Marshal(a)
	if err != nil {
		fmt.Println("marshal:", err)
		return
	}
	fmt.Println("b:", string(b))
	var c = new(notTag)
	err = json.Unmarshal(b, &c)
	if err != nil {
		fmt.Println("unmarshal:", err)
		return
	}
	fmt.Printf("c:%+v\n", c)
}
