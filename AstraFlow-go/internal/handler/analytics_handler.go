package handler

import (
	"AstraFlow-go/internal/service"
	typeUtils "AstraFlow-go/pkg/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// AnalyticsHandler handles analytics-related HTTP requests
type AnalyticsHandler struct {
	analyticsService *service.AnalyticsService
}

// NewAnalyticsHandler creates a new analytics handler instance
func NewAnalyticsHandler() *AnalyticsHandler {
	return &AnalyticsHandler{
		analyticsService: service.NewAnalyticsService(),
	}
}

// AnalyticsResponse represents the standard response format for analytics endpoints
type AnalyticsResponse struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data,omitempty"`
}

// GetDashboardData retrieves all dashboard analytics data
func (h AnalyticsHandler) GetDashboardData(c *gin.Context) {
	// Extract user and tenant information from context
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, AnalyticsResponse{
			Code:    401,
			Message: "登录过期, 请重新登录",
		})
		return
	}

	userIdInt, ok := typeUtils.AnyToInt64(userID)
	if !ok {
		c.JSON(http.StatusInternalServerError, AnalyticsResponse{
			Code:    500,
			Message: "user_id类型转换错误",
		})
		return
	}

	tenantID, tenantExists := c.Get("tenant_id")
	var tenantIdInt *int64
	if !tenantExists || tenantID == nil {
		tenantIdInt = nil
	} else {
		if tid, ok := tenantID.(*int64); ok && tid != nil {
			tenantIdInt = tid
		} else {
			tenantIdInt = nil
		}
	}

	now := time.Now()            // 当前时间
	year, month, _ := now.Date() // 取出年、月
	// 本月第一天 00:00:00
	firstOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, now.Location())
	// 下月第一天 00:00:00 再减 1 纳秒 → 本月最后 1 秒
	lastOfMonth := time.Date(year, month+1, 1, 0, 0, 0, -1, now.Location())

	startDate := &firstOfMonth
	endDate := &lastOfMonth

	// Get dashboard data from service
	dashboardData, err := h.analyticsService.GetDashboardData(tenantIdInt, userIdInt, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, AnalyticsResponse{
			Code:    500,
			Message: "获取仪表板数据失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, AnalyticsResponse{
		Code:    200,
		Message: "获取仪表板数据成功",
		Data: map[string]interface{}{
			"dashboard": dashboardData,
		},
	})
}

// GetMetrics retrieves key dashboard metrics only
func (h AnalyticsHandler) GetMetrics(c *gin.Context) {
	// Extract user and tenant information from context
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, AnalyticsResponse{
			Code:    401,
			Message: "登录过期, 请重新登录",
		})
		return
	}

	userIdInt, ok := typeUtils.AnyToInt64(userID)
	if !ok {
		c.JSON(http.StatusInternalServerError, AnalyticsResponse{
			Code:    500,
			Message: "user_id类型转换错误",
		})
		return
	}

	tenantID, tenantExists := c.Get("tenant_id")
	var tenantIdInt *int64
	if !tenantExists || tenantID == nil {
		tenantIdInt = nil
	} else {
		if tid, ok := tenantID.(*int64); ok && tid != nil {
			tenantIdInt = tid
		} else {
			tenantIdInt = nil
		}
	}

	// Parse query parameters
	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")

	var startDate, endDate *time.Time

	if startDateStr != "" {
		if parsedDate, err := time.Parse("2006-01-02", startDateStr); err == nil {
			startDate = &parsedDate
		}
	}

	if endDateStr != "" {
		if parsedDate, err := time.Parse("2006-01-02", endDateStr); err == nil {
			endDate = &parsedDate
		}
	}

	// Get metrics from service
	metrics, err := h.analyticsService.GetMetrics(tenantIdInt, userIdInt, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, AnalyticsResponse{
			Code:    500,
			Message: "获取指标数据失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, AnalyticsResponse{
		Code:    200,
		Message: "获取指标数据成功",
		Data: map[string]interface{}{
			"metrics": metrics,
		},
	})
}

