package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
)

func main() {
	//testRead()
	//testWrite()
	testSize()
}

// 从r中读出字节数据并反序列化成结构数据
func testRead() {
	var pi float64
	b := []byte{0x18, 0x2d, 0x44, 0x54, 0xfb, 0x21, 0x09, 0x40}
	buf := bytes.NewReader(b)
	err := binary.Read(buf, binary.LittleEndian, &pi)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
		return
	}
	fmt.Println(pi)
}

// 将数据序列化成字节流写入w中
func testWrite() {
	buf := new(bytes.Buffer)
	var pi = math.Pi
	err := binary.Write(buf, binary.LittleEndian, pi)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
		return
	}
	fmt.Printf("% x\n", buf.Bytes()) // % x 在字节之间用空格隔开
}

// 返回数据序列化后的字节长度
func testSize() {
	var a int8 = 1
	p := &a
	b := [10]int64{1}
	s := "abcd"
	bs := make([]byte, 20)

	fmt.Println(binary.Size(a))
	fmt.Println(binary.Size(p))
	fmt.Println(binary.Size(b))
	fmt.Println(binary.Size(s))
	fmt.Println(binary.Size(bs))
}
