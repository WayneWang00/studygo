package rpc

import (
	"PGDemo/app/arithproto"
	"framework/rpcclient/szmq" //zmq日志操作
	//"putil/log"
)

type Arith int

//var (
//	Logger *szmq.SzmqPushClient //通过zmq写日志！
//)

//方法需要手动注册

//非反射
//func (arith *Arith) Add(rq *arithproto.ArithRequest) *arithproto.ArithResponse {
//	rp := new(arithproto.ArithResponse)
//	rp.A3 = int32(rq.A1 + rq.A2)
//	//plog.Debug("the result is: ", rp.A3)
//	go szmq.Logger.WriteNormalLog("test", "hello,the log comes from arith!")

//	return rp
//}

//反射
func (arith *Arith) Add(rq *arithproto.ArithRequest) (rp *arithproto.ArithResponse) {
	//rp := new(arithproto.ArithResponse)
	rp.A3 = int32(rq.A1 + rq.A2)
	//plog.Debug("the result is: ", rp.A3)
	go szmq.Logger.WriteNormalLog("test", "hello,the log comes from arith!")

	return rp
}

//乘法服务
func (arith *Arith) Multiply(rq *arithproto.ArithRequest) *arithproto.ArithResponse {
	rp := new(arithproto.ArithResponse)
	rp.A3 = int32(rq.A1 * rq.A2)
	return rp
}
