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

	// 添加CORS中间件
	r.Use(middleware.CORSMiddleware())

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
			protected.GET("/auth/me", authHandler.GetCurrentUser)     // 获取当前用户信息

			// 租户相关接口
			tenantHandler := handler.NewTenantHandler()
			protected.POST("/tenants", tenantHandler.CreateTenant)       // 创建租户
			protected.GET("/tenants", tenantHandler.GetTenantList)       // 获取租户列表
			protected.GET("/tenants/:id", tenantHandler.GetTenant)       // 获取租户详情
			protected.PUT("/tenants/:id", tenantHandler.UpdateTenant)    // 更新租户
			protected.DELETE("/tenants/:id", tenantHandler.DeleteTenant) // 删除租户

			invoiceHandler := handler.NewInvoiceHandler()
			protected.POST("/invoices", invoiceHandler.CreateInvoice)                      // 创建发票
			protected.GET("/invoices", invoiceHandler.GetAllInvoicePage)                   // 分页获取发票列表
			protected.GET("/invoicesByUser", invoiceHandler.GetAllInvoicePageByUserId)     // 根据用户id分页获取发票列表
			protected.GET("/invoicesByTenant", invoiceHandler.GetAllInvoicePageByTenantId) // 根据租户id分页获取发票列表
			protected.PUT("/invoices/:id", invoiceHandler.UpdateInvoice)                   // 更新发票信息
			protected.DELETE("/invoices/:id", invoiceHandler.DeleteInvoice)                // 删除发票

			// 分析和报告相关接口
			analyticsHandler := handler.NewAnalyticsHandler()
			protected.GET("/analytics/dashboard", analyticsHandler.GetDashboardData)                          // 仪表板数据
			protected.GET("/analytics/metrics", analyticsHandler.GetMetrics)                                  // 关键指标
			protected.GET("/analytics/categories", analyticsHandler.GetExpenseCategories)                     // 支出类别
			protected.GET("/analytics/weekly", analyticsHandler.GetWeeklyExpenses)                            // 每周支出趋势
			protected.GET("/analytics/recent-bills", analyticsHandler.GetRecentBills)                         // 近期账单
			protected.GET("/analytics/reimbursement", analyticsHandler.GetReimbursementStatisticsData)        // 报销统计数据
			protected.GET("/analytics/reimbursement/statistics", analyticsHandler.GetReimbursementStatistics) // 报销统计详情
			protected.GET("/analytics/reimbursement/trends", analyticsHandler.GetReimbursementMonthlyTrends)  // 报销月度趋势
		}
	}

	return r
}
