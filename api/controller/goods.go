package controller

import (
	"github.com/gin-gonic/gin"
	"mall/api/httputils"
	"mall/internal/entity"
	"mall/internal/logger"
	"mall/internal/service"
	"net/http"
)

func CreateGoods(c *gin.Context) {
	req := entity.CreateGoodsReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		logger.WithContext(c).Errorf("ShouldBindJSON err:%v", err)
		c.JSON(http.StatusOK, httputils.Error(httputils.PARAM_ERROR))
		return
	}

	resp, err := service.CreateGoods(c, req)
	if err != nil {
		c.JSON(http.StatusOK, httputils.Error(err))
		return
	}
	c.JSON(http.StatusOK, httputils.SuccessWithData(resp))
}

func UpdateGoods(c *gin.Context) {
	req := entity.UpdateGoodsReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		logger.WithContext(c).Errorf("ShouldBindJSON err:%v", err)
		c.JSON(http.StatusOK, httputils.Error(httputils.PARAM_ERROR))
		return
	}

	resp, err := service.UpdateGoods(c, req)
	if err != nil {
		c.JSON(http.StatusOK, httputils.Error(err))
		return
	}
	c.JSON(http.StatusOK, httputils.SuccessWithData(resp))
}

func DeleteGoods(c *gin.Context) {
	req := entity.DeleteGoodsReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		logger.WithContext(c).Errorf("ShouldBindJSON err:%v", err)
		c.JSON(http.StatusOK, httputils.Error(httputils.PARAM_ERROR))
		return
	}

	resp, err := service.DeleteGoods(c, req)
	if err != nil {
		c.JSON(http.StatusOK, httputils.Error(err))
		return
	}
	c.JSON(http.StatusOK, httputils.SuccessWithData(resp))
}

func GoodsList(c *gin.Context) {
	req := entity.GoodsListReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		logger.WithContext(c).Errorf("ShouldBindJSON err:%v", err)
		c.JSON(http.StatusOK, httputils.Error(httputils.PARAM_ERROR))
		return
	}

	resp, err := service.GoodsList(c, req)
	if err != nil {
		c.JSON(http.StatusOK, httputils.Error(err))
		return
	}
	c.JSON(http.StatusOK, httputils.SuccessWithData(resp))
}

func GetGoodsDetail(c *gin.Context) {
	req := entity.GetGoodsDetailReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		logger.WithContext(c).Errorf("ShouldBindJSON err:%v", err)
		c.JSON(http.StatusOK, httputils.Error(httputils.PARAM_ERROR))
		return
	}

	resp, err := service.GetGoodsDetail(c, req)
	if err != nil {
		c.JSON(http.StatusOK, httputils.Error(err))
		return
	}
	c.JSON(http.StatusOK, httputils.SuccessWithData(resp))
}
