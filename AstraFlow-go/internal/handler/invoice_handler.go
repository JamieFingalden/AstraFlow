package handler

import (
	"AstraFlow-go/internal/client"
	"AstraFlow-go/internal/service"
	typeUtils "AstraFlow-go/pkg/utils"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cast"

	"github.com/gin-gonic/gin"
)

// InvoiceHandler 发票处理器
type InvoiceHandler struct {
	invoiceService    *service.InvoiceService   // 发票服务实例
	attachmentService service.AttachmentService // 附件服务实例
}

// NewInvoiceHandler 创建发票处理器实例
func NewInvoiceHandler() *InvoiceHandler {
	return &InvoiceHandler{
		invoiceService:    service.NewInvoiceService(),
		attachmentService: service.NewAttachmentService(),
	}
}

// CreateInvoiceRequest 创建发票请求体
type CreateInvoiceRequest struct {
	TenantID      *int64     `json:"tenant_id,omitempty" form:"tenant_id"`
	UserId        int64      `json:"user_id" form:"user_id"`
	AttachmentID  int64      `json:"attachment_id" form:"attachment_id"` // Optional if file is provided
	InvoiceNumber *string    `json:"invoice_number" form:"invoice_number"`
	InvoiceDate   *time.Time `json:"invoice_date" form:"invoice_date" time_format:"2006-01-02T15:04:05Z07:00"`
	Amount        *float64   `json:"amount" form:"amount"`
	Vendor        string     `json:"vendor" form:"vendor" binding:"required"`
	Description   *string    `json:"description,omitempty" form:"description"`
	Category      *string    `json:"category" form:"category" binding:"required"`
	Status        string     `json:"status" form:"status"`
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
	// Support both JSON and Form data
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, InvoiceResponse{
			Code:    400,
			Message: "请求参数错误：" + err.Error(),
		})
		return
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
	var tenantIdPtr *int64
	if !exists || tenantId == nil {
		tenantIdInt = 0
	} else {
		tenantIdInt = cast.ToInt64(tenantId)
		tenantIdPtr = &tenantIdInt
	}

	// Handle file upload if present
	file, err := c.FormFile("file")
	var attachmentID = req.AttachmentID
	if err == nil && file != nil {
		// Upload file without OCR
		attachment, err := h.attachmentService.UploadFile(file, UserIdInt, tenantIdPtr)
		if err != nil {
			c.JSON(http.StatusInternalServerError, InvoiceResponse{
				Code:    500,
				Message: "文件上传失败: " + err.Error(),
			})
			return
		}
		attachmentID = attachment.ID
		// Update status to success since we are manually filling it or just storing it
		h.attachmentService.UpdateAttachmentStatus(attachment.ID, 1) // 1 corresponds to AttachmentStatusSuccess
	}

	// If no file and no attachmentID, we might still allow it but usually it needs one or the other
	// for manual entry without file, attachmentID might be 0.

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

	category := ""
	if req.Category != nil {
		category = *req.Category
	}

	var description = ""
	if req.Description != nil {
		description = *req.Description
	}

	invoice, err := h.invoiceService.CreateInvoice(tenantIdInt, UserIdInt, attachmentID, invoiceDate, amount, invoiceNumber, req.Vendor, category, description)
	if err != nil {
		c.JSON(http.StatusConflict, InvoiceResponse{
			Code:    409,
			Message: err.Error(),
		})
		return
	}

	// If we created a new attachment, link it to the invoice
	if attachmentID != req.AttachmentID && attachmentID != 0 {
		h.attachmentService.UpdateAttachmentInvoiceID(attachmentID, invoice.ID)
	}

	c.JSON(http.StatusOK, InvoiceResponse{
		Code:    200,
		Message: "发票创建成功",
		Data: map[string]interface{}{
			"invoice": invoice,
		},
	})
}

