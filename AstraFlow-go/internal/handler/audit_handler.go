package handler

import (
	"AstraFlow-go/internal/service"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuditHandler 财务审核处理器
type AuditHandler struct {
	invoiceService *service.InvoiceService
}

// NewAuditHandler 创建财务审核处理器
func NewAuditHandler() *AuditHandler {
	return &AuditHandler{
		invoiceService: service.NewInvoiceService(),
	}
}

// ReviewInvoiceRequest 审核请求体
type ReviewInvoiceRequest struct {
	Action      string   `json:"action" binding:"required,oneof=approve reject"`
	FinalAmount *float64 `json:"final_amount"`
	Remarks     string   `json:"remarks"`
}

// GetPendingInvoices
// [GET] /api/v1/audit/invoices/pending
// 获取待审核任务池（status = pending），支持分页并返回提交人姓名。
func (h *AuditHandler) GetPendingInvoices(c *gin.Context) {
	page, size, ok := parsePageAndSize(c)
	if !ok {
		return
	}

	result, err := h.invoiceService.GetPendingInvoices(page, size)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "获取待审核单据失败: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取待审核单据成功",
		"data":    result,
	})
}

// GetInvoiceDetail
// [GET] /api/v1/audit/invoices/:id
// 获取单据详情（只读，状态不限），包含提交人信息和图片 URL。
func (h *AuditHandler) GetInvoiceDetail(c *gin.Context) {
	id, ok := parsePathID(c)
	if !ok {
		return
	}

	detail, err := h.invoiceService.GetInvoiceDetailForAudit(id)
	if err != nil {
		if errors.Is(err, service.ErrInvoiceNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"code": 404, "message": "单据不存在"})
			return
		}
		if errors.Is(err, service.ErrInvalidInvoiceID) {
			c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "id 参数不合法"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "获取单据详情失败: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取单据详情成功",
		"data":    detail,
	})
}

// ReviewInvoice
// [PUT] /api/v1/audit/invoices/:id/review
// 财务审核批复：approve/reject，要求单据状态必须为 pending。
func (h *AuditHandler) ReviewInvoice(c *gin.Context) {
	id, ok := parsePathID(c)
	if !ok {
		return
	}

	var req ReviewInvoiceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "请求参数错误: " + err.Error()})
		return
	}

	actorID, ok := getCurrentUserID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "未授权"})
		return
	}

	result, err := h.invoiceService.ReviewInvoice(id, *actorID, req.Action, req.FinalAmount, strings.TrimSpace(req.Remarks))

	if err != nil {
		switch {
		case errors.Is(err, service.ErrInvoiceAlreadyHandled):
			c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "单据已被处理"})
			return
		case errors.Is(err, service.ErrFinalAmountRequired):
			c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "审批通过时 final_amount 必填"})
			return
		case errors.Is(err, service.ErrRejectRemarksRequired):
			c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "驳回时 remarks 不能为空"})
			return
		case errors.Is(err, service.ErrInvalidReviewAction):
			c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "action 只能是 approve 或 reject"})
			return
		case errors.Is(err, service.ErrInvoiceNotFound):
			c.JSON(http.StatusNotFound, gin.H{"code": 404, "message": "单据不存在"})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "审核处理失败: " + err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "审核处理成功",
		"data":    result,
	})
}

// parsePageAndSize 解析分页参数
func parsePageAndSize(c *gin.Context) (int, int, bool) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "page 参数必须为大于 0 的整数"})
		return 0, 0, false
	}

	size, err := strconv.Atoi(c.DefaultQuery("size", "10"))
	if err != nil || size < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "size 参数必须为大于 0 的整数"})
		return 0, 0, false
	}

	if size > 100 {
		size = 100
	}

	return page, size, true
}

// parsePathID 解析路径参数 ID
func parsePathID(c *gin.Context) (int64, bool) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "id 参数不合法"})
		return 0, false
	}
	return id, true
}

// getCurrentUserID 获取当前登录用户 ID
func getCurrentUserID(c *gin.Context) (*int64, bool) {
	v, exists := c.Get("user_id")
	if !exists {
		return nil, false
	}

	switch val := v.(type) {
	case int64:
		return &val, true
	case int:
		id := int64(val)
		return &id, true
	case float64:
		id := int64(val)
		return &id, true
	case string:
		id, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return nil, false
		}
		return &id, true
	default:
		return nil, false
	}
}
