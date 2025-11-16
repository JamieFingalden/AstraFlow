package api

import (
	"AstraFlow-go/internal/database"
	"AstraFlow-go/internal/handler"
	"AstraFlow-go/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	// 初始化数据库
	database.InitDB()

	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/health", handler.Health)

		// 认证相关接口
		authHandler := handler.NewAuthHandler()
		v1.POST("/auth/register", authHandler.Register) // 用户注册
		v1.POST("/auth/login", authHandler.Login)       // 用户登录

		// 需要认证的接口
		protected := v1.Group("/")
		protected.Use(middleware.AuthMiddleware())
		{
			protected.POST("/auth/refresh", authHandler.RefreshToken) // 刷新Token
			protected.GET("/auth/me", authHandler.GetCurrentUser)    // 获取当前用户信息
		}
	}

	return r
}
