package handler

import (
	"AstraFlow-go/internal/client"
	"AstraFlow-go/internal/model"
	"AstraFlow-go/internal/service"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Response 定义通用响应体
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type AttachmentHandler struct {
	service        service.AttachmentService
	invoiceService *service.InvoiceService
}

func NewAttachmentHandler() *AttachmentHandler {
	return &AttachmentHandler{
		service:        service.NewAttachmentService(),
		invoiceService: service.NewInvoiceService(),
	}
}

// UploadFile 上传发票文件并识别发票信息
func (h *AttachmentHandler) UploadFile(c *gin.Context) {
	// 获取当前用户信息
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, Response{
			Code:    401,
			Message: "未授权",
		})
		return
	}

	tenantID, exists := c.Get("tenant_id")
	// Handle nil tenantID
	var tenantIDPtr *int64
	if exists && tenantID != nil {
		if tid, ok := tenantID.(*int64); ok {
			tenantIDPtr = tid
		}
	}

	// 获取上传的文件
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    400,
			Message: "文件上传失败: " + err.Error(),
		})
		return
	}

	// 校验文件格式是否正确
	if err = h.service.ValidateFile(file); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    400,
			Message: "文件格式错误: " + err.Error(),
		})
		return
	}

	// 调用服务层上传文件
	attachment, err := h.service.UploadFile(file, userID.(int64), tenantIDPtr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	// 创建回调URL
	callbackURL := fmt.Sprintf("%s://%s/api/callback/ocr-result", c.Request.URL.Scheme, c.Request.Host)
	if c.Request.URL.Scheme == "" {
		proto := c.GetHeader("X-Forwarded-Proto")
		if proto != "" {
			callbackURL = fmt.Sprintf("%s://%s/api/callback/ocr-result", proto, c.Request.Host)
		} else {
			callbackURL = fmt.Sprintf("http://%s/api/callback/ocr-result", c.Request.Host)
		}
	}

	// 发布 OCR 任务
	rabbitmqClient, err := client.NewRabbitMQOCRClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Code:    500,
			Message: "RabbitMQ 客户端初始化失败: " + err.Error(),
		})
		return
	}
	defer rabbitmqClient.Close()

	_, err = rabbitmqClient.AddTask(attachment.ID, attachment.FileURL, &callbackURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Code:    500,
			Message: "发送 OCR 任务失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "文件上传成功，OCR 任务已发送到队列等待处理",
		Data:    attachment,
	})
}

// GetAttachmentByID 根据ID获取附件信息
func (h *AttachmentHandler) GetAttachmentByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    400,
			Message: "ID格式错误",
		})
		return
	}

	attachment, err := h.service.GetAttachmentByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Code:    500,
			Message: "获取附件信息失败",
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "获取附件信息成功",
		Data:    attachment,
	})
}

// GetAttachmentsByUserID 根据用户ID获取附件列表
func (h *AttachmentHandler) GetAttachmentsByUserID(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, Response{
			Code:    401,
			Message: "未授权",
		})
		return
	}

	attachments, err := h.service.GetAttachmentsByUserID(userID.(int64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Code:    500,
			Message: "获取附件列表失败",
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "获取附件列表成功",
		Data:    attachments,
	})
}

// GetAttachmentsByTenantID 根据租户ID获取附件列表
func (h *AttachmentHandler) GetAttachmentsByTenantID(c *gin.Context) {
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

	tenantID, exists := c.Get("tenant_id")
	if !exists || tenantID == nil {
		c.JSON(http.StatusUnauthorized, Response{
			Code:    401,
			Message: "未授权",
		})
		return
	}
	
	tenantIDPtr, ok := tenantID.(*int64)
	if !ok || tenantIDPtr == nil {
		c.JSON(http.StatusUnauthorized, Response{
			Code:    401,
			Message: "未授权或租户ID无效",
		})
		return
	}

	attachments, total, err := h.service.GetAttachmentsByTenantID(*tenantIDPtr, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Code:    500,
			Message: "获取附件列表失败",
		})
		return
	}

	totalPages := int(total) / pageSize
	if int(total)%pageSize > 0 {
		totalPages++
	}

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "获取附件列表成功",
		Data: map[string]interface{}{
			"attachments": attachments,
			"page":        page,
			"size":        pageSize,
			"total":       total,
			"totalPages":  totalPages,
		},
	})
}

// GetAttachmentsByInvoiceID 根据发票ID获取附件列表
func (h *AttachmentHandler) GetAttachmentsByInvoiceID(c *gin.Context) {
	invoiceID, err := strconv.ParseInt(c.Param("invoice_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    400,
			Message: "Invoice ID格式错误",
		})
		return
	}

	attachments, err := h.service.GetAttachmentsByInvoiceID(invoiceID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Code:    500,
			Message: "获取附件列表失败",
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "获取附件列表成功",
		Data:    attachments,
	})
}

