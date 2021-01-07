package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

//用户名
var uName string

//接收数据通道
var recvDataClient = make(chan string)

//发送数据通道
var sendDataClient = make(chan string)

//关闭通道
//var stopChan = make(chan bool)

//关闭信息显示通道
var stopMessage = make(chan bool)

//大厅信息
type HallInfo struct {
	Room_id    int    `json:"room_id"`
	Room_name  string `json:"room_name"`
	Total_num  int    `json:"total_num"`
	Online_num int    `json:"online_num"`
}

//聊天室列表
var hallInfo []HallInfo

//用户信息
type UserInfo struct {
	User_name  string `json:"user_name"`
	Head_photo string `json:"head_photo"`
	Status     string `json:"status"`
	Grade      string `json:"grade"`
}

//玩家列表
var userInfo []UserInfo

func main() {
	//连接server
	connectServer()
}

//连接服务器
func connectServer() {
	//连通
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	checkErrClient(err)
	fmt.Println("连接成功")
	go recvClient(recvDataClient, conn) //启动数据接收协程
	go sendClient(sendDataClient, conn) //启动数据发送协程
	//启动发送心跳协程
	go func() {
		heartInterval := time.Tick(10 * time.Second)
		for {
			select {
			case <-heartInterval:
				sendHeardPacket()
				//case <-stopChan:
				//	return
			}
		}
	}()
	fmt.Println("1.注册  2.登录")
	//输入
	//inputReader := bufio.NewReader(os.Stdin)
	clientLog() //进入注册登录

}

//发送心跳包
func sendHeardPacket() {
	timeStamp := time.Now().Format("2006-01-02 15:04:05")
	send := "heardPacket-" + uName + timeStamp
	sendDataClient <- send
}

//接收数据
func recvClient(recvData chan string, conn net.Conn) {
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	checkErrClient(err)
	recvData <- string(buf[:n])
}

//发送数据
func sendClient(sendData chan string, conn net.Conn) {
	buf := <-sendData
	_, err := conn.Write([]byte(buf))
	checkErrClient(err)
}

//注册登录
func clientLog() {
LOOP:
	fmt.Println("请输入“1”或“2”来选择注册或登录：")
	var option string
	fmt.Scanln(&option)
	switch option {
	case "1":
		register() //进入注册
	case "2":
		login() //进入登录
	default:
		fmt.Println("输入的序号错误！")
		goto LOOP
	}
}

//注册程序
func register() {
	for {
		var userName, pwd, pwdCheck string
		fmt.Println("请输入用户名：")
		fmt.Scanln(&userName)
		fmt.Println("请输入密码：")
		fmt.Scanln(&pwd)
		fmt.Println("请确认密码：")
		fmt.Scanln(&pwdCheck)
		if pwdCheck != pwd {
			fmt.Println("两次密码不一样，请重新注册！")
		} else {
			send := "Reg-" + userName + "-" + pwd
			sendDataClient <- send
			res := <-recvDataClient
			if res == "SUCCESS" {
				fmt.Println("注册成功，请登录")
				break
			} else {
				fmt.Println(res)
			}
		}
	}
	login() //登录
}

//登录程序
func login() {
LOOP:
	for {
		var userName, password string
		fmt.Println("请输入用户名：")
		fmt.Scanln(&userName)
		if userName == "" {
			goto LOOP
		}
		fmt.Println("请输入密码：")
		fmt.Scanln(&password)
		send := "Log-" + userName + "-" + password
		sendDataClient <- send
		res := <-recvDataClient
		if res == "SUCCESS" {
			uName = userName
			fmt.Println("登录成功")
			res := <-recvDataClient
			//chatRoomList = res
			json.Unmarshal([]byte(res), &hallInfo)
			chatRoom(hallInfo) //进入聊天室
			break
		} else {
			fmt.Println(res)
		}
	}
}

//选择聊天室
func chatRoom(chatroomlist []HallInfo) {
	var ordinal string
	for i, _ := range chatroomlist {
		fmt.Println("room_id: ", chatroomlist[i].Room_id, "room_name: ", chatroomlist[i].Room_name, "online_num: ", chatroomlist[i].Online_num, "total_num: ", chatroomlist[i].Total_num)
	}
	//fmt.Println(chatroomlist) //显示聊天室列表
	fmt.Println("请选择聊天室序号：")
	fmt.Scanln(&ordinal)
	send := "List-" + uName + "-" + ordinal
	sendDataClient <- send
	res := <-recvDataClient
	if res == "SUCCESS" {
		fmt.Println("你已进入聊天室!")
		userList := <-recvDataClient
		json.Unmarshal([]byte(userList), &userInfo)
		for i, _ := range userInfo {
			fmt.Println("user_name: ", userInfo[i].User_name, "head_photo: ", userInfo[i].Head_photo, "status: ", userInfo[i].Status, "grade: ", userInfo[i].Grade)
		}
		//fmt.Println(userList) //显示玩家信息
		go showMessage() //启动聊天内容显示协程
	} else {
		fmt.Println(res)
	}
	inputMessage() //启动聊天内容输入
}

//聊天内容
func showMessage() {
	for {
		message := <-recvDataClient
		fmt.Println(message)
		if <-stopMessage {
			break
		}
	}
	return
}

//聊天内容输入
func inputMessage() {
	var input string
	for {
		fmt.Scanln(&input)
		switch input {
		case "quit":
			//stopChan <- true
			fmt.Println("再见")
			os.Exit(0)
		case "leave":
			stopMessage <- true
			chatRoom(hallInfo)
		default:
			if len(input) != 0 {
				input = uName + "说:" + input
			}
		}
		sendDataClient <- input
		input = ""
	}
}

//检查错误
func checkErrClient(err error) {
	if err != nil {
		log.Fatal("an error: ", err.Error())
	}
}
