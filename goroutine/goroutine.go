package main

import (
	"runtime"
	"fmt"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	//var wg sync.WaitGroup
	//wg.Add(2)
	go func(){
		//defer wg.Done()
		for i := 0; i < 10; i++{
			fmt.Println("A: ", i)
		}
	}()
	go func(){
		//defer wg.Done()
		for i := 0; i < 10; i++{
			fmt.Println("B: ", i)
		}
	}()
	//wg.Wait()
	time.Sleep(100*time.Millisecond)
}
