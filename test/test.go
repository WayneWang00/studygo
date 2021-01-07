package main

import (
	"fmt"
	"unicode/utf8"
	"unsafe"
)

func testString() {
	var str = "hello, 世界"
	p := (*struct {
		str uintptr
		len int
	})(unsafe.Pointer(&str))
	fmt.Printf("%+v\n", p)                                         // &{str:4990566 len:13}
	fmt.Println("len(str):", len(str))                             // len(str): 13
	fmt.Println("RuneCountInString:", utf8.RuneCountInString(str)) // RuneCountInString: 9
	fmt.Println("rune:", len([]rune(str)))                         // rune: 9
}

func testString1() {
	var str = "hello, 世界"
	b := []byte(str)
	b[0] = 'a'
	str = string(b)
	fmt.Println("str:", str) // str: aello, 世界
	a := str[0]
	fmt.Println("a:", a)
}

func testSlice1() {
	var slice []int32 = make([]int32, 5, 10)
	p := (*struct {
		array uintptr
		len   int
		cap   int
	})(unsafe.Pointer(&slice))
	fmt.Printf("%+v\n", p) // &{array:824634114144 len:5 cap:10}
}

func testSlice2() {
	var array = [...]int32{1, 2, 3, 4, 5}
	var slice = array[2:4]
	fmt.Printf("改变前：array=%+v, slice=%+v\n", array, slice) // 改变前：array=[1 2 3 4 5], slice=[3 4]
	slice[0] = 123
	fmt.Printf("改变后：array=%+v, slice=%+v\n", array, slice) // 改变后：array=[1 2 123 4 5], slice=[123 4]
}

func testSlice3() {
	var slice1 = []int32{1, 2, 3, 4, 5}
	var slice2 = slice1[2:4]
	fmt.Printf("改变前：slice1=%+v, slice2=%+v\n", slice1, slice2) // 改变前：slice1=[1 2 3 4 5], slice2=[3 4]
	slice2[0] = 123
	fmt.Printf("改变后：slice1=%+v, slice2=%+v\n", slice1, slice2) // 改变后：slice1=[1 2 123 4 5], slice2=[123 4]
}

func testSlice4() {
	var array = [...]int32{1, 2, 3, 4, 5}
	var slice = array[2:4]
	fmt.Printf("改变前：array=%+v, slice=%+v\n", array, slice) // 改变前：array=[1 2 3 4 5], slice=[3 4]
	fmt.Printf("len=%d, cap=%d\n", len(slice), cap(slice)) // len=2, cap=3
	slice = append(slice, 6, 7, 8)
	fmt.Printf("改变后：array=%+v, slice=%+v\n", array, slice) // 改变后：array=[1 2 3 4 5], slice=[3 4 6 7 8]
	fmt.Printf("len=%d, cap=%d\n", len(slice), cap(slice)) // len=5, cap=8
}

func testSlice5() {
	var slice1 = []int32{1, 2, 3, 4, 5}
	var slice2 = slice1[2:4]
	fmt.Printf("改变前：slice1=%+v, slice2=%+v\n", slice1, slice2) // 改变前：slice1=[1 2 3 4 5], slice2=[3 4]
	//slice3 := append(slice2, 6, 7, 8)
	//slice3[0] = 123
	//fmt.Printf("扩容：slice1=%+v, slice2=%+v, slice3=%+v\n", slice1, slice2, slice3) // 扩容：slice1=[1 2 3 4 5], slice2=[3 4], slice3=[123 4 6 7 8]
	slice3 := append(slice2, 9)
	slice3[0] = 234
	fmt.Printf("未扩容：slice1=%+v, slice2=%+v, slice3=%+v\n", slice1, slice2, slice3) // 未扩容：slice1=[1 2 234 4 9], slice2=[234 4], slice3=[234 4 9]
}

func testSlice6() {
	var slice = []int64{1, 2, 3}
	fmt.Printf("before: len=%+v, cap=%+v\n", len(slice), cap(slice)) // before: len=3, cap=3
	//slice = append(slice, 3, 4, 5)
	slice = append(slice, 3)
	fmt.Printf("after: len=%+v, cap=%+v\n", len(slice), cap(slice)) // after: len=6, cap=8
}

func testMap() {
	var m = make(map[string]int32, 10)
	p := (*struct {
		count      int
		flags      uint32
		hash0      uint32
		B          uint8
		keysize    uint8
		valuesize  uint8
		bucketsize uint16
		buckets    uintptr
		oldbuckets uintptr
		nevacuate  uintptr
	})(unsafe.Pointer(&m))
	fmt.Printf("%+v\n", p) // &{count:824634122816 flags:5600384 hash0:0 B:0 keysize:0 valuesize:0 bucketsize:0 buckets:0 oldbuckets:0 nevacuate:0}
}

func testInterface() {
	var str interface{} = "Hello World!"
	p := (*struct {
		tab  uintptr
		data uintptr
	})(unsafe.Pointer(&str))
	fmt.Println("interface{}.tab:", unsafe.Sizeof(p.tab))      // interface{}.tab: 8
	fmt.Println("interface{}.data 大小:", unsafe.Sizeof(p.data)) // interface{}.data 大小: 8
	fmt.Printf("%+v\n", p)                                     // &{tab:4854304 data:5071792}
}

func main() {
	//testString()
	//testString1()
	//testSlice1()
	//testSlice2()
	//testSlice3()
	//testSlice4()
	//testSlice5()
	//testSlice6()
	//testMap()
	//testInterface()
	//testNil()
}

func testNil() {
	var a = new(*teststr)
	fmt.Println("a:", *a)
}

type teststr struct {
	name string
	age  int32
}

type struct1 = struct {
	teststr
}
