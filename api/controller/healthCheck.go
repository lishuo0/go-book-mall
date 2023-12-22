package controller

import (
	"github.com/gin-gonic/gin"
	"mall/api/httputils"
	"mall/internal/logger"
	"mall/internal/service"
	"net/http"
)

func HealthCheck(c *gin.Context) {
	logger.WithContext(c).Info("handle http request, healthCheck controller")
	resp, _ := service.HealthCheck(c)
	c.JSON(http.StatusOK, httputils.SuccessWithData(resp))
}

func HealthCheckV1(c *gin.Context) {
	logger.WithGoID().Info("handle http request, healthCheck controller")
	resp, _ := service.HealthCheckV1()
	c.JSON(http.StatusOK, httputils.SuccessWithData(resp))
}
