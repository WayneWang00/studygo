package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"sync"
	"time"
)

/*var (
	mutex1 = sync.Mutex{}
	mutex2 = sync.Mutex{}
)
*/
var (
	mu sync.Mutex
)

func main() {
	runtime.SetBlockProfileRate(1)
	runtime.SetMutexProfileFraction(1)
	f, e := os.OpenFile("src/pprof/block/block.pprof", os.O_RDWR|os.O_CREATE, 644)
	ff, ee := os.OpenFile("src/pprof/block/mutex.pprof", os.O_RDWR|os.O_CREATE, 644)
	if e != nil {
		fmt.Print("创建文件失败")
	}
	if ee != nil {
		fmt.Print("创建文件失败", ee)
	}
	defer f.Close()
	defer ff.Close()

	go func() {
		mu.Lock()
		time.Sleep(time.Second * 5)
		defer mu.Unlock()

	}()
	go func() {
		mu.Lock()
		mu.Lock()

		defer mu.Unlock()

	}()

	time.Sleep(time.Second * 5)
	pprof.Lookup("mutex").WriteTo(ff, 2)
	pprof.Lookup("block").WriteTo(f, 2)

}
