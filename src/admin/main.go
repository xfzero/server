package main

import (
	"gm/admin/controller"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

type S2C_UserInfo struct {
	Code uint32 `json:"code"`
	User string `json:"user,omitempty"`
}

func main() {
	// // 1.创建路由
	// r := gin.Default()
	// // 2.绑定路由规则，执行的函数
	// // gin.Context，封装了request和response
	// r.GET("/", func(c *gin.Context) {
	// 	c.String(http.StatusOK, "hello World!")
	// })

	// ret := S2C_UserInfo{Code: 0, User: "ags"}
	// r.GET("/test1", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, ret)
	// })
	// // 3.监听端口，默认在8080
	// // Run("里面不指定端口号默认为8080")
	// r.Run(":8000")

	initGin()
}

func initGin() {
	//创建日志文件，os.Stdout同时将日志写入控制台
	f, _ := os.Create("../../gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r := gin.Default()

	// r.GET("/", func(c *gin.Context) {
	// 	c.String(http.StatusOK, "hello World!")
	// })

	controller.InitCommand(r)

	r.Run(":8000")
}
