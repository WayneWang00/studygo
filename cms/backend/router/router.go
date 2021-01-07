package router

import (
	"Wayne/cms/backend/ws"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Init() {
	fmt.Println("router Init")
	router := gin.Default()
	//router.LoadHTMLFiles("./index.html")
	//router.GET("/", func(c *gin.Context) {
	//	c.HTML(200, "index.html", nil)
	//})
	router.GET("/ws", ws.Ping)
	router.Run("localhost:8000")
}
