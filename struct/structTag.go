package main

import (
	"fmt"
	"reflect"
)

type Info struct {
	Name string `tag:"name"`
	Age  int32  `tag:"age"`
}

func main() {
	// 测试获取结构体tag对应的值
	var a = Info{
		Name: "Wayne",
		Age:  20,
	}
	getStructTag(a)
	//a := new(Info)
	//b := *a
	//fmt.Println("b value: ", b)
}

func getStructTag(info Info) {
	a := reflect.ValueOf(info)
	if !a.IsValid() {
		fmt.Printf("%v is not valid\n", a)
	}
	fmt.Println("a.Kind value: ", a.Kind())
	if a.Kind() == reflect.Ptr && !a.IsNil() {
		a.Elem()
	}
	var typ = a.Type()

	var m = make(map[string]interface{}, typ.NumField())
	for i := 0; i < typ.NumField(); i++ {
		tag := typ.Field(i).Tag.Get("tag")
		if tag == "" || tag == "_" {
			continue
		}
		switch a.Field(i).Kind() {
		case reflect.String:
			m[tag] = a.Field(i).String()
		case reflect.Int32:
			m[tag] = a.Field(i).Int()
		}
	}
	fmt.Printf("m value: %v\n", m)
}