// DeleteAttachment 删除附件
func (h *AttachmentHandler) DeleteAttachment(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    400,
			Message: "ID格式错误",
		})
		return
	}

	err = h.service.DeleteAttachment(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Code:    500,
			Message: "删除附件失败",
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "附件删除成功",
	})
}

// GetAttachmentsByStatus 根据状态获取附件列表
func (h *AttachmentHandler) GetAttachmentsByStatus(c *gin.Context) {
	status := c.Query("status")
	if status == "" {
		c.JSON(http.StatusBadRequest, Response{
			Code:    400,
			Message: "状态参数不能为空",
		})
		return
	}

	// 验证状态参数是否有效
	var attachmentStatus model.AttachmentStatus
	switch status {
	case "pending":
		attachmentStatus = model.AttachmentStatusPending
	case "success":
		attachmentStatus = model.AttachmentStatusSuccess
	case "failed":
		attachmentStatus = model.AttachmentStatusFailed
	default:
		c.JSON(http.StatusBadRequest, Response{
			Code:    400,
			Message: "无效的附件状态值",
		})
		return
	}

	attachments, err := h.service.GetAttachmentsByStatus(attachmentStatus)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Code:    500,
			Message: "获取附件列表失败",
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "获取附件列表成功",
		Data:    attachments,
	})
}

// HandleOCRResultCallback 处理OCR结果的回调
func (h *AttachmentHandler) HandleOCRResultCallback(c *gin.Context) {
	// 从请求体获取OCR结果
	var result client.OCRQueueResult
	if err := c.ShouldBindJSON(&result); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    400,
			Message: "解析回调数据失败: " + err.Error(),
		})
		return
	}

	log.Printf("接收到OCR结果回调，任务ID: %s，状态: %s", result.TaskID, result.Status)

	if result.Status == "success" {
		// 从result.Data中获取文件ID和其他OCR识别的数据
		fileID, ok := result.Data["file_id"].(float64) // JSON解析会将整数转换为float64
		if !ok {
			c.JSON(http.StatusBadRequest, Response{
				Code:    400,
				Message: "回调数据中缺少文件ID",
			})
			return
		}

		// 获取附件信息
		attachment, err := h.service.GetAttachmentByID(int64(fileID))
		if err != nil {
			log.Printf("获取附件信息失败，文件ID: %d, 错误: %v", int64(fileID), err)
			c.JSON(http.StatusInternalServerError, Response{
				Code:    500,
				Message: "获取附件信息失败",
			})
			return
		}

		// 从OCR结果中提取发票信息
		invoiceNumber, _ := result.Data["invoice_number"].(string)
		vendor, _ := result.Data["vendor"].(string)
		description, _ := result.Data["description"].(string)
		amountFloat, _ := result.Data["amount"].(float64)
		amount := amountFloat
		category, _ := result.Data["category"].(string)
		// Ignored: paymentSource, taxId

		// 处理日期
		var invoiceDate time.Time
		if dateStr, ok := result.Data["invoice_date"].(string); ok && dateStr != "" {
			if parsedDate, err := time.Parse("2006-01-02", dateStr); err == nil {
				invoiceDate = parsedDate
			} else {
				invoiceDate = time.Now()
			}
		} else {
			invoiceDate = time.Now()
		}
		
		var tenantId int64
		if attachment.TenantID != nil {
			tenantId = *attachment.TenantID
		}

		// 创建发票记录
		invoice, err := h.invoiceService.CreateInvoice(
			tenantId, 
			attachment.UserID, 
			attachment.ID, // Pass AttachmentID
			invoiceDate, 
			amount, 
			invoiceNumber, 
			vendor, 
			category, 
			description,
		)
		if err != nil {
			log.Printf("创建发票失败，文件ID: %d, 错误: %v", int64(fileID), err)
			c.JSON(http.StatusInternalServerError, Response{
				Code:    500,
				Message: "创建发票失败: " + err.Error(),
			})
			return
		}

		// 更新附件记录的InvoiceID字段，关联到新创建的发票
		err = h.service.UpdateAttachmentInvoiceID(int64(fileID), invoice.ID)
		if err != nil {
			log.Printf("更新附件InvoiceID失败，文件ID: %d, 发票ID: %d, 错误: %v",
				int64(fileID), invoice.ID, err)
		}
		
		// Also update status
		h.service.UpdateAttachmentStatus(int64(fileID), model.AttachmentStatusSuccess)

		log.Printf("OCR任务处理成功，发票ID: %d，文件ID: %d", invoice.ID, int64(fileID))

		c.JSON(http.StatusOK, Response{
			Code:    200,
			Message: "OCR结果处理成功",
		})
	} else {
		// OCR处理失败
		fileID, ok := result.Data["file_id"].(float64)
		if ok {
			h.service.UpdateAttachmentStatus(int64(fileID), model.AttachmentStatusFailed)
			log.Printf("OCR任务处理失败，文件ID: %d, 错误: %s", int64(fileID), result.ErrorMsg)
		}
		c.JSON(http.StatusOK, Response{
			Code:    200,
			Message: "OCR结果处理失败，已记录错误状态",
		})
	}
}