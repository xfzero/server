package output

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response(c *gin.Context, ret interface{}) {
	c.JSON(http.StatusOK, ret)
}

func Ok(c *gin.Context, message string) {
	//model.AddOpLog(c)
	// Response(c, cs.S2C_Default{
	// 	Code: errmsg.CodeOk,
	// 	Message: message,
	// })
}
