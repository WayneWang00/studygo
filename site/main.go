package main

import (
	"site/static"
	"fmt"
	"net/http"
)

func main() {
	//// 获取当前目录，绝对路径
	//pwd, pwdErr := os.Getwd()//在Linux下为“/”，在Windows下为“\”。
	//if pwdErr != nil {
	//	fmt.Println(pwdErr)
	//	os.Exit(1)
	//}
	//// 设置路由
	//http.HandleFunc("/getcdn", index)
	//log.Println(pwd)
	//log.Println(pwd + "/static")
	//// 设置静态文件服务目录，需要绝对路径
	//fsh := http.FileServer(http.Dir("E:/GOPATH/Global/src/Wayne/site/static"))
	////fsh :=http.FileServer(http.Dir(pwd + "/static"))
	//fmt.Println(fsh)
	//// 设置静态资源路由
	//http.Handle("/static/", http.StripPrefix("/static/", fsh))
	//// http服务启动监听端口
	//err := http.ListenAndServe(":9001", nil)
	//if err != nil {
	//	log.Fatal("ListenAndServe: ", err)
	//}

	//测试相对路径
	static.GetFile()
}

func index(w http.ResponseWriter, r *http.Request) {
	json := `{"hall":[{"ip":"192.168.201.77","port":7000},{"ip":"192.168.201.77","port":7001},{"ip":"192.168.201.77","port":7002},{"ip":"192.168.201.77","port":7003},{"ip":"192.168.201.77","port":7004},{"ip":"dfaccess.oa.com","port":7000},{"ip":"dfaccess.oa.com","port":7001},{"ip":"dfaccess.oa.com","port":7002},{"ip":"dfaccess.oa.com","port":7003},{"ip":"dfaccess.oa.com","port":7004}],"hall_backup":{},"php":[{"url":"http:\/\/dfqptest01.oa.com\/dfqp\/"}],"cdn":[{"url":"http://produce.oa.com:9001/static/cdn.json"}]}`
	fmt.Fprintf(w, json)
}
