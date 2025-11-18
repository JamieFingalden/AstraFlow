package middleware

import (
	"AstraFlow-go/pkg/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware JWT认证中间件
// 验证请求中的JWT Token，提取用户信息并存储到上下文中
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		// 检查Authorization header是否存在
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    401,
				"message": "缺少授权Token",
			})
			c.Abort()
			return
		}

		// 移除 "Bearer " 前缀
		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		}

		// 解析JWT Token
		claims, err := jwt.ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    401,
				"message": "Token无效或已过期",
			})
			c.Abort()
			return
		}

		// 将用户信息存储到上下文中
		// 后续处理函数可以通过c.Get("user_id")等方式获取用户信息
		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)
		c.Set("role", claims.Role)

		// 继续执行后续处理函数
		c.Next()
	}
}
