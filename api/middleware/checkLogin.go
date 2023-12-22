package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"mall/api/httputils"
	"mall/internal/service"
	"net/http"
)

func CheckLogin(c *gin.Context) {
	token := c.GetHeader("mall-auth-token")
	if token == "" {
		c.AbortWithStatusJSON(http.StatusOK, httputils.Error(httputils.UserNotLogin))
		return
	}
	// 解析token
	userMap, err := service.ParseAPIToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, httputils.Error(httputils.UserNotLogin))
		return
	}

	userId := cast.ToInt(userMap["user_id"])
	// 设置上下文userId
	c.Set("userId", userId)
}
