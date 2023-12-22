package controller

import (
	"github.com/gin-gonic/gin"
	"mall/api/httputils"
	"mall/internal/entity"
	"mall/internal/service"
	"net/http"
)

func AdminLogin(c *gin.Context) {
	req := entity.AdminLoginReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, httputils.Error(httputils.PARAM_ERROR))
		return
	}
	resp, err := service.AdminLogin(c, req)
	if err != nil {
		c.JSON(http.StatusOK, httputils.Error(err))
		return
	}
	c.JSON(http.StatusOK, httputils.SuccessWithData(resp))
}
