package api

import (
	"AstraFlow-go/internal/handler"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/health", handler.Health)
	}

	return r
}
