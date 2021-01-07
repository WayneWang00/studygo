package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	//testBuffer()
	//testReader()
	//testCompare()
	testContains()
	var a, b = 1, 2
	fmt.Println(a &^ b)
}

func testCompare() {
	var a, b []byte
	a = []byte{1}
	b = []byte{2}
	// 返回一个按字典顺序比较两个字节切片的整数。a == b, 则为0；a < b, 则为-1；a > b，则为1
	if bytes.Compare(a, b) < 0 {
		fmt.Println("a < b")
	}

	a = []byte{2}
	b = []byte{2}
	if bytes.Compare(a, b) == 0 {
		fmt.Println("a == b")
	}

	a = []byte{3}
	b = []byte{2}
	if bytes.Compare(a, b) > 0 {
		fmt.Println("a > b")
	}

	// 判断两个字节切片是否相等
	if bytes.Equal(a, b) {
		fmt.Println("a == b")
	} else {
		fmt.Println("a not equal b")
	}
}

// 判断子字节切片是否在b中
func testContains() {
	var b1, b2 []byte
	b1 = []byte("abcfoo")
	b2 = []byte("abc")
	if bytes.Contains(b1, b2) {
		fmt.Println("b1 contains b2")
	}

	fmt.Println(bytes.Contains([]byte("seafood"), []byte("food")))
	fmt.Println(bytes.Contains([]byte("seafood"), []byte("bar")))
	fmt.Println(bytes.Contains([]byte("seafood"), []byte("")))
	fmt.Println(bytes.Contains([]byte(""), []byte("")))
}

func testReader() {
	b1 := []byte("hello world!!")
	read := bytes.NewReader(b1)
	fmt.Println("b1", read.Len())

	//b2 := make([]byte, 5)
	//n, err := read.Read(b2)
	//fmt.Println(n, err)
	//fmt.Println(string(b2))
	//fmt.Println("b1", read.Len())
	//fmt.Println("size", read.Size())

	b3 := make([]byte, 5)
	m, err := read.ReadAt(b3, 3)
	fmt.Println(m, err)
	fmt.Println(string(b3))
	fmt.Println("b1", read.Len())

	for {
		// 如果之前进行了read.Read操作，会导致ReadByte读取的数据只能从Read之后的元素开始。即：返回未被读取的byte
		b, err := read.ReadByte()
		if err == io.EOF {
			break
		}
		println(string(b))
	}

	b4 := []byte("hello 世界！！")
	read1 := bytes.NewReader(b4)
	for {
		b, n, err := read1.ReadRune()
		if err == io.EOF {
			break
		}
		println(string(b), n)
	}

	b5 := []byte("string builder")
	read1.Reset(b5)
	println(read1.Len())

	read2 := bytes.NewReader(b1)
	abs, err := read2.Seek(-2, 2)
	fmt.Println(abs, err)
	b, err := read2.ReadByte()
	fmt.Println(string(b), err)
}

func testBuffer() {
	b1 := []byte("hello world!!")
	buf := bytes.NewBuffer(b1)
	fmt.Println(buf.Len())
	fmt.Println(buf.Cap())

	buf.Grow(100)
	fmt.Println(buf.Len())
	fmt.Println(buf.Cap())

	b2 := make([]byte, 6)
	n, err := buf.Read(b2)
	fmt.Println(n, err)
	fmt.Println("b2", string(b2))

	b3 := buf.Next(5)
	fmt.Println("b3", string(b3))
	b4 := buf.Next(3)
	fmt.Println("b4", string(b4))

	buf1 := bytes.NewBuffer(b1)
	b5, err := buf1.ReadBytes(byte(' '))
	fmt.Println(len(b5))
	fmt.Println(string(b5))

	b6 := []byte("go programming")
	buf2 := bytes.NewBuffer(b1)
	buf2.Write(b6)
	fmt.Println(string(buf2.Bytes()))
	fmt.Println(buf2.String())
}
