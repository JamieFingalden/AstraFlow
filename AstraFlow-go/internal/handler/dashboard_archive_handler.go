package handler

import (
	"AstraFlow-go/internal/model"
	"AstraFlow-go/internal/service"
	typeUtils "AstraFlow-go/pkg/utils"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// DashboardArchiveHandler 统计与归档处理器
type DashboardArchiveHandler struct {
	invoiceService *service.InvoiceService
}

// NewDashboardArchiveHandler 创建统计与归档处理器实例
func NewDashboardArchiveHandler() *DashboardArchiveHandler {
	return &DashboardArchiveHandler{invoiceService: service.NewInvoiceService()}
}

// GetArchiveInvoices 历史归档高级搜索
// 入参：Query 参数 page、size、status、start_date、end_date、keyword；Context 中 user_id、role、tenant_id。
// 出参：统一响应结构，data 包含归档分页结果。
func (h *DashboardArchiveHandler) GetArchiveInvoices(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "page 参数必须为大于 0 的整数"})
		return
	}

	size, err := strconv.Atoi(c.DefaultQuery("size", "10"))
	if err != nil || size < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "size 参数必须为大于 0 的整数"})
		return
	}

	userValue, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "未授权"})
		return
	}

	userID, ok := typeUtils.AnyToInt64(userValue)
	if !ok || userID <= 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "用户身份无效"})
		return
	}

	roleValue, exists := c.Get("role")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "角色信息缺失"})
		return
	}

	role, ok := typeUtils.AnyToString(roleValue)
	if !ok || strings.TrimSpace(role) == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "角色信息无效"})
		return
	}

	status := strings.TrimSpace(c.Query("status"))
	keyword := strings.TrimSpace(c.Query("keyword"))

	var startDate *time.Time
	startDateStr := strings.TrimSpace(c.Query("start_date"))
	if startDateStr != "" {
		parsed, parseErr := time.Parse("2006-01-02", startDateStr)
		if parseErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "start_date 格式错误，应为 YYYY-MM-DD"})
			return
		}
		startDate = &parsed
	}

	var endDate *time.Time
	endDateStr := strings.TrimSpace(c.Query("end_date"))
	if endDateStr != "" {
		parsed, parseErr := time.Parse("2006-01-02", endDateStr)
		if parseErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "end_date 格式错误，应为 YYYY-MM-DD"})
			return
		}
		endDate = &parsed
	}

	var tenantID *int64
	if tenantValue, exists := c.Get("tenant_id"); exists && tenantValue != nil {
		if tid, convertOK := typeUtils.AnyToInt64(tenantValue); convertOK && tid != 0 {
			tenantID = &tid
		}
	}

	result, err := h.invoiceService.SearchArchiveInvoices(page, size, status, startDate, endDate, keyword, userID, role, tenantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "归档搜索失败: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "归档搜索成功",
		"data":    result,
	})
}

// GetDashboardStats 获取仪表盘统计数据
// 入参：Context 中 user_id、role、tenant_id。
// 出参：统一响应结构，data 为按角色口径计算后的统计面板。
func (h *DashboardArchiveHandler) GetDashboardStats(c *gin.Context) {
	userValue, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "未授权"})
		return
	}

	userID, ok := typeUtils.AnyToInt64(userValue)
	if !ok || userID <= 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "用户身份无效"})
		return
	}

	roleValue, exists := c.Get("role")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "角色信息缺失"})
		return
	}

	role, ok := typeUtils.AnyToString(roleValue)
	if !ok || strings.TrimSpace(role) == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "角色信息无效"})
		return
	}

	var tenantID *int64
	if tenantValue, exists := c.Get("tenant_id"); exists && tenantValue != nil {
		if tid, convertOK := typeUtils.AnyToInt64(tenantValue); convertOK && tid != 0 {
			tenantID = &tid
		}
	}

	stats, err := h.invoiceService.GetDashboardStats(userID, role, tenantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "获取统计数据失败: " + err.Error()})
		return
	}

	role = strings.TrimSpace(role)
	if role == model.RoleKeyEmployee {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "获取统计数据成功",
			"data": gin.H{
				"unconfirmed":             stats.Unconfirmed,
				"processing_amount":       stats.ProcessingAmount,
				"total_reimbursed_amount": stats.TotalReimbursedAmount,
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取统计数据成功",
		"data": gin.H{
			"company_pending_count": stats.CompanyPendingCount,
			"monthly_paid_amount":   stats.MonthlyPaidAmount,
			"company_to_pay_count":  stats.CompanyToPayCount,
		},
	})
}
