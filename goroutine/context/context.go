package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//testWithCancel()
	testWithValue()
}

func testWithCancel() {
	ctx, cancel := context.WithCancel(context.Background())
	go watchByName(ctx, "【协程1】：")
	go watchByName(ctx, "【协程2】：")
	go watchByName(ctx, "【协程3】：")
	//go func(ctx context.Context) {
	//	for {
	//		select {
	//		case <-ctx.Done():
	//			fmt.Println("监控退出，停止了...")
	//			return
	//		default:
	//			fmt.Println("goroutine监控中...")
	//			time.Sleep(2 * time.Second)
	//		}
	//	}
	//}(ctx)
	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	cancel()
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}

func watchByName(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "监控退出，停止了...")
			return
		default:
			fmt.Println(name, "goroutine监控中...")
			time.Sleep(2 * time.Second)
		}
	}
}

type key string

func testWithValue() {
	ctx, cancel := context.WithCancel(context.Background())
	//附加值
	valueCtx := context.WithValue(ctx, key("name"), "【监控1】")
	go watchByKey(valueCtx, key("name"))
	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	cancel()
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}

func watchByKey(ctx context.Context, k key) {
	for {
		select {
		case <-ctx.Done():
			//取出值
			fmt.Println(ctx.Value(k), "监控退出，停止了...")
			return
		default:
			//取出值
			fmt.Println(ctx.Value(k), "goroutine监控中...")
			time.Sleep(2 * time.Second)
		}
	}
}
