package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"mall/api/httputils"
	"net/http"
	"runtime"
)

const size = 64 << 10

func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			//打印错误堆栈信息
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			log.Printf("http: panic serving err: %v\n%s", r, buf)
			c.JSON(http.StatusOK, httputils.Error(httputils.InterNalError))
			c.Abort()
		}
	}()
	// 必须手动调用c.Next()
	c.Next()
}