// GetExpenseCategories retrieves expense breakdown by category
func (h AnalyticsHandler) GetExpenseCategories(c *gin.Context) {
	// Extract user and tenant information from context
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, AnalyticsResponse{
			Code:    401,
			Message: "登录过期, 请重新登录",
		})
		return
	}

	userIdInt, ok := typeUtils.AnyToInt64(userID)
	if !ok {
		c.JSON(http.StatusInternalServerError, AnalyticsResponse{
			Code:    500,
			Message: "user_id类型转换错误",
		})
		return
	}

	tenantID, tenantExists := c.Get("tenant_id")
	var tenantIdInt *int64
	if !tenantExists || tenantID == nil {
		tenantIdInt = nil
	} else {
		if tid, ok := tenantID.(*int64); ok && tid != nil {
			tenantIdInt = tid
		} else {
			tenantIdInt = nil
		}
	}

	// Parse query parameters
	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")

	var startDate, endDate *time.Time

	if startDateStr != "" {
		if parsedDate, err := time.Parse("2006-01-02", startDateStr); err == nil {
			startDate = &parsedDate
		}
	}

	if endDateStr != "" {
		if parsedDate, err := time.Parse("2006-01-02", endDateStr); err == nil {
			endDate = &parsedDate
		}
	}

	// Get expense categories from service
	categories, err := h.analyticsService.GetExpenseCategories(tenantIdInt, userIdInt, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, AnalyticsResponse{
			Code:    500,
			Message: "获取支出类别数据失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, AnalyticsResponse{
		Code:    200,
		Message: "获取支出类别数据成功",
		Data: map[string]interface{}{
			"expense_categories": categories,
		},
	})
}

// GetWeeklyExpenses retrieves weekly expense trends
func (h AnalyticsHandler) GetWeeklyExpenses(c *gin.Context) {
	// Extract user and tenant information from context
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, AnalyticsResponse{
			Code:    401,
			Message: "登录过期, 请重新登录",
		})
		return
	}

	userIdInt, ok := typeUtils.AnyToInt64(userID)
	if !ok {
		c.JSON(http.StatusInternalServerError, AnalyticsResponse{
			Code:    500,
			Message: "user_id类型转换错误",
		})
		return
	}

	tenantID, tenantExists := c.Get("tenant_id")
	var tenantIdInt *int64
	if !tenantExists || tenantID == nil {
		tenantIdInt = nil
	} else {
		if tid, ok := tenantID.(*int64); ok && tid != nil {
			tenantIdInt = tid
		} else {
			tenantIdInt = nil
		}
	}

	// Parse query parameters
	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")

	var startDate, endDate *time.Time

	if startDateStr != "" {
		if parsedDate, err := time.Parse("2006-01-02", startDateStr); err == nil {
			startDate = &parsedDate
		}
	}

	if endDateStr != "" {
		if parsedDate, err := time.Parse("2006-01-02", endDateStr); err == nil {
			endDate = &parsedDate
		}
	}

	// Get monthly expenses from service
	monthlyExpenses, err := h.analyticsService.GetMonthlyExpenses(tenantIdInt, userIdInt, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, AnalyticsResponse{
			Code:    500,
			Message: "获取每月支出数据失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, AnalyticsResponse{
		Code:    200,
		Message: "获取每月支出数据成功",
		Data: map[string]interface{}{
			"monthly_expenses": monthlyExpenses,
		},
	})
}

