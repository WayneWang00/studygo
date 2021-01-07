package rpc

import (
	rpcclient "byrpc/sapi/framework/rpcclient/core"
	plog "byrpc/sapi/putil/log"
	"runtime"
)

type MyDisPatch struct {
	client *rpcclient.RpcCall
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	defer plog.CatchPanic()
	quitFlag := make(chan bool)
	client, err := rpcclient.NewRpcCall()
	if err != nil {
		plog.Fatal("Fatal")
		return
	}
	err = client.RpcInit("./conf/clientapp.conf")
	if err != nil {
		plog.Debug("rpc初始化异常")
	}
	mydisp := new(MyDisPatch)
	mydisp.client = client
	err = client.LaunchRpcClient(mydisp)
	client.WriteNormalLog("test", "this is normal log")
	client.WriteRealLog("test", "this is real log")
	client.WriteDebugLog("test", "this is debug log")
	if err != nil {
		plog.Fatal("launch failed", err)
	}
	_ = <-quitFlag
}