// UploadOCR [POST] /api/v1/invoices/upload-ocr
// AI 智能流：上传图片并发布 OCR 任务到队列
func (h InvoiceHandler) UploadOCR(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, InvoiceResponse{
			Code:    400,
			Message: "请上传发票图片 (file 字段缺失)",
		})
		return
	}

	userId, _ := c.Get("user_id")
	userIdInt := cast.ToInt64(userId)

	tenantId, exists := c.Get("tenant_id")
	var tenantIdPtr *int64
	if exists && tenantId != nil {
		id := cast.ToInt64(tenantId)
		tenantIdPtr = &id
	}

	// 1. 保存图片到附件表
	attachment, err := h.attachmentService.UploadFile(file, userIdInt, tenantIdPtr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, InvoiceResponse{
			Code:    500,
			Message: "文件保存失败: " + err.Error(),
		})
		return
	}

	// 2. 创建发票记录 (状态为 recognizing, 来源为 ocr)
	invoice, err := h.invoiceService.CreateOCRInvoice(tenantIdPtr, userIdInt, attachment.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, InvoiceResponse{
			Code:    500,
			Message: "单据初始化失败: " + err.Error(),
		})
		return
	}

	// 3. 发布 OCR 任务到 RabbitMQ 队列
	// 后续由 Python AI Worker 消费并回调更新状态
	if err := client.PublishOCRTask(attachment.ID, invoice.ID, attachment.FileURL); err != nil {
		// 注意：此处可以根据业务决定是否需要回滚操作
		// 暂时只记录错误，因为单据记录已创建，可由后台任务补偿
		log.Printf("发布 OCR 任务到队列失败: %v", err)
		c.JSON(http.StatusInternalServerError, InvoiceResponse{
			Code:    500,
			Message: "发布 AI 识别任务失败，请稍后重试",
		})
		return
	}

	c.JSON(http.StatusCreated, InvoiceResponse{
		Code:    201,
		Message: "上传成功，已加入AI识别队列",
		Data: map[string]interface{}{
			"invoice_id": invoice.ID,
			"status":     invoice.Status,
		},
	})
}

// GetMyInvoices [GET] /api/v1/invoices/my-invoices
// 获取当前用户自己的发票列表（可根据状态筛选）
func (h InvoiceHandler) GetMyInvoices(c *gin.Context) {
	// 从认证中间件获取用户ID
	userId, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, InvoiceResponse{
			Code:    401,
			Message: "无法获取用户信息，请重新登录",
		})
		return
	}
	userIdInt := cast.ToInt64(userId)

	// 获取查询参数
	status := c.Query("status") // 'recognizing', 'unconfirmed', 'draft'
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	invoices, total, err := h.invoiceService.GetInvoicesByUserIDAndStatus(page, pageSize, userIdInt, status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, InvoiceResponse{
			Code:    500,
			Message: "获取发票列表失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, InvoiceResponse{
		Code:    200,
		Message: "获取成功",
		Data: gin.H{
			"items":      invoices,
			"total":      total,
			"page":       page,
			"pageSize":   pageSize,
			"totalPages": (total + int64(pageSize) - 1) / int64(pageSize),
		},
	})
}