// GetRecentBills retrieves recent bills
func (h AnalyticsHandler) GetRecentBills(c *gin.Context) {
	// Extract user and tenant information from context
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, AnalyticsResponse{
			Code:    401,
			Message: "登录过期, 请重新登录",
		})
		return
	}

	userIdInt, ok := typeUtils.AnyToInt64(userID)
	if !ok {
		c.JSON(http.StatusInternalServerError, AnalyticsResponse{
			Code:    500,
			Message: "user_id类型转换错误",
		})
		return
	}

	tenantID, tenantExists := c.Get("tenant_id")
	var tenantIdInt *int64
	if !tenantExists || tenantID == nil {
		tenantIdInt = nil
	} else {
		if tid, ok := tenantID.(*int64); ok && tid != nil {
			tenantIdInt = tid
		} else {
			tenantIdInt = nil
		}
	}

	// Parse query parameters
	limitParam := c.DefaultQuery("limit", "10")
	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		limit = 10
	}

	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")

	var startDate, endDate *time.Time

	if startDateStr != "" {
		if parsedDate, err := time.Parse("2006-01-02", startDateStr); err == nil {
			startDate = &parsedDate
		}
	}

	if endDateStr != "" {
		if parsedDate, err := time.Parse("2006-01-02", endDateStr); err == nil {
			endDate = &parsedDate
		}
	}

	// Get recent bills from service
	bills, err := h.analyticsService.GetRecentBills(tenantIdInt, userIdInt, limit, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, AnalyticsResponse{
			Code:    500,
			Message: "获取近期账单数据失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, AnalyticsResponse{
		Code:    200,
		Message: "获取近期账单数据成功",
		Data: map[string]interface{}{
			"recent_bills": bills,
		},
	})
}

// GetReimbursementStatisticsData retrieves all reimbursement statistics data
func (h AnalyticsHandler) GetReimbursementStatisticsData(c *gin.Context) {
	// Extract user and tenant information from context
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, AnalyticsResponse{
			Code:    401,
			Message: "登录过期, 请重新登录",
		})
		return
	}

	userIdInt, ok := typeUtils.AnyToInt64(userID)
	if !ok {
		c.JSON(http.StatusInternalServerError, AnalyticsResponse{
			Code:    500,
			Message: "user_id类型转换错误",
		})
		return
	}

	tenantID, tenantExists := c.Get("tenant_id")
	var tenantIdInt *int64
	if !tenantExists || tenantID == nil {
		tenantIdInt = nil
	} else {
		if tid, ok := tenantID.(*int64); ok && tid != nil {
			tenantIdInt = tid
		} else {
			tenantIdInt = nil
		}
	}

	now := time.Now()            // 当前时间
	year, month, _ := now.Date() // 取出年、月
	// 本月第一天 00:00:00
	firstOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, now.Location())
	// 下月第一天 00:00:00 再减 1 纳秒 → 本月最后 1 秒
	lastOfMonth := time.Date(year, month+1, 1, 0, 0, 0, -1, now.Location())

	startDate := &firstOfMonth
	endDate := &lastOfMonth

	// Get reimbursement statistics data from service
	reimbursementData, err := h.analyticsService.GetReimbursementStatisticsData(tenantIdInt, userIdInt, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, AnalyticsResponse{
			Code:    500,
			Message: "获取报销统计数据分析失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, AnalyticsResponse{
		Code:    200,
		Message: "获取报销统计数据分析成功",
		Data: map[string]interface{}{
			"reimbursement_statistics": reimbursementData,
		},
	})
}

// GetReimbursementStatistics retrieves key reimbursement statistics only
func (h AnalyticsHandler) GetReimbursementStatistics(c *gin.Context) {
	// Extract user and tenant information from context
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, AnalyticsResponse{
			Code:    401,
			Message: "登录过期, 请重新登录",
		})
		return
	}

	userIdInt, ok := typeUtils.AnyToInt64(userID)
	if !ok {
		c.JSON(http.StatusInternalServerError, AnalyticsResponse{
			Code:    500,
			Message: "user_id类型转换错误",
		})
		return
	}

	tenantID, tenantExists := c.Get("tenant_id")
	var tenantIdInt *int64
	if !tenantExists || tenantID == nil {
		tenantIdInt = nil
	} else {
		if tid, ok := tenantID.(*int64); ok && tid != nil {
			tenantIdInt = tid
		} else {
			tenantIdInt = nil
		}
	}

	// Parse query parameters
	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")

	var startDate, endDate *time.Time

	if startDateStr != "" {
		if parsedDate, err := time.Parse("2006-01-02", startDateStr); err == nil {
			startDate = &parsedDate
		}
	}

	if endDateStr != "" {
		if parsedDate, err := time.Parse("2006-01-02", endDateStr); err == nil {
			endDate = &parsedDate
		}
	}

	// Get reimbursement statistics from service
	statistics, err := h.analyticsService.GetReimbursementStatistics(tenantIdInt, userIdInt, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, AnalyticsResponse{
			Code:    500,
			Message: "获取报销统计数据失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, AnalyticsResponse{
		Code:    200,
		Message: "获取报销统计数据成功",
		Data: map[string]interface{}{
			"reimbursement_statistics": statistics,
		},
	})
}

