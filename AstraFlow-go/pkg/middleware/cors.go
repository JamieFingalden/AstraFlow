package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CORSMiddleware CORS中间件
// 处理跨域请求，允许前端应用访问API
//
// CORS (Cross-Origin Resource Sharing) 是一种机制，它使用额外的 HTTP 头来告诉浏览器
// 让运行在一个 origin (domain) 上的Web应用被准许访问来自不同源服务器上的指定的资源
//
// 功能说明：
// 1. 设置允许的源域名（Access-Control-Allow-Origin）
// 2. 设置允许的HTTP方法（Access-Control-Allow-Methods）
// 3. 设置允许的请求头（Access-Control-Allow-Headers）
// 4. 处理预检请求（OPTIONS方法）
// 5. 设置允许携带认证信息（Access-Control-Allow-Credentials）
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 设置允许跨域请求的源，* 表示允许所有域名
		// 在生产环境中应该设置为具体的域名，如：http://localhost:3000 或 https://yourdomain.com
		c.Header("Access-Control-Allow-Origin", "*")

		// 设置允许的HTTP方法列表
		// GET: 获取资源
		// POST: 创建资源
		// PUT: 更新资源（完整更新）
		// DELETE: 删除资源
		// OPTIONS: 预检请求，浏览器在发送跨域请求前会先发送OPTIONS请求来检查是否允许
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		// 设置允许的请求头列表
		// Origin: 请求的源域名
		// Content-Type: 请求体的媒体类型
		// Content-Length: 请求体的长度
		// Accept-Encoding: 客户端能理解的内容编码方式
		// X-CSRF-Token: CSRF防护令牌
		// Authorization: 认证信息（如Bearer Token）
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		// 设置允许暴露给客户端的响应头
		// 这些响应头可以在前端通过 response.headers 获取
		c.Header("Access-Control-Expose-Headers", "Content-Length")

		// 设置是否允许携带认证信息（cookies、authorization headers等）
		// 当设置为true时，Access-Control-Allow-Origin 不能设置为 *
		// 这里为了开发方便设置为true，生产环境需要配合具体域名使用
		c.Header("Access-Control-Allow-Credentials", "true")

		// 处理预检请求（Preflight Request）
		// 浏览器在发送跨域请求前，会先发送一个OPTIONS请求来询问服务器是否允许该跨域请求
		// 如果是OPTIONS请求，直接返回204 No Content，表示允许跨域
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		// 继续执行后续的中间件和处理函数
		// 如果不是OPTIONS请求，继续处理正常的业务逻辑
		c.Next()
	}
}