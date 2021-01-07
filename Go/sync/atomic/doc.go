package main

import (
	"fmt"
	"sync/atomic"
	"unsafe"
)

func main() {
	swapInt32()
	swapInt64()
	swapUint32()
	swapUint64()
	swapUintptr()
	swapPointer()
}

// 原子性保存新的int32，返回旧的int32
func swapInt32() {
	var a int32 = -1
	oldAddr := &a
	var newAddr int32 = -2
	old := atomic.SwapInt32(oldAddr, newAddr)
	fmt.Println("oldInt32:", old)
}

// 原子性保存新的int64，返回旧的int64
func swapInt64() {
	var a int64 = -1
	oldAddr := &a
	var newAddr int64 = -2
	old := atomic.SwapInt64(oldAddr, newAddr)
	fmt.Println("oldInt64:", old)
}

// 原子性保存新的uint32，返回旧的uint32
func swapUint32() {
	var a uint32 = 1
	oldAddr := &a
	var newAddr uint32 = 2
	old := atomic.SwapUint32(oldAddr, newAddr)
	fmt.Println("oldUint32:", old)
}

// 原子性保存新的uint64，返回旧的uint64
func swapUint64() {
	var a uint64 = 1
	oldAddr := &a
	var newAddr uint64 = 2
	old := atomic.SwapUint64(oldAddr, newAddr)
	fmt.Println("oldUint64:", old)
}

// 原子性保存新的uintptr，返回旧的uintptr
func swapUintptr() {
	var a uint = 1
	oldAddr := uintptr(unsafe.Pointer(&a))
	var b uint = 2
	newAddr := uintptr(unsafe.Pointer(&b))
	old := atomic.SwapUintptr(&oldAddr, newAddr)
	fmt.Println("oldUintptr:", old)
}

// 原子性保存新的unsafe.Pointer，返回旧的unsafe.Pointer
func swapPointer() {
	var a int32 = 1
	oldAddr := unsafe.Pointer(&a)
	fmt.Println("oldAddr:", oldAddr)
	var b int32 = 2
	newAddr := unsafe.Pointer(&b)
	//old := (*int64)(atomic.SwapPointer(&oldAddr, newAddr))
	old := atomic.SwapPointer(&oldAddr, newAddr)
	fmt.Println("oldPointer:", old)
}
