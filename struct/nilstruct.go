package main

import (
	"fmt"
	"reflect"
	"time"
)

type A struct {
	Name  string
	Level int
}

type B struct {
	Skill string
}

func main() {
	//a := A{"Momo", 1}
	//b := B{"Starfall"}
	//clear(&a)
	//clear(&b)
	//fmt.Println("a: ", a)
	//fmt.Println("b: ", b)
	//
	//isStruct()
	//clearArray()
	//clearMap()

	a := new(A)
	b := a.Name
	c := a.Level
	fmt.Printf("a:%+v\n", a)
	fmt.Println("b:", b, "c:", c)

	//ret := retNilStruct()
	//if ret == nil {
	//	retChange := ret.toChange()
	//	fmt.Println("retChange value: ", retChange)
	//	fmt.Println("ret is nil")
	//}
	//if ret.Name == "" { // 会报错
	//	fmt.Println("ret.Name is nil")
	//}
}

//清空已经赋值的struct
func clear(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	fmt.Println("clear struct", reflect.Zero(p.Type()))
}

//判断是否为空的struct
type Session struct {
	playerId  string
	beehive   string
	timestamp time.Time
}

func isStruct() {
	session := Session{}
	if (Session{}) == session {
		fmt.Println("struct is zero value")
	}
	//或者判断其中一个值
	if session.playerId == "" {
		fmt.Println("struct is zero value")
	}
}

//清空array
func clearArray() {
	slice := []string{"first"}
	slice = slice[:0] //或者 slice = nil
	fmt.Println("array value: ", slice)
}

//清空map
func clearMap() {
	m := make(map[string]string)
	m["name"] = "yourname"
	m = nil
	fmt.Println("map value: ", m)
	//或者这样
	/*
		for k := range m {
			delete(m, k)
		}
	*/
}

func retNilStruct() *A {
	return nil
}

func (a *A) toChange() A {
	return A{
		Name:  a.Name, // 会报错
		Level: a.Level,
	}
}
