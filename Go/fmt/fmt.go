package main

import "fmt"

func main() {
	//general()
	//boolean()
	integer()
	//floatAndComplex()
	//stringAndByte()
	//sliceAndPoint()
	//otherFlag()
}

type testStruct struct {
	name string
	age  int32
	sex  string
}

var in interface{}

// 通用
func general() {
	a := new(testStruct)
	fmt.Printf("%v\n", *a)  // 值的默认格式
	fmt.Printf("%+v\n", *a) // 类似%v,为结构体是会输出字段名
	fmt.Printf("%#v\n", *a) // 值的go语法表示
	fmt.Printf("%T\n", *a)  // 值类型的go语法表示
	fmt.Printf("%%\n")      // 百分号

	b := make(map[string]interface{})
	b["1"] = *a
	b["2"] = "test"
	fmt.Printf("%v\n", b)
	fmt.Printf("%+v\n", b)
	fmt.Printf("%#v\n", b)
	fmt.Printf("%T\n", b)

	in = b
	fmt.Printf("%v\n", in)
	fmt.Printf("%+v\n", in)
	fmt.Printf("%#v\n", in)
	fmt.Printf("%T\n", in)
}

// 布尔类型
func boolean() {
	var a bool
	fmt.Printf("%t\n", a) // 单词true或false
}

// 整数
func integer() {
	var b uint = 45
	fmt.Printf("b:%b\n", b)
	var c int = 45
	fmt.Printf("c:%b\n", c)
	var a = 45
	fmt.Printf("%b\n", a) // 二进制
	fmt.Printf("%c\n", a) // 对应的unicode码值
	fmt.Printf("%d\n", a) // 十进制
	fmt.Printf("%o\n", a) // 八进制
	fmt.Printf("%q\n", a) // 用单引号括起来的go语法字面值，必要时会采用安全的转义表示
	fmt.Printf("%x\n", a) // 十六进制（a-f）
	fmt.Printf("%X\n", a) // 十六进制（A-F）
	fmt.Printf("%U\n", a) // Unicode格式（U+1234）
}

// 浮点数和复数
func floatAndComplex() {
	var a = 123.456
	fmt.Println("--------------float--------------")
	fmt.Printf("%b\n", a) // 无小数部分、二进制指数的科学计数法
	fmt.Printf("%e\n", a) // 科学计数法（e）
	fmt.Printf("%E\n", a) // 科学计数法（E）
	fmt.Printf("%f\n", a) // 有小数部分，但无指数部分
	fmt.Printf("%F\n", a) // 有小数部分，但无指数部分
	fmt.Printf("%g\n", a) // 根据实际情况采用e%或f%格式（以获得更简洁、准确的输出）
	fmt.Printf("%G\n", a) // 根据实际情况采用E%或F%格式（以获得更简洁、准确的输出）

	c := 12.34 + 5.6i
	fmt.Println("---------------complex-------------")
	fmt.Printf("%b\n", c)
	fmt.Printf("%e\n", c)
	fmt.Printf("%E\n", c)
	fmt.Printf("%f\n", c)
	fmt.Printf("%F\n", c)
	fmt.Printf("%g\n", c)
	fmt.Printf("%G\n", c)
}

// 字符串和[]byte
func stringAndByte() {
	var s = "abcd"
	fmt.Println("------------string--------------")
	fmt.Printf("%s\n", s) // 直接输出字符串或[]byte
	fmt.Printf("%q\n", s) // 该值对应的双引号括起来的go语法字符串字面值
	fmt.Printf("%x\n", s) // 每个字节都用两字符十六进制表示（a-f）
	fmt.Printf("%X\n", s) // 每个字节都用两字符十六进制表示（A-F）

	bs := []byte(s)
	fmt.Println("------------[]byte-------------")
	fmt.Printf("%s\n", bs)
	fmt.Printf("%q\n", bs)
	fmt.Printf("%x\n", bs)
	fmt.Printf("%X\n", bs)
}

// slice和指针
func sliceAndPoint() {
	var s = "dcba"
	bs := []byte(s)
	fmt.Println("---------------slice-------------")
	fmt.Printf("%p\n", bs) // 表示为十六进制，并加上前导的0x

	var p = &s
	fmt.Println("---------------point-------------")
	fmt.Printf("%p\n", p)
}

// 其他flag
func otherFlag() {
	var a = 123
	var b = -12
	fmt.Printf("%+d\n", a) // 总是输出数值的正负号
	fmt.Printf("%+d\n", b)
	fmt.Printf("% d\n", a) // 对数值，正数前加空格，负数前加负号
	fmt.Printf("% d\n", b)
	fmt.Printf("%#o\n", a) // 切换格式：八进制数
	fmt.Printf("%#o\n", b)
	fmt.Printf("%#x\n", a) // 切换格式：十六进制数
	fmt.Printf("%#x\n", b)
	fmt.Printf("%#p\n", &a) // 指针，去掉前面的0x
	fmt.Printf("%#p\n", &b)
	fmt.Printf("%#U\n", a) // 输出Unicode格式，如符号可打印，还会输出空格加单引号括起来的符号
	fmt.Printf("%#U\n", b)
	fmt.Printf("%0d\n", a)
	fmt.Printf("%0d\n", b)

	var c = 123.4
	var d = -12.34
	fmt.Printf("%+f\n", c)
	fmt.Printf("%+f\n", d)

	var s = "ab/cd"
	var bs = []byte(s)
	fmt.Printf("%+q\n", s)
	fmt.Printf("%+q\n", bs)
	fmt.Printf("%#q\n", s) // 输出反引号括起来的未被转义的字符串
	fmt.Printf("%#q\n", bs)
	fmt.Printf("% x\n", s) // 给各打印字节间加空格
	fmt.Printf("% x\n", bs)
}
