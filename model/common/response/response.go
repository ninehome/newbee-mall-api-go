package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	ResultCode int         `json:"resultCode"`
	Data       interface{} `json:"data"`
	Msg        string      `json:"message"`
}

const (
	ERROR       = 500
	SUCCESS     = 200
	UNLOGIN     = 416
	Passworderr = 417
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func Ok(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "成功", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "成功", c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, "失敗 code=12", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, c)
}

// 密码错误
func FailWithPSW(message string, c *gin.Context) {
	Result(Passworderr, map[string]interface{}{}, message, c)
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(ERROR, data, message, c)
}

func UnLogin(data interface{}, c *gin.Context) {
	Result(UNLOGIN, data, "未登陸！code=11", c)
}
