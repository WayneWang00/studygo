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
	fmt.Printf("len=%d, cap=%d\n", len(slice), cap(slice)) // len=5, cap=8 	扩容操作：旧容量*2=6，大于append的容量5，所以新容量为6。int32占4字节*6=24字节，查表发现24不符合向上取整得到32，再除以4得到8，即最终容量。
}

func testSlice5() {
	var slice1 = []int32{1, 2, 3, 4, 5}
	var slice2 = slice1[2:4]
	fmt.Printf("改变前：slice1=%+v, slice2=%+v\n", slice1, slice2) // 改变前：slice1=[1 2 3 4 5], slice2=[3 4]
	fmt.Println("slice2:", len(slice2), cap(slice2))           // len:2 cap:3 	slice2截取slice1，所以底层数组和slice1共用，不过slice2是从截取的位置开始计算，即从3到开始。所以长度为2，容量为剩余数组的长度3。
	slice3 := append(slice2, 6, 7, 8)
	slice3[0] = 123
	fmt.Printf("扩容：slice1=%+v, slice2=%+v, slice3=%+v\n", slice1, slice2, slice3) // 扩容：slice1=[1 2 3 4 5], slice2=[3 4], slice3=[123 4 6 7 8]
	fmt.Println("slice3:", len(slice3), cap(slice3))                              // len:5 cap:8 	slice3在slice2上增加三个元素，长度为5，超过了容量3。所以会进行扩容操作，具体看第83行。
	slice4 := append(slice2, 9)
	slice4[0] = 234
	fmt.Printf("未扩容：slice1=%+v, slice2=%+v, slice4=%+v\n", slice1, slice2, slice4) // 未扩容：slice1=[1 2 234 4 9], slice2=[234 4], slice4=[234 4 9]
	fmt.Println("slice4:", len(slice4), cap(slice4))                               // len:3 cap:3 	slice4在slice2上增加1一个元素，长度为3，没有超过容量3。所以不扩容
}

/*
	扩容操作：
		1.判断append的容量大于旧容量的两倍，就以append的容量作为新容量。反之，就判断旧容量是否大于1024，如果大于，则按照旧容量的1.25倍为新容量；如果不大于，则按照旧容量的2倍为新容量。
		2.通过第一步产生的容量，乘以当前slice类型占的字节数，得到总字节数。然后进行内存对齐操作，即根据得到的总字节数查阅允许的字节数表（一般是8的偶数倍），向上取整得到新的总字节数。再用新总字节数除以当前slice类型的字节数，得到最终的容量。
*/
func testSlice6() {
	var slice = []int64{1, 2, 3}
	fmt.Printf("before: len=%+v, cap=%+v\n", len(slice), cap(slice)) // before: len=3, cap=3
	slice2 := append(slice, 3)
	fmt.Printf("append 1: len=%+v, cap=%+v\n", len(slice2), cap(slice2)) // append 1: len=3, cap=6
	slice3 := append(slice, 3, 4, 5)
	fmt.Printf("append 3: len=%+v, cap=%+v\n", len(slice3), cap(slice3)) // append 3: len=6, cap=6
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
	testSlice5()
	testSlice6()
	//testMap()
	//testInterface()
	//testNil()
	lenAndDisplace()
}

const s = "aabbccdde"

// https://zhuanlan.zhihu.com/p/74543420
func lenAndDisplace() {
	// 因为s[:]不是常量，所以 len(s[:]) 的结果不是常量，则 1<<len(s[:]) 结果的类型与接收者类型一样为byte。但是byte最大为255，而 1<<len(s[:]) 结果为512，超出了范围所以结果为0，最后 /128 的结果为0。
	var n1 byte = 1 << len(s[:])
	fmt.Println("n1:", n1)
	// 因为s为常量，所以 len(s) 的结果也为常量，则 1<<len(s) 结果的类型就为整数常量，所以结果就为512，最后 /128 的结果为4。
	var n2 byte = 1 << len(s) / 128
	fmt.Println("n2:", n2)
}

func testNil() {
	var a = new(*testStr)
	fmt.Println("a:", *a)
}

type testStr struct {
	name string
	age  int32
}

type struct1 = struct {
	testStr
}
