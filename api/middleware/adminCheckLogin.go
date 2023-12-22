package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"mall/api/httputils"
	"mall/internal/logger"
	"mall/internal/service"
	"net/http"
)

func AdminCheckLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("admin_token")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusOK, httputils.Error(httputils.UserNotLogin))
			return
		}

		// 解析token
		userMap, err := service.ParseAdminToken(token)
		if err != nil {
			logger.WithContext(c).Errorf("service.ParseAdminToken err:%v", err)
			c.AbortWithStatusJSON(http.StatusOK, httputils.Error(httputils.UserCheckLoginError))
			return
		}

		//
		adminId := cast.ToInt(userMap["admin_id"])
		c.Set("adminId", adminId)
	}
}