// UploadManual [POST] /api/v1/invoices/upload-manual
// 手动兜底流：上传图片并手动填报单据
func (h InvoiceHandler) UploadManual(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, InvoiceResponse{
			Code:    400,
			Message: "请上传发票图片",
		})
		return
	}

	// 解析表单数据
	amount := cast.ToFloat64(c.PostForm("amount"))
	invoiceNumber := strings.TrimSpace(c.PostForm("invoice_number"))
	vendor := strings.TrimSpace(c.PostForm("vendor"))
	paymentMethod := strings.TrimSpace(c.PostForm("payment_method"))
	category := strings.TrimSpace(c.PostForm("category"))
	description := strings.TrimSpace(c.PostForm("description"))
	dateStr := c.PostForm("invoice_date")
	invoiceDate, _ := time.Parse("2006-01-02", dateStr)

	if invoiceNumber == "" || amount <= 0 || vendor == "" || dateStr == "" || paymentMethod == "" || category == "" {
		c.JSON(http.StatusBadRequest, InvoiceResponse{
			Code:    400,
			Message: "请完整填写：发票号码、金额、商户名称、日期、支付方式、消费类别",
		})
		return
	}

	if invoiceDate.IsZero() {
		c.JSON(http.StatusBadRequest, InvoiceResponse{
			Code:    400,
			Message: "发票日期格式错误，应为 YYYY-MM-DD",
		})
		return
	}

	userId, _ := c.Get("user_id")
	userIdInt := cast.ToInt64(userId)

	tenantId, exists := c.Get("tenant_id")
	var tenantIdPtr *int64
	if exists && tenantId != nil {
		id := cast.ToInt64(tenantId)
		tenantIdPtr = &id
	}

	// 1. 保存图片到附件表
	attachment, err := h.attachmentService.UploadFile(file, userIdInt, tenantIdPtr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, InvoiceResponse{
			Code:    500,
			Message: "文件上传失败",
		})
		return
	}

	// 2. 创建发票记录 (状态直接为 draft, 来源为 manual)
	invoice, err := h.invoiceService.CreateManualInvoice(
		tenantIdPtr,
		userIdInt,
		attachment.ID,
		amount,
		invoiceDate,
		invoiceNumber,
		vendor,
		paymentMethod,
		category,
		description,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, InvoiceResponse{
			Code:    500,
			Message: "保存单据失败",
		})
		return
	}

	c.JSON(http.StatusCreated, InvoiceResponse{
		Code:    201,
		Message: "单据已存入草稿箱",
		Data: map[string]interface{}{
			"invoice_id": invoice.ID,
			"status":     invoice.Status,
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
	invoices, total, err := h.invoiceService.GetAllInvoicePage(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, InvoiceResponse{
			Code:    500,
			Message: "获取发票列表失败: " + err.Error(),
		})
		return
	}

	totalPages := int(total) / pageSize
	if int(total)%pageSize > 0 {
		totalPages++
	}

	c.JSON(http.StatusOK, InvoiceResponse{
		Code:    200,
		Message: "获取发票列表成功",
		Data: map[string]interface{}{
			"invoices":   invoices,
			"page":       page,
			"size":       pageSize,
			"total":      total,
			"totalPages": totalPages,
		},
	})
}

func (h InvoiceHandler) GetAllInvoicePageByUserId(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("page_size", "10")

	userId, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, InvoiceResponse{
			Code:    401,
			Message: "登录过期, 请重新登录",
		})
		return
	}

	UserIdInt := cast.ToInt64(userId)

	tenantId, exists := c.Get("tenant_id")
	var tenantIdInt int64
	if !exists {
		tenantIdInt = int64(0)
	}

	tenantIdInt = cast.ToInt64(tenantId)

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		pageSize = 10
	}

	invoices, total, err := h.invoiceService.GetAllInvoicePageByUserId(page, pageSize, UserIdInt, tenantIdInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, InvoiceResponse{
			Code:    500,
			Message: "获取发票列表失败: " + err.Error(),
		})
		return
	}

	totalPages := int(total) / pageSize
	if int(total)%pageSize > 0 {
		totalPages++
	}

	c.JSON(http.StatusOK, InvoiceResponse{
		Code:    200,
		Message: "获取发票列表成功",
		Data: map[string]interface{}{
			"invoices":   invoices,
			"page":       page,
			"size":       pageSize,
			"total":      total,
			"totalPages": totalPages,
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
			Message: "无法获取租户id",
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
	invoices, total, err := h.invoiceService.GetAllInvoicePageByTenantId(page, pageSize, TenantIdInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, InvoiceResponse{
			Code:    500,
			Message: "获取发票列表失败: " + err.Error(),
		})
		return
	}

	totalPages := int(total) / pageSize
	if int(total)%pageSize > 0 {
		totalPages++
	}

	c.JSON(http.StatusOK, InvoiceResponse{
		Code:    200,
		Message: "获取发票列表成功",
		Data: map[string]interface{}{
			"invoices":   invoices,
			"page":       page,
			"size":       pageSize,
			"total":      total,
			"totalPages": totalPages,
		},
	})
}

func (h InvoiceHandler) UpdateInvoice(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, InvoiceResponse{
			Code:    400,
			Message: "发票ID格式错误",
		})
		return
	}

	var req CreateInvoiceRequest
	if err = c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, InvoiceResponse{
			Code:    400,
			Message: "请求参数错误：" + err.Error(),
		})
		return
	}

	var status = ""
	if req.Status != "" {
		status = req.Status
	}

	var description = ""
	if req.Description != nil {
		description = *req.Description
	}

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

	category := ""
	if req.Category != nil {
		category = *req.Category
	}

	invoice, err := h.invoiceService.UpdateInvoice(id, invoiceDate, amount, invoiceNumber, req.Vendor, category, description, status)
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

