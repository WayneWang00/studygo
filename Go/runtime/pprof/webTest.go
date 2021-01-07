package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"sync"
	"time"
)

var Count int64 = 0
var mu = sync.Mutex{}

type info struct {
	Name  string
	Adrss string
	Phone string
}

func main() {
	go calCount()
	runtime.SetBlockProfileRate(1)
	runtime.SetMutexProfileFraction(1)
	http.HandleFunc("/test", test)
	http.HandleFunc("/data", handlerData)
	go func() {
		mu.Lock()
		time.Sleep(time.Second * 10)
		defer mu.Unlock()

	}()
	go func() {
		mu.Lock()
		time.Sleep(time.Second * 10)
		defer mu.Unlock()

	}()
	go func() {
		mu.Lock()
		time.Sleep(time.Second * 10)
		defer mu.Unlock()

	}()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func handlerData(w http.ResponseWriter, r *http.Request) {
	fmt.Println(runtime.Caller(0))
	info := new(info)
	info.Name = "===================="
	info = nil

	qUrl := r.URL
	fmt.Println(qUrl)
	fibRev := Fib()
	var fib uint64
	for i := 0; i < 5000; i++ {
		fib = fibRev()
		fmt.Println("fib = ", fib)
	}
	//str := RandomStr(RandomInt(100, 500))
	str := fmt.Sprintf("Fib = %d", fib)
	w.Write([]byte(str))
}

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Println(runtime.Caller(0))
	info := new(info)
	info.Name = "===================="
	info = nil

	fibRev := Fib()
	var fib uint64
	index := Count
	arr := make([]uint64, index)
	var i int64
	for ; i < index; i++ {
		fib = fibRev()
		arr[i] = fib
		fmt.Println("fib = ", fib)
	}
	time.Sleep(time.Millisecond * 500)
	str := fmt.Sprintf("Fib = %v", arr)
	w.Write([]byte(str))
}

func Fib() func() uint64 {
	var x, y uint64 = 0, 1
	return func() uint64 {
		x, y = y, x+y
		return x
	}
}

func calCount() {
	timeInterval := time.Tick(time.Second)

	for {
		select {
		case i := <-timeInterval:
			Count = int64(i.Second())
		}
	}
}
