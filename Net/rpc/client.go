package main

import (
	"fmt"
	"net/rpc"
	"sync"
	"time"
)

const (
	CLIENT_PORT  = "127.0.0.1:1234"
	CLIENT_PROXY = "tcp"
)

type ClientHello struct {
	StrHello string
}

type ClientResStr struct {
	StrHi string
}

func main() {
	client1()
}

func client1() {
	client, err := rpc.DialHTTP(CLIENT_PROXY, CLIENT_PORT)
	if err != nil {
		fmt.Println("Dialing: ", err)
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		var args = ClientHello{"Call"}
		var reply string
		start := time.Now()
		err = client.Call("Say.SayHello", args, &reply)
		fmt.Println("duration:", time.Now().Sub(start))
		if err != nil {
			fmt.Println("airth error: ", err)
		}
		fmt.Println("Call: ", reply)
	}()
	//异步调用
	go func() {
		defer wg.Done()
		time.Sleep(3 * time.Second)
		args1 := ClientHello{"GO"}
		res := new(ClientResStr)
		start := time.Now()
		resCall := client.Go("Say.SayHi", args1, &res, nil)
		replyCall := <-resCall.Done
		fmt.Println("duration:", time.Now().Sub(start))
		if replyCall.Error != nil {
			fmt.Println("airth error: ", replyCall.Error)
		}
		fmt.Printf("resCall:%p, %+v\n", resCall.Done, replyCall)
		fmt.Println("Go: ", res.StrHi)
	}()

	wg.Wait()
	fmt.Println("client end")
}

func client2() {
	client, err := rpc.DialHTTP(CLIENT_PROXY, CLIENT_PORT)
	if err != nil {
		fmt.Println("Dialing: ", err)
	}
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(3 * time.Second)
		timeout <- true
	}()
	//var args = Hello{"Call"}
	//var reply string
	//err = client.Call("Say.SayHello", args, &reply)
	//if err != nil {
	//	fmt.Println("Call error: ", err)
	//}
	//fmt.Println("Call: ", reply)
	//异步调用
	args1 := ClientHello{"GO"}
	var res ClientResStr
	resCall := client.Go("Say.SayHi", args1, &res, nil)
	//replyCall := <-resCall.Done
	select {
	case replyCall := <-resCall.Done:
		if replyCall.Error == nil {
			fmt.Println("Go: ", res.StrHi)
		}
	case <-timeout:
		fmt.Println("timeout!")
	}
	//if replyCall.Error != nil {
	//	fmt.Println("Go error: ", replyCall.Error)
	//}
	fmt.Println("Go: ", res.StrHi)
}
