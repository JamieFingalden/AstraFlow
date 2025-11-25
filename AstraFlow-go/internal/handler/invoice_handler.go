package handler

import (
	"AstraFlow-go/internal/service"
	typeUtils "AstraFlow-go/pkg/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// InvoiceHandler 发票处理器
type InvoiceHandler struct {
	invoiceService *service.InvoiceService // 发票服务实例
}

// NewInvoiceHandler 创建发票处理器实例
func NewInvoiceHandler() *InvoiceHandler {
	return &InvoiceHandler{
		invoiceService: service.NewInvoiceService(),
	}
}

// CreateInvoiceRequest 创建发票请求体
type CreateInvoiceRequest struct {
	TenantID      *int64     `json:"tenant_id,omitempty"`
	UserId        int64      `json:"user_id"`
	InvoiceNumber *string    `json:"invoice_number" binding:"required"`
	InvoiceDate   *time.Time `json:"invoice_date" binding:"required"`
	Amount        *float64   `json:"amount" binding:"required"`
	Vendor        *string    `json:"vendor" binding:"required"`
	TaxID         *string    `json:"taxId"`
	Category      *string    `json:"category" binding:"required"`
	PaymentSource *string    `json:"payment_source"`
	Status        string     `json:"status"`
}

// InvoiceResponse 发票响应体
type InvoiceResponse struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data,omitempty"`
}

// CreateInvoice 创建发票接口
func (h InvoiceHandler) CreateInvoice(c *gin.Context) {
	var req CreateInvoiceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, InvoiceResponse{
			Code:    400,
			Message: "请求参数错误：" + err.Error(),
		})
	}

	userId, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, InvoiceResponse{
			Code:    401,
			Message: "登录过期, 请重新登录",
		})
		return
	}

	UserIdInt, ok := typeUtils.AnyToInt64(userId)
	if !ok {
		c.JSON(http.StatusInternalServerError, InvoiceResponse{
			Code:    500,
			Message: "user_id类型转换错误",
		})
		return
	}

	tenantId, exists := c.Get("tenant_id")
	var tenantIdInt int64
	if !exists {
		tenantId = int64(0)
	} else {
		tmpTenantIdInt, ok := tenantId.(int64)
		if ok {
			c.JSON(http.StatusInternalServerError, InvoiceResponse{
				Code:    500,
				Message: "tenant_id类型转换错误",
			})
			return
		}
		tenantIdInt = tmpTenantIdInt
	}
	fmt.Println(tenantIdInt)
	var invoiceDate time.Time
	if req.InvoiceDate != nil {
		invoiceDate = *req.InvoiceDate
	}

	amount := 0.0
	if req.Amount != nil {
		amount = *req.Amount
	}

	invoiceNumber := ""
	if req.InvoiceNumber != nil {
		invoiceNumber = *req.InvoiceNumber
	}

	vendor := ""
	if req.Vendor != nil {
		vendor = *req.Vendor
	}

	category := ""
	if req.Category != nil {
		category = *req.Category
	}

	var taxId = ""
	if req.TaxID != nil {
		taxId = *req.TaxID
	}

	var paymentSource = ""
	if req.PaymentSource != nil {
		paymentSource = *req.PaymentSource
	}

	var status = ""
	if req.Status != "" {
		status = req.Status
	}

	invoice, err := h.invoiceService.CreateInvoice(tenantIdInt, UserIdInt, invoiceDate, amount, invoiceNumber, vendor, taxId, category, paymentSource, status)
	if err != nil {
		c.JSON(http.StatusConflict, InvoiceResponse{
			Code:    409,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, InvoiceResponse{
		Code:    200,
		Message: "发票创建成功",
		Data: map[string]interface{}{
			"invoice": invoice,
		},
	})
}

// GetAllInvoice 分页获取所有发票接口
func (h InvoiceHandler) GetAllInvoicePage(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("page_size", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		pageSize = 10
	}
	invoices, err := h.invoiceService.GetAllInvoicePage(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, InvoiceResponse{
			Code:    500,
			Message: "获取发票列表失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, InvoiceResponse{
		Code:    200,
		Message: "获取发票列表成功",
		Data: map[string]interface{}{
			"tenants": invoices,
			"page":    page,
			"size":    len(invoices),
		},
	})
}

func (h InvoiceHandler) GetAllInvoicePageByUserId(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("page_size", "10")

	userId, exists := c.Get("user_id")
	if (!exists) || userId.(*int64) == nil {
		c.JSON(http.StatusUnauthorized, InvoiceResponse{
			Code:    401,
			Message: "登录过期, 请重新登录",
		})
		return
	}

	UserIdInt, ok := typeUtils.AnyToInt64(userId)
	if !ok {
		c.JSON(http.StatusInternalServerError, InvoiceResponse{
			Code:    500,
			Message: "user_id类型转换错误",
		})
		return
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		pageSize = 10
	}
	invoices, err := h.invoiceService.GetAllInvoicePageByUserId(page, pageSize, UserIdInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, InvoiceResponse{
			Code:    500,
			Message: "获取发票列表失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, InvoiceResponse{
		Code:    200,
		Message: "获取发票列表成功",
		Data: map[string]interface{}{
			"tenants": invoices,
			"page":    page,
			"size":    len(invoices),
		},
	})
}

func (h InvoiceHandler) GetAllInvoicePageByTenantId(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("page_size", "10")

	tenantId, exists := c.Get("tenant_id")
	if !exists || tenantId.(*int64) == nil {
		c.JSON(http.StatusConflict, InvoiceResponse{
			Code:    409,
			Message: "找不到该租户",
		})
		return
	}

	TenantIdInt, ok := typeUtils.AnyToInt64(tenantId)
	if !ok {
		c.JSON(http.StatusInternalServerError, InvoiceResponse{
			Code:    500,
			Message: "tenant_id类型转换错误",
		})
		return
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		pageSize = 10
	}
	invoices, err := h.invoiceService.GetAllInvoicePageByTenantId(page, pageSize, TenantIdInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, InvoiceResponse{
			Code:    500,
			Message: "获取发票列表失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, InvoiceResponse{
		Code:    200,
		Message: "获取发票列表成功",
		Data: map[string]interface{}{
			"tenants": invoices,
			"page":    page,
			"size":    len(invoices),
		},
	})
}

func (h InvoiceHandler) UpdateInvoice(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, TenantResponse{
			Code:    400,
			Message: "发票ID格式错误",
		})
		return
	}

	var req CreateInvoiceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, InvoiceResponse{
			Code:    400,
			Message: "请求参数错误：" + err.Error(),
		})
		return
	}

	var taxId = ""
	if req.TaxID != nil {
		taxId = *req.TaxID
	}

	var paymentSource = ""
	if req.PaymentSource != nil {
		paymentSource = *req.PaymentSource
	}

	var status = ""
	if req.Status != "" {
		status = req.Status
	}

	invoice, err := h.invoiceService.UpdateInvoice(id, *req.InvoiceDate, *req.Amount, *req.InvoiceNumber, *req.Vendor, taxId, *req.PaymentSource, paymentSource, status)
	if err != nil {
		c.JSON(http.StatusConflict, InvoiceResponse{
			Code:    409,
			Message: "更新时错误：" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, InvoiceResponse{
		Code:    200,
		Message: "发票更新成功",
		Data: map[string]interface{}{
			"invoice": invoice,
		},
	})
}
