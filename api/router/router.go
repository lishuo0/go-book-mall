package router

import (
	"errors"
	"github.com/gin-gonic/gin"
	"mall/api/controller"
	"mall/api/middleware"
)

func RegisterRouter(router *gin.Engine) {
	router.Use(middleware.Qps, middleware.Delay)

	router.GET("/metrics", controller.Metrics)

	router.Use(middleware.Context, middleware.AccessLogger)

	// 管理后台 admin 接口
	admin := router.Group("/admin")
	RegisterAdminRouter(admin)

	// 对外API接口
	api := router.Group("/api")
	RegisterApiRouter(api)
}

// 控制面
func RegisterAdminRouter(router *gin.RouterGroup) {
	//--------------不需要校验登录接口--------------
	//router.POST("/user/login", controller.AdminLogin)

	//--------------校验登录接口-----------------
	//router.Use(middleware.AdminCheckLogin())

	//商品相关接口
	router.POST("/goods", controller.CreateGoods)
	router.PUT("/goods", controller.UpdateGoods)
	router.DELETE("/goods", controller.DeleteGoods)
	router.GET("/goods/detail", controller.GetGoodsDetail)
	router.GET("/goods/list", controller.GoodsList)
}

func RegisterApiRouter(router *gin.RouterGroup) {
	router.Any("/healthCheck", controller.HealthCheckV1)

	router.Any("/panic", middleware.Recover, func(c *gin.Context) {
		panic(errors.New("this is a panic"))
		return
	})

	router.POST("/users/login", controller.APILogin)

	//router.Use(middleware.CheckLogin)
	router.POST("/users", controller.SetUserInfo)

	//商品相关接口
	router.POST("/goods/detail", controller.GetGoodsDetail)
	router.POST("/goods/list", controller.GoodsList)

	//订单相关接口
	router.POST("/order", controller.CreateOrder)
	router.POST("/order/v1", controller.CreateOrderV1)
	router.POST("/order/v2", controller.CreateOrderV2)
	router.POST("/order/list", controller.OrderList)
	router.POST("/order/detail", controller.GetOrderDetail)

}
