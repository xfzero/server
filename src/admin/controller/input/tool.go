package input

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

//路由权限校验及输入参数解析绑定
func CheckAndBind(c *gin.Context, obj interface{}) bool {
	//参数解析
	var err error
	if c.Request.Method == http.MethodGet {
		err = c.ShouldBind(obj)
	} else {
		err = c.ShouldBindBodyWith(obj, binding.JSON)
	}

	if err != nil {
		return false
	}

	//权限验证done
	return true

	//return false
}
