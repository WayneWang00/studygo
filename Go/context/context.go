package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//testWithValue()
	//testManyWithCancel()
	testWithDeadline()
}

var key = "name"

func watch(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Value(key), "监控停止了，退出...")
			return
		default:
			fmt.Println(ctx.Value(key), "goroutine监控中")
			time.Sleep(2 * time.Second)
		}
	}
}

// 附加元素
func testWithValue() {
	ctx, cancel := context.WithCancel(context.Background())
	// 附加值
	valueCtx := context.WithValue(ctx, key, "[监控1]")
	go watch(valueCtx)
	time.Sleep(10 * time.Second)
	fmt.Println("结束，通知监控停止")
	cancel()
	// 检测是否监控停止
	time.Sleep(5 * time.Second)
}

func withCancel(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "监控停止了，退出...")
			return
		default:
			fmt.Println(name, "goroutine监控中")
			time.Sleep(2 * time.Second)
		}
	}
}

// 控制多个goroutine
func testManyWithCancel() {
	ctx, cancel := context.WithCancel(context.Background())
	go withCancel(ctx, "[监控1]")
	go withCancel(ctx, "[监控2]")
	go withCancel(ctx, "[监控3]")

	time.Sleep(10 * time.Second)
	fmt.Println("结束，通知监控停止")
	cancel()

	time.Sleep(5 * time.Second)
}

// 设置截止时间
func testWithDeadline() {
	d := time.Now().Add(50 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
