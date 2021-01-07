package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"time"
)

type User1 struct {
	userName string
	passWord string
}

//登陆信息集合
var users1 []User1

////接收数据通道
//var recvDataServer1 = make(chan string)
//
////发送数据通道
//var sendDataServer1 = make(chan string)

//大厅信息
var hallinfo1 = `[{"room_id":1,"room_name":"chat1","online_num":2,"total_num":3},{"room_id":2,"room_name":"chat2","online_num":2,"total_num":3}]`

//玩家信息
var userinfo1 = `[{"user_name":"Bob","head_photo":"image1","status":"on","grade":"一级"},{"user_name":"Tom","head_photo":"image2","status":"off","grade":"二级"}]`

type HallInfoServer1 struct {
	Room_id    int    `json:"room_id"`
	Room_name  string `json:"room_name"`
	Total_num  int    `json:"total_num"`
	Online_num int    `json:"online_num"`
}

//聊天室列表
var hallInfoServer1 []HallInfoServer1

func main() {
	//开启服务
	startServer1()
}

//开始服务器
func startServer1() {
	listen, err := net.Listen("tcp", "127.0.0.1:8080")
	checkErrServer1(err)
	fmt.Println("建立成功")
	for {
		//等待客户端连接
		conn, err := listen.Accept()
		a := checkErrServer1(err)
		fmt.Println("a: ", a)
		fmt.Println("客户端: ", conn.RemoteAddr().String(), "连接服务器成功")
		//开一个goroutine处理客户端消息
		go handleClient1(conn)
	}
}

//处理客户端消息
func handleClient1(conn net.Conn) {
	defer conn.Close()

	//nameInfo := make([]byte, 512)
	//_, err := conn.Read(nameInfo)
	//checkErrServer1(err)
	//fmt.Println(string(nameInfo))
	dayTime := time.Now().Format("2006-01-02 15:04:05")
	conn.Write([]byte(dayTime))
	for {
		buf := make([]byte, 512)
		//读取客户端发送的请求
		_, err := conn.Read(buf)
		flag := checkErrServer1(err)
		if flag == 0 {
			break
		}
		fmt.Println(string(buf))
		for {
			data := strings.Split(string(buf), "-")
			if len(data) == 1 {
				fmt.Println(data[0])
				continue
			}
			flag, username, password := data[0], data[1], data[2]
			a := User1{
				userName: username,
				passWord: password,
			}
			if flag == "Reg" {
				users1 = append(users1, a)
				conn.Write([]byte("SUCCESS"))
			} else if flag == "Log" {
				for i, _ := range users1 {
					if users1[i].userName == username && users1[i].passWord == password {
						conn.Write([]byte("SUCCESS"))
						conn.Write([]byte(hallinfo1))
					}
				}
			} else if flag == "List" {
				json.Unmarshal([]byte(hallinfo1), &hallInfoServer1)
				for i, _ := range hallInfoServer1 {
					passwordint, _ := strconv.Atoi(password)
					if hallInfoServer1[i].Room_id == passwordint {
						conn.Write([]byte("SUCCESS"))
						conn.Write([]byte(userinfo1))
					}
				}
			} else if flag == "heardPacket" {
				fmt.Println(password)
			}
		}
	}
}

//检查错误
func checkErrServer1(err error) int {
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
