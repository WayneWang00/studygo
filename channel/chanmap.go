package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type Bank struct {
	save map[string]int64 // 所有用户金额
}

type Request struct {
	op    string       // 操作
	name  string       // 用户
	value int64        // 金额
	ret   chan *Result // 响应
}

type Result struct {
	success bool  // 是否操作成功
	value   int64 // 剩余金额
}

func main() {
	reqChan := make(chan *Request, 100)
	bank := NewBank()

	go bank.Loop(reqChan)

	var wg sync.WaitGroup
	var n = 5
	wg.Add(n)
	//go User1(&wg, reqChan)
	//go User2(&wg, reqChan)
	for i := 0; i < n; i++ {
		go User(&wg, reqChan, i)
	}

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
		case "query":
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

// 存钱
func (b *Bank) Save(req *Request) {
	user := req.name
	value := req.value

	if _, ok := b.save[user]; !ok {
		b.save[user] = 0
	}

	b.save[user] += value
	ret := &Result{
		success: true,
		value:   b.save[user],
	}
	req.ret <- ret
}

// 取钱
func (b *Bank) WithDraw(req *Request) {
	user := req.name
	value := req.value

	var status bool
	var money int64
	if v, ok := b.save[user]; !ok || v < value {
		status = false
		money = v
	} else {
		b.save[user] -= value
		status = true
		money = b.save[user]
	}

	ret := &Result{
		status,
		money,
	}
	req.ret <- ret
}

// 查询
func (b *Bank) Query(req *Request) {
	user := req.name

	money := b.save[user]
	ret := &Result{
		success: true,
		value:   money,
	}
	req.ret <- ret
}

func User(wg *sync.WaitGroup, ch chan<- *Request, n int) {
	fmt.Println("user", n)
	retChan := make(chan *Result)
	name := "user" + strconv.Itoa(n)
	value := 100 * int64(n+1)
	draw := 20 * int64(n+1)

	defer func() {
		close(retChan)
		wg.Done()
	}()

	saveReq := &Request{
		"save",
		name,
		value,
		retChan,
	}
	withDrawReq := &Request{
		"withDraw",
		name,
		draw,
		retChan,
	}
	queryReq := &Request{
		op:   "query",
		name: name,
		ret:  retChan,
	}

	reqs := []*Request{saveReq, withDrawReq, queryReq}
	for _, v := range reqs {
		ch <- v
		WaitRet(v)
	}
}

func User1(wg *sync.WaitGroup, ch chan<- *Request) {
	name := "user1"
	retChan := make(chan *Result)

	defer func() {
		close(retChan)
		wg.Done()
	}()

	saveReq := &Request{
		op:    "save",
		name:  name,
		value: 100,
		ret:   retChan,
	}
	withDrawReq := &Request{
		"withDraw",
		name,
		30,
		retChan,
	}
	queryReq := &Request{
		op:   "query",
		name: name,
		ret:  retChan,
	}
	fmt.Printf("save:%p, withDraw:%p, query:%p\n", saveReq, withDrawReq, queryReq)

	reqs := []*Request{saveReq, withDrawReq, queryReq}
	for k, v := range reqs {
		fmt.Printf("%d:&v:%p\n", k, v)
		ch <- v
		WaitRet(v)
	}
}

func User2(wg *sync.WaitGroup, ch chan<- *Request) {
	name := "user2"
	retChan := make(chan *Result)

	defer func() {
		close(retChan)
		wg.Done()
	}()

	saveReq := &Request{
		"save",
		name,
		100,
		retChan,
	}
	withDrawReq := &Request{
		op:    "withDraw",
		name:  name,
		value: 120,
		ret:   retChan,
	}
	queryReq := &Request{
		"query",
		name,
		0,
		retChan,
	}

	reqs := []*Request{saveReq, withDrawReq, queryReq}
	for _, v := range reqs {
		ch <- v
		WaitRet(v)
	}
}

func WaitRet(req *Request) {
	ret := <-req.ret

	if ret.success {
		fmt.Printf("Success: name:%s op:%s value:%d ret value:%d\n", req.name, req.op, req.value, ret.value)
		return
	}

	fmt.Printf("Failed: name:%s op:%s value:%d ret value:%d\n", req.name, req.op, req.value, ret.value)
}
