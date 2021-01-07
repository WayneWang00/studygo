package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//testGameConf()
	//testJson()
	testStu()
}

type ListConf struct {
	Vip           int32 `json:"vip"`
	BankruptType  int32 `json:"bankrupt_type"`
	Number        int32 `json:"number"`
	GrandMoney    int64 `json:"grand_money"`
	BankruptMoney int64 `json:"bankrupt_money"`
}

type GameConf struct {
	GameId int32      `json:"game_id"`
	List   []ListConf `json:"list"`
}

func testGameConf() {
	gameConf := `[{"game_id": 1, "list": [{"vip": 1, "bankrupt_type": 1, "number": 5, "grand_money": 3000, "bankrupt_money": 1500},{"vip": 0, "bankrupt_type": 1, "number": 3, "grand_money": 1500, "bankrupt_money": 1000}] }]`
	var newGameConf = make([]GameConf, 0)
	if err := json.Unmarshal([]byte(gameConf), &newGameConf); err != nil {
		fmt.Println("Unmarshal Failed: ", err)
		return
	}
	fmt.Printf("Unmarshal Success: %+v\n", newGameConf)
}

type AutoStu struct {
	Age   int    `json:"age"`
	Name  string `json:"name"`
	Child []int  `json:"child"`
}

func testJson() {
	json1 := `{"age": 10, "name": "Tom", "child": [1, 2, 3]}`
	stu := AutoStu{}
	if err := json.Unmarshal([]byte(json1), &stu); err != nil {
		fmt.Println("json1 unmarshal failed:", err)
		return
	}
	fmt.Printf("json1:%+v\t len:%d\t cap:%d\t\n", stu, len(stu.Child), cap(stu.Child))
	a := stu.Child
	fmt.Printf("a:%+v\t len:%d\t cap:%d\t\n", a, len(a), cap(a))
	fmt.Printf("json1:%p\t a:%p\t\n", stu.Child, a)
	fmt.Printf("&json1:%p\t &a:%p\t\n", &stu.Child, &a)

	//json2 := `{"age": 12, "name": "Wayne", "child": [4, 5, 6, 7, 8, 9]}`
	json2 := `{"name": "Wayne", "child": [4, 5, 6, 7, 8, 9]}`
	if err := json.Unmarshal([]byte(json2), &stu); err != nil {
		fmt.Printf("json2:%+v\n", stu)
		return
	}
	fmt.Printf("json2:%+v\t len:%d\t cap:%d\t\n", stu, len(stu.Child), cap(stu.Child))
	fmt.Printf("a:%+v\t len:%d\t cap:%d\t\n", a, len(a), cap(a))
	fmt.Printf("json2:%p\t a:%p\t\n", stu.Child, a)
	fmt.Printf("&json2:%p\t &a:%p\t\n", &stu.Child, &a)
}

type TestStu struct {
	Name string   `json:"name"`
	Conf ListConf `json:"conf"`
}

func testStu() {
	testJson := `{"name": "Jack", "conf": {"vip": 1, "bankrupt_type": 1, "number": 5, "grand_money": 3000, "bankrupt_money": 1500}}`
	a := TestStu{}
	if err := json.Unmarshal([]byte(testJson), &a); err != nil {
		fmt.Println("unmarshal failed:", err)
		return
	}
	fmt.Printf("a:%+v\n", a)
}