func (h InvoiceHandler) DeleteInvoice(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, InvoiceResponse{
			Code:    400,
			Message: "发票ID格式错误",
		})
		return
	}

	userId, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, InvoiceResponse{
			Code:    401,
			Message: "登录过期, 请重新登录",
		})
		return
	}

	userIdInt := cast.ToInt64(userId)
	err = h.invoiceService.DeleteInvoice(id, userIdInt)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrInvalidInvoiceID), errors.Is(err, service.ErrInvoiceDeleteNotAllowed):
			c.JSON(http.StatusBadRequest, InvoiceResponse{
				Code:    400,
				Message: "当前状态不可删除，单据已进入流转流程",
			})
			return
		case errors.Is(err, service.ErrInvoiceForbidden):
			c.JSON(http.StatusForbidden, InvoiceResponse{
				Code:    403,
				Message: "您没有权限删除该发票",
			})
			return
		case errors.Is(err, service.ErrInvoiceNotFound):
			c.JSON(http.StatusNotFound, InvoiceResponse{
				Code:    404,
				Message: "发票不存在",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, InvoiceResponse{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, InvoiceResponse{
		Code:    200,
		Message: "发票删除成功",
	})
}

// GetInvoiceDetail 获取当前用户可访问的单据详情
func (h InvoiceHandler) GetInvoiceDetail(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, InvoiceResponse{
			Code:    400,
			Message: "发票ID格式错误",
		})
		return
	}

	userId, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, InvoiceResponse{
			Code:    401,
			Message: "登录过期, 请重新登录",
		})
		return
	}
	userIdInt := cast.ToInt64(userId)

	detail, err := h.invoiceService.GetInvoiceDetailForUser(id, userIdInt)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrInvalidInvoiceID):
			c.JSON(http.StatusBadRequest, InvoiceResponse{
				Code:    400,
				Message: "发票ID格式错误",
			})
			return
		case errors.Is(err, service.ErrInvoiceNotFound):
			c.JSON(http.StatusNotFound, InvoiceResponse{
				Code:    404,
				Message: "发票不存在",
			})
			return
		case errors.Is(err, service.ErrInvoiceForbidden):
			c.JSON(http.StatusForbidden, InvoiceResponse{
				Code:    403,
				Message: "您没有权限查看该发票",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, InvoiceResponse{
			Code:    500,
			Message: "获取发票详情失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, InvoiceResponse{
		Code:    200,
		Message: "获取发票详情成功",
		Data: map[string]interface{}{
			"invoice": detail,
		},
	})
}

