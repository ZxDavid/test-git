package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*{
	code: 2001,
	data: xxx,
	msg: xx
}*/

func Response(c *gin.Context, httpStatus int, code int, data gin.H, msg string) {
	c.JSON(httpStatus, gin.H{"code": code, "data": data, "msg": msg})
}

//成功
func Success(c *gin.Context, data gin.H, msg string) {
	Response(c, http.StatusOK, 200, data, msg)
}

//失败
func Fail(c *gin.Context, msg string, data gin.H) {
	Response(c, http.StatusOK, 400, data, msg)
}
