package main

import (
	"fmt"
	"sync"
)

func main() {
	// print1()
	print2()
}

func print1() {
	oddChan := make(chan struct{})  // 奇数通道
	evenChan := make(chan struct{}) // 偶数通道
	done := make(chan struct{})

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 1; i <= 100; i += 2 {
			<-oddChan
			fmt.Println(i)
			if i+1 <= 100 {
				evenChan <- struct{}{}
			} else {
				done <- struct{}{}
			}
		}
	}()

	go func() {
		defer wg.Done()
		for i := 2; i <= 100; i += 2 {
			<-evenChan
			fmt.Println(i)
			if i+1 <= 100 {
				oddChan <- struct{}{}
			}
		}
	}()

	oddChan <- struct{}{}
	<-done // 会出现死锁
	wg.Wait()
}

func print2() {
	oddChan := make(chan struct{}, 1)
	eventChan := make(chan struct{}, 1)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 1; i < 100; i += 2 {
			<-oddChan
			fmt.Println(i)
			eventChan <- struct{}{}
		}
	}()

	go func() {
		defer wg.Done()
		for i := 2; i <= 100; i += 2 {
			<-eventChan
			fmt.Println(i)
			oddChan <- struct{}{}
		}
	}()

	oddChan <- struct{}{}
	wg.Wait()
}
