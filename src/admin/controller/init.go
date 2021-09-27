package controller

import (
	"gm/admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitCommand(r *gin.Engine) {
	//CORS
	r.Use(middleware.CORS)

	r.POST("/login", UserLogin)

	auth := r.Use(middleware.Auth)
	{
		//账号管理
		auth.GET("/account/list", AccountList)
		auth.POST("account/update", UpdateAccount)

		//grpc-sayHello测试
		auth.GET("grpc/sayhello", SayHello)

		//grpc-gmcmd测试
		auth.GET("grpc/activity", CommonActivity)
	}

}
