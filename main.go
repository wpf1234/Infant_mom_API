package main

import (
	"app/m/base"
	"app/m/handler"
	"app/m/middleware"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/web"
	log "github.com/sirupsen/logrus"
)

func main() {
	service := web.NewService(
		web.Name("infant.micro.api.v1.infant"),
		web.Version("latest"),
		web.Address(":12345"),
	)
	err := service.Init()
	if err != nil {
		log.Error("服务初始化失败: ", err)
		return
	}

	base.Init()

	// 全局设置环境，debug 为开发环境，线上环境为 gin.ReleaseMode
	gin.SetMode(gin.ReleaseMode)
	// 创建 Restful handler
	g := new(handler.Gin)
	router := gin.Default()
	router.Use(middleware.Cors())

	// 登录与注册接口不需要token认证
	r1 := router.Group("/v1/infant")
	{
		r1.POST("/login", g.Login)
		r1.POST("/register", g.Register)
	}

	r2 := router.Group("/v1/infant/auth")
	r2.Use(middleware.JWTAuth())
	// 以下接口需要经过 token 认证才能访问
	{
		// 刷新 token之值
		r2.GET("/refresh", g.Refresh)
		r2.GET("/mine", g.Mine)
		r2.GET("/order", g.MyOrder)
		r2.GET("/address", g.GetAddress)
		r2.POST("/address", g.PostAddress)
		r2.PUT("/address", g.PutAddress)
		r2.DELETE("/address", g.DelAddress)
		r2.POST("/image", g.UploadImage)
		// 更改商品状态，包括 关注/取消，喜爱/取消，浏览/删除
		r2.GET("/change", g.ChangeState)
	}
	// 注册 handler
	service.Handle("/", router)

	err = service.Run()
	if err != nil {
		log.Error("服务启动失败: ", err)
		return
	}
}
