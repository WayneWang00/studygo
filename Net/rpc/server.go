package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
)

const (
	SAY          = "Test"
	SERVER_PORT  = ":1234"
	SERVER_PROXY = "tcp"
)

type ServerHello struct {
	StrHello string
}

type ServerResStr struct {
	StrHi string
}

type Say int

func (s *Say) SayHello(args *ServerHello, reply *string) error {
	//fmt.Println("等待5秒...")
	//time.Sleep(5 * time.Second)
	//fmt.Println("结束等待")
	*reply = SAY + args.StrHello
	return nil
}

func (s *Say) SayHi(args *ServerHello, reply *ServerResStr) error {
	//fmt.Println("等待5秒...")
	//time.Sleep(5 * time.Second)
	//fmt.Println("结束等待")
	reply.StrHi = SAY + args.StrHello
	return nil
}

func main() {
	hi := new(Say)
	rpc.Register(hi)
	fmt.Println("服务启动")
	rpc.HandleHTTP()
	l, e := net.Listen(SERVER_PROXY, SERVER_PORT)
	if e != nil {
		fmt.Println("Listen error:", e)
	}
	http.Serve(l, nil)
}
