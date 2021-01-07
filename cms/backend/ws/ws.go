package ws

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"strings"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Ping(c *gin.Context) {
	fmt.Println("websocket Ping")
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("localAddr:", ws.LocalAddr())
	fmt.Println("remoteAddr:", ws.RemoteAddr())
	defer ws.Close()

	for {
		mt, message, err := ws.ReadMessage()
		fmt.Println("mt:", mt, "message:", string(message), "err:", err)
		if err != nil {
			//if strings.Contains(err.Error(), "close 1001") {
			//	fmt.Println("执行解锁操作")
			//	return
			//}
			if err.Error() == "websocket: invalid close code" || strings.Contains(err.Error(), "close 1001") {
				fmt.Println("websocket close")
				return
			}
			fmt.Println("err:", err)
			return
		}
		fmt.Println("mt:", mt, "messageStr:", string(message))

		ws.SetCloseHandler(func(code int, text string) error {
			fmt.Println("code", code, "text", text)
			fmt.Println("执行解锁操作")
			return nil
		})

		if string(message) == "close" {
			message = []byte("websocket close")
			err = ws.WriteMessage(websocket.CloseMessage, message)
			if err != nil {
				fmt.Println("close failed:", err)
				break
			}
		}

		if string(message) == "ping" {
			message = []byte("pong")
			err = ws.WriteMessage(mt, message)
			if err != nil {
				fmt.Println("write failed:", err)
				break
			}
		}
		// 返回json字符串
		//v := gin.H{
		//	"message": message,
		//}
		//err = ws.WriteJSON(v)
		// 二进制返回
	}
}
