package main

import (
	"encoding/json"
	"fmt"
)

type user1 struct {
	Name_first string `json:"name"`
	Age_num    int    `json:"age"`
}

var users []user1

func main() {
	jsonstr := `[{"name":"Bob","age":30},{"name":"Tom","age":40}]`
	jsonstr1 := `{"name":"Bob","age":30}`
	b := user1{
		Name_first: "name",
		Age_num:    20,
	}
	var a user1
	//var b users
	json.Unmarshal([]byte(jsonstr), &users)
	json.Unmarshal([]byte(jsonstr1), &a)
	fmt.Println(a)
	//log.Fatal("test: ", a)
	fmt.Println(len(users))
	for i, _ := range users {
		fmt.Println("name: ", users[i].Name_first, "age: ", users[i].Age_num)
	}
	users = append(users, b)
	for i, _ := range users {
		fmt.Println("name: ", users[i].Name_first, "age: ", users[i].Age_num)
	}
	var c uint8
	c = 0X07
	fmt.Println(c)
	d := 4
	if d != 3 && d != 5 {
		fmt.Println("d: ", d)
	}
}
