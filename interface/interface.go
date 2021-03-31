package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a interface{} // interface变量被赋值什么类型的值，该变量就是什么类型
	a = int64(16)
	a = "test"
	fmt.Println(reflect.TypeOf(a).Kind(), a)

	nilInterface()
}

type T interface{}

func nilInterface() {
	var (
		p1 T           // nil类型	值为nil
		p2 *T          // *T类型 	值为nil
		t1 interface{} = p1
		t2 interface{} = p2
	)

	fmt.Printf("t1:\nvalue:%+v\ttype:%T\np1:\nvalue:%+v\ttype:%T\n", t1, t1, p1, p1)
	fmt.Println(t1 == p1, t1 == nil) // ==或!= 判断的是interface的动态类型。 值可以通过反射来判断 reflect.ValueOf(a).isNil()
	fmt.Printf("t2:\nvalue:%+v\ttype:%T\np2:\nvalue:%+v\ttype:%T\n", t2, t2, p2, p2)
	fmt.Println(t2 == p2, t2 == nil) // true false

	var i interface{}
	var i2 interface{} = nil
	fmt.Println(i == nil, i2 == nil) // true true
}
