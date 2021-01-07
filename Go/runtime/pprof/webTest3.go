package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"sync"
	"time"
)

var (
	mu    sync.Mutex
	mu2   sync.Mutex
	items = make(map[int]struct{})
)

func main() {
	runtime.SetMutexProfileFraction(5)
	runtime.SetBlockProfileRate(1)
	for i := 0; i < 100*100; i++ {
		go func(i int) {
			mu.Lock()
			mut()
			defer mu.Unlock()
			items[i] = struct{}{}
		}(i)
		go func(i int) {
			mu2.Lock()
			mut()
			defer mu2.Unlock()
			//items[i] = struct{}{}
		}(i)
	}
	fmt.Println("====")
	http.ListenAndServe(":8182", nil)
	// http://localhost:8888/debug/pprof/
}
func mut() {
	time.Sleep(time.Second)
	mute()
}
func mute() {
	time.Sleep(time.Second)
}
