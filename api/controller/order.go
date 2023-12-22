package controller

import (
	"github.com/gin-gonic/gin"
	"mall/api/httputils"
	"mall/internal/entity"
	"mall/internal/logger"
	"mall/internal/service"
	"net/http"
)

func CreateOrder(c *gin.Context) {
	req := entity.CreateOrderReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		logger.WithContext(c).Errorf("ShouldBindJSON err:%v", err)
		c.JSON(http.StatusOK, httputils.Error(httputils.PARAM_ERROR))
		return
	}

	resp, err := service.CreateOrder(c, req)
	if err != nil {
		c.JSON(http.StatusOK, httputils.Error(err))
		return
	}
	c.JSON(http.StatusOK, httputils.SuccessWithData(resp))
}

func CreateOrderV1(c *gin.Context) {
	req := entity.CreateOrderReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		logger.WithContext(c).Errorf("ShouldBindJSON err:%v", err)
		c.JSON(http.StatusOK, httputils.Error(httputils.PARAM_ERROR))
		return
	}

	resp, err := service.CreateOrderV1(c, req)
	if err != nil {
		c.JSON(http.StatusOK, httputils.Error(err))
		return
	}
	c.JSON(http.StatusOK, httputils.SuccessWithData(resp))
}

func CreateOrderV2(c *gin.Context) {
	req := entity.CreateOrderReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		logger.WithContext(c).Errorf("ShouldBindJSON err:%v", err)
		c.JSON(http.StatusOK, httputils.Error(httputils.PARAM_ERROR))
		return
	}

	resp, err := service.CreateOrderV2(c, req)
	if err != nil {
		c.JSON(http.StatusOK, httputils.Error(err))
		return
	}
	c.JSON(http.StatusOK, httputils.SuccessWithData(resp))
}

func OrderList(c *gin.Context) {
	req := entity.OrderListReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		logger.WithContext(c).Errorf("ShouldBindJSON err:%v", err)
		c.JSON(http.StatusOK, httputils.Error(httputils.PARAM_ERROR))
		return
	}

	req.UserId = c.GetInt("userId")
	resp, err := service.OrderList(c, req)
	if err != nil {
		c.JSON(http.StatusOK, httputils.Error(err))
		return
	}
	c.JSON(http.StatusOK, httputils.SuccessWithData(resp))
}

func GetOrderDetail(c *gin.Context) {
	req := entity.GetOrderDetailReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		logger.WithContext(c).Errorf("ShouldBindJSON err:%v", err)
		c.JSON(http.StatusOK, httputils.Error(httputils.PARAM_ERROR))
		return
	}

	resp, err := service.GetOrderDetail(c, req)
	if err != nil {
		c.JSON(http.StatusOK, httputils.Error(err))
		return
	}
	c.JSON(http.StatusOK, httputils.SuccessWithData(resp))
}
