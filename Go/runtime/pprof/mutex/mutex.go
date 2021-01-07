package main

import (
	"fmt"
	"os"
	"runtime/pprof"
	"sync"
	"time"
)

func main() {
	f, e := os.OpenFile("src/pprof/mutex/mutex.pprof", os.O_RDWR|os.O_CREATE, 644)
	if e != nil {
		fmt.Print("创建文件失败")
	}
	defer f.Close()
	var m sync.Mutex
	var wait sync.WaitGroup
	wait.Add(3)
	for i := 0; i < 3; i++ {
		go func() {
			m.Lock()
			time.Sleep(time.Second * 10)
			m.Unlock()
			wait.Done()
		}()
	}
	time.Sleep(time.Second * 2)
	err := pprof.Lookup("mutex").WriteTo(f, 2)
	if err != nil {
		fmt.Println("mutex 文件写入错误")
	}
	wait.Wait()
}
