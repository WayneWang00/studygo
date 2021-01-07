package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	//"time"
)

//登陆信息
type User struct {
	userName string
	passWord string
}

//登陆信息集合
var user []User

//接收数据通道
var recvDataServer = make(chan string)

//发送数据通道
var sendDataServer = make(chan string)

//大厅信息
var hallinfo = `[{"room_id":1,"room_name":"chat1","online_num":2,"total_num":3},{"room_id":2,"room_name":"chat2","online_num":2,"total_num":3}]`

//玩家信息
var userinfo = `[{"user_name":"Bob","head_photo":"image1","status":"on","grade":"一级"},{"user_name":"Tom","head_photo":"image2","status":"off","grade":"二级"}]`

type HallInfoServer struct {
	Room_id    int    `json:"room_id"`
	Room_name  string `json:"room_name"`
	Total_num  int    `json:"total_num"`
	Online_num int    `json:"online_num"`
}

//聊天室列表
var hallInfoServer []HallInfoServer

//用户信息
type UserInfoServer struct {
	User_name  string `json:"user_name"`
	Head_photo string `json:"head_photo"`
	Status     string `json:"status"`
	Grade      string `json:"grade"`
}

//玩家列表
var userInfoServer []UserInfoServer

func main() {
	//开启服务
	startServer()
}

//开始服务器
func startServer() {
	listen, err := net.Listen("tcp", "127.0.0.1:8080")
	checkErrServer(err)
	fmt.Println("建立成功")
	for {
		//等待客户端连接
		conn, err := listen.Accept()
		a := checkErrServer(err)
		fmt.Println("a: ", a)
		fmt.Println("客户端: ", conn.RemoteAddr().String(), "连接服务器成功")
		//开一个goroutine处理客户端消息
		go handleClient(conn)
	}
}

//处理客户端消息
func handleClient(conn net.Conn) {
	defer conn.Close()
	go recvServer(recvDataServer, conn) //启动数据接收协程
	go sendServer(sendDataServer, conn) //启动数据发送协程
	go authUser(recvDataServer)

	//nameInfo := make([]byte, 512)
	//_, err := conn.Read(nameInfo)
	//checkErrServer(err)
	//fmt.Println(string(nameInfo))
	//dayTime := time.Now().Format("2006-01-02 15:04:05")
	//conn.Write([]byte(dayTime))
	//for {
	//	buf := make([]byte, 512)
	//	//读取客户端发送的请求
	//	_, err := conn.Read(buf)
	//	flag := checkErrServer(err)
	//	if flag == 0 {
	//		break
	//	}
	//	fmt.Println(string(buf))
	//}
}

//接收数据
func recvServer(recvData chan string, conn net.Conn) {
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	checkErrServer(err)
	recvData <- string(buf[:n])
}

//发送数据
func sendServer(sendData chan string, conn net.Conn) {
	buf := <-sendData
	_, err := conn.Write([]byte(buf))
	checkErrServer(err)
}

//验证用户
func authUser(recvData chan string) {
	for {
		data := strings.Split(<-recvData, "-")
		if len(data) == 1 {
			fmt.Println(data[0])
			continue
		}
		flag, username, password := data[0], data[1], data[2]
		a := User{
			userName: username,
			passWord: password,
		}
		if flag == "Reg" {
			user = append(user, a)
			sendDataServer <- "SUCCESS"
		} else if flag == "Log" {
			for i, _ := range user {
				if user[i].userName == username && user[i].passWord == password {
					sendDataServer <- "SUCCESS"
					sendDataServer <- hallinfo
				}
			}
		} else if flag == "List" {
			json.Unmarshal([]byte(hallinfo), &hallInfoServer)
			for i, _ := range hallInfoServer {
				passwordint, _ := strconv.Atoi(password)
				if hallInfoServer[i].Room_id == passwordint {
					sendDataServer <- "SUCCESS"
					sendDataServer <- userinfo
				}
			}
		} else if flag == "heardPacket" {
			fmt.Println(password)
		}
	}
}

//检查错误
func checkErrServer(err error) int {
	if err != nil {
		if err.Error() == "EOF" {
			//fmt.Println("用户退出了")
			return 0
		}
		log.Fatal("an error: ", err.Error())
		return -1
	}
	return 1
}
