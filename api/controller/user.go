package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"mall/api/httputils"
	"mall/internal/entity"
	"mall/internal/service"
	"net/http"
)

func APILogin(c *gin.Context) {
	req := entity.LoginReq{}
	err := c.ShouldBindJSON(&req)
	// 没有参数：ShouldBindJSON err:EOF
	if err != nil {
		c.JSON(http.StatusOK, httputils.Error(httputils.PARAM_ERROR))
		return
	}
	resp, err := service.APILogin(c, req)
	if err != nil {
		c.JSON(http.StatusOK, httputils.Error(err))
		return
	}
	c.JSON(http.StatusOK, httputils.SuccessWithData(resp))
}

func SetUserInfo(c *gin.Context) {
	var req = entity.SetUserInfoReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, httputils.Error(err))
		return
	}
	userId, _ := c.Get("userId")
	req.UserId = cast.ToInt(userId)
	err = service.SetUserInfo(c, req)
	if err != nil {
		c.JSON(http.StatusOK, httputils.Error(err))
		return
	}

	c.JSON(http.StatusOK, httputils.Success())
}