// GetReimbursementMonthlyTrends retrieves monthly reimbursement trends only
func (h AnalyticsHandler) GetReimbursementMonthlyTrends(c *gin.Context) {
	// Extract user and tenant information from context
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, AnalyticsResponse{
			Code:    401,
			Message: "登录过期, 请重新登录",
		})
		return
	}

	userIdInt, ok := typeUtils.AnyToInt64(userID)
	if !ok {
		c.JSON(http.StatusInternalServerError, AnalyticsResponse{
			Code:    500,
			Message: "user_id类型转换错误",
		})
		return
	}

	tenantID, tenantExists := c.Get("tenant_id")
	var tenantIdInt *int64
	if !tenantExists || tenantID == nil {
		tenantIdInt = nil
	} else {
		if tid, ok := tenantID.(*int64); ok && tid != nil {
			tenantIdInt = tid
		} else {
			tenantIdInt = nil
		}
	}

	// Parse query parameters
	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")

	var startDate, endDate *time.Time

	if startDateStr != "" {
		if parsedDate, err := time.Parse("2006-01-02", startDateStr); err == nil {
			startDate = &parsedDate
		}
	}

	if endDateStr != "" {
		if parsedDate, err := time.Parse("2006-01-02", endDateStr); err == nil {
			endDate = &parsedDate
		}
	}

	// Get monthly trends from service
	trends, err := h.analyticsService.GetReimbursementMonthlyTrends(tenantIdInt, userIdInt, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, AnalyticsResponse{
			Code:    500,
			Message: "获取报销月度趋势数据失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, AnalyticsResponse{
		Code:    200,
		Message: "获取报销月度趋势数据成功",
		Data: map[string]interface{}{
			"monthly_trends": trends,
		},
	})
}

// GetMonthlyExpenses retrieves monthly expense trends for the main dashboard
func (h AnalyticsHandler) GetMonthlyExpenses(c *gin.Context) {
	// Extract user and tenant information from context
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, AnalyticsResponse{
			Code:    401,
			Message: "登录过期, 请重新登录",
		})
		return
	}

	userIdInt, ok := typeUtils.AnyToInt64(userID)
	if !ok {
		c.JSON(http.StatusInternalServerError, AnalyticsResponse{
			Code:    500,
			Message: "user_id类型转换错误",
		})
		return
	}

	tenantID, tenantExists := c.Get("tenant_id")
	var tenantIdInt *int64
	if !tenantExists || tenantID == nil {
		tenantIdInt = nil
	} else {
		if tid, ok := tenantID.(*int64); ok && tid != nil {
			tenantIdInt = tid
		} else {
			tenantIdInt = nil
		}
	}

	now := time.Now()            // 当前时间
	year, month, _ := now.Date() // 取出年、月
	// 本月第一天 00:00:00
	firstOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, now.Location())
	// 下月第一天 00:00:00 再减 1 纳秒 → 本月最后 1 秒
	lastOfMonth := time.Date(year, month+1, 1, 0, 0, 0, -1, now.Location())

	startDate := &firstOfMonth
	endDate := &lastOfMonth

	// Get monthly expenses from service
	monthlyExpenses, err := h.analyticsService.GetMonthlyExpenses(tenantIdInt, userIdInt, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, AnalyticsResponse{
			Code:    500,
			Message: "获取每月支出数据失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, AnalyticsResponse{
		Code:    200,
		Message: "获取每月支出数据成功",
		Data: map[string]interface{}{
			"monthly_expenses": monthlyExpenses,
		},
	})
}
