package main

import (
	"runtime"
	"sync"
	"time"

	"fmt"
)

func paramsTest(s map[string]string, i ...interface{}) interface{} {
	return nil
}

func main() {
	//paramsTest(map[string]string{})
	//go32()
	Add26()
}

func go32() {
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

func Add26() {
	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	wg.Add(2)
	fmt.Println("start goroutine")

	go func() {
		defer wg.Done()

		for count := 0; count < 3; count++ {
			fmt.Println("a count:", count)
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
			fmt.Println()
		}
	}()

	go func() {
		defer wg.Done()

		for count := 0; count < 3; count++ {
			fmt.Println("A count:", count)
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
			fmt.Println()
		}
	}()

	fmt.Println("waiting ...")
	wg.Wait()
	fmt.Println("end goroutine")
}
