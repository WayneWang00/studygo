package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"syscall"
	"unsafe"
)

var kernel = syscall.NewLazyDLL("Kernel32.dll")

type memoryStatusEx struct {
	cbSize                  uint32
	dwMemoryLoad            uint32
	ullTotalPhys            uint64 // in bytes
	ullAvailPhys            uint64
	ullTotalPageFile        uint64
	ullAvailPageFile        uint64
	ullTotalVirtual         uint64
	ullAvailVirtual         uint64
	ullAvailExtendedVirtual uint64
}

func getMemoryInfo() {

	GlobalMemoryStatusEx := kernel.NewProc("GlobalMemoryStatusEx")
	var memInfo memoryStatusEx
	memInfo.cbSize = uint32(unsafe.Sizeof(memInfo))
	mem, _, _ := GlobalMemoryStatusEx.Call(uintptr(unsafe.Pointer(&memInfo)))
	if mem == 0 {
		return
	}
	fmt.Printf("total=:", memInfo.ullTotalPhys)
	fmt.Printf("free=:", memInfo.ullAvailPhys)
}

func main() {
	getMemoryInfo()
}

func Encode(data interface{}) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

//解码时如下，data为需要解码的字节数组，to为相应的接收结构体，记住to的结构体结构应与被编码的data相一致，解码后内容保存在to里面，直接使用to即可
func Decode(data []byte, to interface{}) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	return dec.Decode(to)
}

//使用的时候： b, err := Encode(data) if err != nil { //错误处理 } if err := Decode(b, &to); err != nil { //错误处理}
