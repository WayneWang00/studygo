package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
)

/*
 * 用于提供消息完整性保护。
 * 发送一份原数据以及通过原数据进行md5加密的校验和。接收方接收到数据后通过md5加密生成新的校验和，然后和接收到的校验和进行比较。相同则表示收到的数据完整，反之则不完整。
 */

type person struct {
	Name string
	Age  int32
}

func main() {
	a := make([]person, 3)
	fmt.Printf("%+v\n", a)

	testSum()
	fmt.Println("----------")
	testNew()
}

func testSum() {
	p := person{Name: "test", Age: 18}
	b, err := json.Marshal(p)
	if err != nil {
		fmt.Println("marshal failed:", err)
		return
	}
	fmt.Println("marshal b:", string(b))
	sum := md5.Sum(b)
	fmt.Printf("md5 sum: %#x\n", sum)
}

func testNew() {
	h := md5.New()
	io.WriteString(h, "the first")
	io.WriteString(h, "the second")
	sum := h.Sum(nil)
	fmt.Printf("%#x\n", sum)

	data := hex.EncodeToString(sum)
	fmt.Printf("%+v\n", data)
}
