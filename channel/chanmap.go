package main

import (
	"fmt"
	"sync"
	"time"
)

type Bank struct {
	save map[string]int64
}

type Request struct {
	op    string
	name  string
	value int64
	ret   chan *Result
}

type Result struct {
	success bool
	value   int64
}

func main() {
	reqChan := make(chan *Request, 100)
	bank := NewBank()

	go bank.Loop(reqChan)

	var wg sync.WaitGroup
	wg.Add(2)
	go User1(&wg, reqChan)
	go User2(&wg, reqChan)

	wg.Wait()
	close(reqChan)

	time.Sleep(time.Second)
	fmt.Println("The End")
}

// 银行初始化
func NewBank() *Bank {
	b := &Bank{
		save: make(map[string]int64),
	}

	return b
}

func (b *Bank) Loop(reqChan chan *Request) {
	for req := range reqChan {
		switch req.op {
		case "save":
			b.Save(req)
		case "withDraw":
			b.WithDraw(req)
		case "Query":
			b.Query(req)
		default:
			ret := &Result{
				success: false,
				value:   0,
			}
			req.ret <- ret
		}
	}

	fmt.Println("Bank Exit")
}

func (b *Bank) Save(req *Request) {}

func (b *Bank) WithDraw(req *Request) {}

func (b *Bank) Query(req *Request) {}

func User1(wg *sync.WaitGroup, ch chan<- *Request) {
	defer wg.Done()
}

func User2(wg *sync.WaitGroup, ch chan<- *Request) {
	defer wg.Done()
}
