package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex
	sharedResource := 0
	isRunning := false

	for i := 1; i <= 3; i++ {
		go func(id int) {
			fmt.Printf("Goroutine %d 尝试获取锁\n", id)
			mu.Lock()
			fmt.Printf("Goroutine %d 获取锁成功\n", id)
			if isRunning {
				fmt.Printf("Goroutine %d 其他服务正在运行，退出\n", id)
				mu.Unlock()
				return
			}
			isRunning = true

			// 模拟临界区操作
			fmt.Printf("Goroutine %d 开始操作共享资源 (value: %d)\n", id, sharedResource)
			sharedResource++
			time.Sleep(1 * time.Second)
			fmt.Printf("Goroutine %d 完成操作 (new value: %d)\n", id, sharedResource)

			isRunning = false
			mu.Unlock()
			fmt.Printf("Goroutine %d 释放锁\n", id)
		}(i)
	}

	// 等待goroutines执行完毕
	time.Sleep(3 * time.Second)
	fmt.Println("最终共享资源值:", sharedResource)
}
