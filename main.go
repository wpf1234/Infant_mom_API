package main

import (
	"app/m/base"
	"app/m/handler"
	"app/m/middleware"
	"app/m/utils"
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

	//每天零点执行的函数
	go utils.StartTimer(handler.ChangeRec)
	go utils.StartTimer(handler.ChangeNews)
	go utils.StartTimer(handler.ChangePackets)

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
		// 主页，不需要token 也能访问
		r1.GET("/home", g.Home)
		// 查询
		r1.GET("/search", g.Search)
		// 获取详情
		r1.GET("/detail", g.GetDetails)
	}

	r2 := router.Group("/v1/infant/auth")
	r2.Use(middleware.JWTAuth())
	// 以下接口需要经过 token 认证才能访问
	{
		r2.POST("/upload", g.UploadImage)
		r2.GET("/detail", g.GetDetails)
		// 刷新 token之值
		r2.GET("/refresh", g.Refresh)
		r2.GET("/mine", g.Mine)
		// 我的页面相关接口
		mr := r2.Group("/mine")
		{
			// 我的订单
			mr.GET("/order", g.MyOrder)
			// 我的收货地址
			mr.GET("/address", g.GetAddress)
			mr.POST("/address", g.PostAddress)
			mr.PUT("/address", g.PutAddress)
			mr.DELETE("/address", g.DelAddress)
			// 我的记录
			mr.GET("/record", g.Record)
			// 我的卡券
			mr.GET("/packets", g.Packets)
			// 我的钱包
		}
		r2.GET("/cart", g.GetCart)
		r2.POST("/cart", g.AddCart)
		r2.PUT("/cart", g.PutCart)
		r2.DELETE("/cart", g.DelCart)
		// 更改商品状态，包括 关注/取消，喜爱/取消，浏览/删除
		r2.PUT("/change", g.ChangeState)
	}
	// 注册 handler
	service.Handle("/", router)

	err = service.Run()
	if err != nil {
		log.Error("服务启动失败: ", err)
		return
	}
}
