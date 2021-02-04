package main

import (
	"runtime"
	"sync"

	"fmt"
	"time"
)

func paramstest(s map[string]string, i ...interface{}) interface{} {
	return nil
}

func main() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(32)
	//for i := 0; i < 100000000; i++{
	//	if i<10 {
	//		go func() {
	//			fmt.Println("fffff",time.Now())
	//			fmt.Println("A: ", i)
	//			wg.Done()
	//
	//		}()
	//	}
	//	//time.Sleep(10*time.Millisecond)
	//}
	//fmt.Println("outside",time.Now())
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("A:", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("inside:", time.Now())
			fmt.Println("B: ", i)
			wg.Done()
		}(i)
	}
	fmt.Println("outside1:", time.Now())
	for i := 0; i < 10; i++ {
		//defer wg.Done()
		go func() {
			fmt.Println("C: ", i)
			wg.Done()
		}()
	}
	go func() {
		defer wg.Done()
		for i := 0; i < 100000; i++ {
			fmt.Println("D: ", i)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 1000000; i++ {
			fmt.Println("E: ", i)
		}
	}()
	fmt.Println("outside2: ", time.Now())
	//for i :=0; i < 10; i++{
	//	go func(i int) {
	//		fmt.Println("C: ", i)
	//		wg.Done()
	//	}(i)
	//}
	wg.Wait()
	//fmt.Println("测试")
}
