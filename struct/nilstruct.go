package main

import (
	"fmt"
	"reflect"
	"time"
	"unsafe"
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

	sizeOfStruct()
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

/*
	1：值类型
		bool、int、float、(unsafe)pointer、struct、array
	2：复合类型
		channel、map、func、pointer：一个指针，指向具体的数据。8字节
		string、interface：string一个指针，指向底层字节切片，一个整形表示长度；interface两个指针，一个指向类型，一个指向具体数据。16字节
		slice：一个指针，指向底层数组，两个整形分别表示长度和容量。24字节
*/
func sizeOfStruct() {
	fmt.Println(unsafe.Sizeof([]int{}))        // 24
	fmt.Println(unsafe.Sizeof(interface{}(0))) // 16
	fmt.Println(unsafe.Sizeof(string("")))     // 16
	fmt.Println(unsafe.Sizeof(struct{}{}))     // 0
	fmt.Println(unsafe.Sizeof(false))          // 1
	fmt.Println(unsafe.Sizeof(int(0)))         // 8
	fmt.Println(unsafe.Sizeof(float32(0)))     // 4
	fmt.Println(unsafe.Sizeof([1]int{}))       // 8

	// 内存对齐
	fmt.Println(unsafe.Sizeof(struct {
		a bool
		b string
		c bool
	}{})) // 32 bool=1 string=16 第一和第三都为1，分别内存补齐7为8，所以为 1+7 + 16 + 1+7=32
	fmt.Println(unsafe.Sizeof(struct {
		a bool
		b bool
		c string
	}{})) // 24 bool=1 string=16 第一为1，第二接着第一为2，第二补齐6为8，所以为 1 + 1+6 + 16=24
}
