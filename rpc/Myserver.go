package rpc

import (
	"PGDemo/app/arithproto"
	"Wayne/arithmetic" //业务代码
	rpcclient "byrpc/sapi/framework/rpcclient/core"
	plog "byrpc/sapi/putil/log"
	"runtime"
	"strings"
)

var arithobj = new(arithmetic.Arith) //Server实例，为处理接收到的请求做准备
//Arith是一个自定义数据类型，在arith.go中.包含两个方法:
//Add:加法
//Multiply:乘法

type MyDisPatch struct {
	client *rpcclient.RpcCall
	/*RpcCall是个结构体，在rpccall.go中，包含的方法：
		RpcInit : 初始化rpc服务
		SetSvr ：设置地址 sertype，serid
		SetSvrNames ：设置地址 servername，servicename，functionname
		SetLocalAddress ：绑定本地ip，port
		SendAndRecvRespRpcMsg ：接口（rpc点对点有响应）
		SendAndRecvWithTypeId ：接口（rpc点对点有响应，指定svrtype和svrid的情况）
		SendNoRespRpcMsg : rpc点对点无响应(下方原注释仅供参考)
	    给目标[dststype,dstsvrid]调用方法(methodname),调用发送的数据是req，无需返回
	    这里groupids广播的时候用，groupids中的字段可以填充多个[servertype+svrsvrid]，给这组服务发送消息，注意，这里只能提供推送，并且目前这个逻辑在Agent是否实现有待确认[暂不推荐使用该groupid]。
	    目前groupids的设计是基于业务来设计的
		SendNoRespRpcMsgGroup : rpc点对组无响应
		SendPacket : 给源[ssvrtype, ssvrid]返回rpc请求的响应  注意，这里req来自于请求的rpc请求的时交付给业务的req，包含了请求的[type,id],以及请求的头和对应的mtid，不要修改，否则对方有收不到返回的情况
		SendNotifyMsg : 推送接口，推送接口的数据不是rpc数据，是自定义数据，这里是非rpc消息发送， 给客户端发送主要用此接口
		SendMsgToClient : 通过网关服务[dstsvrtype, dstsvrid]给客户端id发送数据msgdata，注意该接口暂不建议使用。
		LaunchRpcClient : 启动rpc的调用客户端，连接到Agent,Agent的真实网络地址是【remoteip,remoteport】对收到的rpc请求通过rev来接口
		WriteNormalLog : 通过zmq发送日志(普通日志)
		WriteRealLog : 通过zmq发送日志（实时日志）
		WriteDebugLog : 通过zmq发送日志（调试日志）
	*/
}

//接收rpc请求并计算返回
func (dis *MyDisPatch) RpcRequest(req *rpcclient.RpcRecvReq, body []byte) { //RpcRecvReq是个结构体,在rpccall.go中.
	plog.Debug("here comes RpcRequest! ")
	serviceAndMethodName := string(req.Rpchead.MethodName) //方法名称(eg：User.getUserInfo)
	var methodName string
	//转成纯方法名称
	pos := strings.LastIndex(serviceAndMethodName, ".") //返回字符串中"."最后出现的位置
	if pos > 0 && pos < len(serviceAndMethodName) {
		methodName = serviceAndMethodName[pos+1:] //避免越界
	}
	rt := []byte{} //最终返回的字节切片
	plog.Debug("here comes RpcRequest!", methodName)
	//TODO检查方法名称是否存在
	switch methodName {
	case "Add ":
		//接收参数并处理
		arithReq := new(arithproto.ArithRequest)
		arithReq.Unmarshal(body)
		//arithResp := new( arithproto.ArithResponse)
		arithResp := arithobj.Add(arithReq)
		plog.Debug("return", arithResp.A3)
		rt, _ = arithResp.Marshal()
	case "Multiply":
		//接收参数并处理
		arithReq := new(arithproto.ArithRequest)
		arithReq.Unmarshal(body)                      //解参数的pb包
		arithResp := arithobj.Multiply(arithReq)      //计算
		plog.Debug("Multiply returns:", arithResp.A3) //日志
		rt, _ = arithResp.Marshal()                   //返回数据的pb格式化
	default:
		//rt := []byte{}
	}
	if req.NeedResp == true { //NeedResp 是结构体ArithRequest中一个bool类型的元素
		dis.client.SendPacket(req, rt) //处理完毕之后返回数据！
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	defer plog.CatchPanic()     //Panic异常处理
	quitFlag := make(chan bool) //退出处理
	//实例化rpc
	client, err := rpcclient.NewRpcCall() //创建RPC调用接口。NewRpcCall函数：第一个参数，RpcCaller结构体
	if err != nil {
		plog.Fatal("Fatal")
		return
	}
	err = client.RpcInit("") //初始化RPC服务
	if err != nil {
		plog.Debug("rpc初始化异常")
		return
	}
	mydisp := new(MyDisPatch)
	mydisp.client = client
	//设置本rpc服务的代理人（Net层）的ip和端口，并启动服务！
	err = client.LaunchRpcClient(mydisp)
	if err != nil {
		plog.Fatal("launch failed", err)
		return
	}
	//控制退出
	_ = <-quitFlag //收消息退出！
}
