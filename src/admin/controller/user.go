package controller

import (
	"fmt"
	"gm/admin/controller/input"
	"gm/admin/controller/output"
	"gm/admin/model"
	"gm/base"

	"github.com/gin-gonic/gin"
)

type S2C_LoginInfo struct {
	Code uint32        `json:"code"`
	Info base.UserInfo `json:"info"`
}

func UserLogin(c *gin.Context) {
	//json格式的不能这样接受
	// userName := c.PostForm("userName")
	// password := c.PostForm("password")

	//通过ShouldBind绑定到struct
	req := &input.C2S_Login{}
	if !input.CheckAndBind(c, req) {
		return
	}

	fmt.Println(req.Username, req.Password)

	user := model.GetUserLogin(req.Username, req.Password)

	info := base.UserInfo{
		Id:            user.Id,
		Username:      user.Username,
		LastLoginTime: user.LastLoginTime,
		Group:         user.Group,
		Token:         "123456",
	}

	output.Response(c, S2C_LoginInfo{
		Code: 0,
		Info: info,
	})

}
