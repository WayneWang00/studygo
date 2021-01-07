package main

import (
	"fmt"
	"os"
	"runtime/pprof"
	"sync"
	"time"
)

func main() {
	f, e := os.OpenFile("src/pprof/goroutine/goroutine.pprof", os.O_RDWR|os.O_CREATE, 644)
	if e != nil {
		fmt.Print("创建文件失败")
	}
	defer f.Close()
	var wg = sync.WaitGroup{}
	wg.Add(20)

	for i := 0; i < 20; i++ {
		go func() {

			for z := 0; z < 100000; z++ {
				z = z
			}
			time.Sleep(time.Second * 3)
			wg.Done()

		}()
	}
	pprof.Lookup("goroutine").WriteTo(f, 2)
	wg.Wait()
	defer f.Close()
}
