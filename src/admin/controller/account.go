package controller

import (
	"gm/admin/controller/output"

	"github.com/gin-gonic/gin"
)

type S2C_UserList struct {
	Code uint32   `json:"code"`
	List []string `json:"list"`
}

func AccountList(c *gin.Context) {
	list := []string{"tony", "alex"}

	output.Response(c, S2C_UserList{
		Code: 0,
		List: list,
	})
}

func UpdateAccount(c *gin.Context) {

}
