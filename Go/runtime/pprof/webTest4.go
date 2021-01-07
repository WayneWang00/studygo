package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	pprof2 "runtime/pprof"
	"sync"
	"time"
)

func main() {
	wait := sync.WaitGroup{}
	wait.Add(1)
	go func() {
		http.HandleFunc("/test", handler)
		err := http.ListenAndServe(":8181", nil)
		if err != nil {
			panic(err)
		}
	}()
	go func() {
		t(1)
	}()
	time.Sleep(time.Second * 2)
	t(1)
	wait.Wait()
}

func t(i int) {
	c := make(chan int, 1)
	fmt.Println(i)
	for {
		select {
		case <-c:
		default:

		}
	}

}
func handler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if nil != err {
		w.Write([]byte(err.Error()))
		return
	}
	doSomeThingOne(10000)
	buff := genSomeBytes()
	b, err := ioutil.ReadAll(buff)
	fmt.Println(pprof2.Lookup("heap").Count(), "===================================================================")
	if nil != err {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(b)
}

func doSomeThingOne(times int) {
	for i := 0; i < times; i++ {
		for j := 0; j < times; j++ {

		}
	}
}

func genSomeBytes() *bytes.Buffer {
	var buff bytes.Buffer
	for i := 1; i < 20000; i++ {
		buff.Write([]byte{'0' + byte(rand.Intn(10))})
	}
	return &buff
}
