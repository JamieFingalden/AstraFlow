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

	// 静态文件服务 - 提供上传的文件访问
	r.Static("/uploads", "./uploads")

	v1 := r.Group("/api/v1")
	{
		v1.GET("/health", handler.Health)

		// 认证相关接口
		authHandler := handler.NewAuthHandler()
		// v1.POST("/auth/register", authHandler.Register) // 用户注册
		v1.POST("/auth/login", authHandler.Login) // 用户登录
		tenantHandler := handler.NewTenantHandler()
		v1.POST("/tenants", tenantHandler.CreateTenant) // 创建租户

		// 需要认证的接口
		protected := v1.Group("/")
		protected.Use(middleware.AuthMiddleware())
		{
			protected.POST("/auth/refresh", authHandler.RefreshToken) // 刷新Token
			protected.GET("/auth/me", authHandler.GetCurrentUser)     // 获取当前用户信息

			// 用户管理相关接口
			userHandler := handler.NewUserHandler()
			protected.GET("/users", userHandler.GetUserList)                        // 获取用户列表 获取员工列表
			protected.POST("/users", userHandler.CreateUser)                        // 创建用户 添加员工
			protected.PUT("/users/:id", userHandler.UpdateUser)                     // 更新用户 更新员工
			protected.DELETE("/users/:id", userHandler.DeleteUser)                  // 删除用户 停用或删除员工
			protected.PATCH("/users/:id/reset-password", userHandler.ResetPassword) // 重置员工密码
			protected.GET("/users/:id", userHandler.GetUserById)                    // 获取用户详情

			// 租户相关接口
			protected.GET("/tenants", tenantHandler.GetTenantList)       // 获取租户列表
			protected.GET("/tenants/:id", tenantHandler.GetTenant)       // 获取租户详情
			protected.PUT("/tenants/:id", tenantHandler.UpdateTenant)    // 更新租户
			protected.DELETE("/tenants/:id", tenantHandler.DeleteTenant) // 删除租户

			// 附件上传相关接口
			attachmentHandler := handler.NewAttachmentHandler()
			protected.POST("/attachments", attachmentHandler.UploadFile)                                   // 上传文件
			protected.GET("/attachments/:id", attachmentHandler.GetAttachmentByID)                         // 根据ID获取附件
			protected.GET("/attachments", attachmentHandler.GetAttachmentsByUserID)                        // 获取用户附件列表
			protected.GET("/attachments/tenant", attachmentHandler.GetAttachmentsByTenantID)               // 获取租户附件列表
			protected.GET("/attachments/invoice/:invoice_id", attachmentHandler.GetAttachmentsByInvoiceID) // 根据发票ID获取附件列表
			protected.GET("/attachments/status", attachmentHandler.GetAttachmentsByStatus)                 // 根据状态获取附件列表
			protected.DELETE("/attachments/:id", attachmentHandler.DeleteAttachment)                       // 删除附件

			// OCR结果回调接口 - 不需要认证，由Python端直接调用
			v1.POST("/callback/ocr-result", attachmentHandler.HandleOCRResultCallback) // 处理OCR结果回调

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
			protected.GET("/analytics/monthly", analyticsHandler.GetMonthlyExpenses)                          // 每月支出趋势
			protected.GET("/analytics/recent-bills", analyticsHandler.GetRecentBills)                         // 近期账单
			protected.GET("/analytics/reimbursement", analyticsHandler.GetReimbursementStatisticsData)        // 报销统计数据
			protected.GET("/analytics/reimbursement/statistics", analyticsHandler.GetReimbursementStatistics) // 报销统计详情
			protected.GET("/analytics/reimbursement/trends", analyticsHandler.GetReimbursementMonthlyTrends)  // 报销月度趋势
		}
	}

	return r
}