// ConfirmInvoice 确认发票信息
func (h InvoiceHandler) ConfirmInvoice(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, InvoiceResponse{
			Code:    400,
			Message: "发票ID格式错误",
		})
		return
	}

	var req CreateInvoiceRequest
	if err = c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, InvoiceResponse{
			Code:    400,
			Message: "请求参数错误：" + err.Error(),
		})
		return
	}

	var description = ""
	if req.Description != nil {
		description = *req.Description
	}

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

	category := ""
	if req.Category != nil {
		category = *req.Category
	}

	invoice, err := h.invoiceService.ConfirmInvoice(id, invoiceDate, amount, invoiceNumber, req.Vendor, category, description)
	if err != nil {
		c.JSON(http.StatusConflict, InvoiceResponse{
			Code:    409,
			Message: "确认失败：" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, InvoiceResponse{
		Code:    200,
		Message: "发票已确认并移入草稿箱",
		Data: map[string]interface{}{
			"invoice": invoice,
		},
	})
}

type PublishInvoicesRequest struct {
	IDs []int64 `json:"ids" binding:"required,min=1"`
}

// BatchPayInvoicesRequest 批量打款请求体
type BatchPayInvoicesRequest struct {
	InvoiceIDs []int64 `json:"invoice_ids" binding:"required,min=1"`
}

// PublishInvoices 发布发票（待发布 -> 待审核）
func (h InvoiceHandler) PublishInvoices(c *gin.Context) {
	var req PublishInvoicesRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, InvoiceResponse{
			Code:    400,
			Message: "请求参数错误：" + err.Error(),
		})
		return
	}

	err := h.invoiceService.PublishInvoices(req.IDs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, InvoiceResponse{
			Code:    500,
			Message: "发布失败：" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, InvoiceResponse{
		Code:    200,
		Message: "发布成功，已提交审核",
	})
}

// GetApprovedInvoices 获取待打款单据列表
// 入参：Query 参数 page/size
// 出参：统一响应结构，data 包含 items/page/size/total/total_pages
func (h InvoiceHandler) GetApprovedInvoices(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		c.JSON(http.StatusBadRequest, InvoiceResponse{
			Code:    400,
			Message: "page 参数必须为大于 0 的整数",
		})
		return
	}

	size, err := strconv.Atoi(c.DefaultQuery("size", "10"))
	if err != nil || size < 1 {
		c.JSON(http.StatusBadRequest, InvoiceResponse{
			Code:    400,
			Message: "size 参数必须为大于 0 的整数",
		})
		return
	}

	result, err := h.invoiceService.GetApprovedInvoices(page, size)
	if err != nil {
		c.JSON(http.StatusInternalServerError, InvoiceResponse{
			Code:    500,
			Message: "获取待打款单据失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, InvoiceResponse{
		Code:    200,
		Message: "获取待打款单据成功",
		Data: map[string]interface{}{
			"items":       result.Items,
			"page":        result.Page,
			"size":        result.Size,
			"total":       result.Total,
			"total_pages": result.TotalPages,
		},
	})
}

// BatchPayInvoices 批量确认打款
// 入参：JSON 请求体 invoice_ids
// 出参：统一响应结构，成功返回 200，失败根据业务错误返回对应状态码
func (h InvoiceHandler) BatchPayInvoices(c *gin.Context) {
	var req BatchPayInvoicesRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, InvoiceResponse{
			Code:    400,
			Message: "请求参数错误: " + err.Error(),
		})
		return
	}

	if len(req.InvoiceIDs) == 0 {
		c.JSON(http.StatusBadRequest, InvoiceResponse{
			Code:    400,
			Message: "invoice_ids 不能为空",
		})
		return
	}

	err := h.invoiceService.BatchPayInvoices(req.InvoiceIDs)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrInvalidInvoiceIDs):
			c.JSON(http.StatusBadRequest, InvoiceResponse{
				Code:    400,
				Message: "invoice_ids 参数不合法",
			})
			return
		case errors.Is(err, service.ErrBatchPayCountMismatch):
			c.JSON(http.StatusBadRequest, InvoiceResponse{
				Code:    400,
				Message: "部分单据不存在或状态不是 approved，批量打款失败",
			})
			return
		default:
			c.JSON(http.StatusInternalServerError, InvoiceResponse{
				Code:    500,
				Message: "批量打款失败: " + err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, InvoiceResponse{
		Code:    200,
		Message: "批量打款成功",
	})
}
