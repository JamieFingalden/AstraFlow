package handler

import (
	"AstraFlow-go/internal/client"
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
//  1. 用户上传文件
//  2. Go 保存文件到本地
//  3. Go 插入 attachment 表（记录文件信息）
//  4. Go 发送 OCR 任务到队列（Python 端消费者处理）
//  5. Python 端处理完成后，通过回调函数更新结果
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
	if !exists {
		c.JSON(http.StatusUnauthorized, Response{
			Code:    401,
			Message: "未授权",
		})
		return
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
	attachment, err := h.service.UploadFile(file, userID.(int64), tenantID.(*int64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	// 创建回调URL，以便Python端处理完成后可以通知Go端更新结果
	// 这里使用当前请求的主机信息构建回调URL
	callbackURL := fmt.Sprintf("%s://%s/api/callback/ocr-result", c.Request.URL.Scheme, c.Request.Host)
	if c.Request.URL.Scheme == "" {
		// 如果请求中没有协议，尝试从X-Forwarded-Proto头获取
		proto := c.GetHeader("X-Forwarded-Proto")
		if proto != "" {
			callbackURL = fmt.Sprintf("%s://%s/api/callback/ocr-result", proto, c.Request.Host)
		} else {
			// 默认使用http
			callbackURL = fmt.Sprintf("http://%s/api/callback/ocr-result", c.Request.Host)
		}
	}

	// 发布 OCR 任务到 RabbitMQ 队列，由 Python 端消费者处理
	rabbitmqClient, err := client.NewRabbitMQOCRClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Code:    500,
			Message: "RabbitMQ 客户端初始化失败: " + err.Error(),
		})
		return
	}
	defer rabbitmqClient.Close()

	// 发送任务到队列，包含回调URL以便Python端处理完成后通知Go端
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
	tenantID, exists := c.Get("tenant_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, Response{
			Code:    401,
			Message: "未授权",
		})
		return
	}

	attachments, err := h.service.GetAttachmentsByTenantID(*tenantID.(*int64))
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

// HandleOCRResultCallback 处理OCR结果的回调
// Python端处理完成后，会向此接口发送请求
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
		amount := amountFloat // 直接使用float64类型，不需要指针
		category, _ := result.Data["category"].(string)
		paymentSource, _ := result.Data["payment_source"].(string)
		taxId, _ := result.Data["tax_id"].(string)

		// 处理日期
		var invoiceDate time.Time
		if dateStr, ok := result.Data["invoice_date"].(string); ok && dateStr != "" {
			if parsedDate, err := time.Parse("2006-01-02", dateStr); err == nil {
				invoiceDate = parsedDate
			} else {
				// 如果日期解析失败，使用当前时间
				invoiceDate = time.Now()
			}
		} else {
			// 如果没有日期信息，使用当前时间
			invoiceDate = time.Now()
		}

		// 创建发票记录
		invoice, err := h.invoiceService.CreateInvoice(
			*attachment.TenantID, // 使用附件的租户ID
			attachment.UserID,    // 使用附件的用户ID
			invoiceDate,          // 使用从OCR结果解析的日期
			amount,               // 从OCR结果获取的金额
			invoiceNumber,        // 发票号码
			vendor,               // 商户名称
			taxId,                // 税号
			category,             // 分类
			paymentSource,        // 支付方式
			"recognized",         // 状态
			&attachment.FileURL,  // 发票图片URL
			&description,         // 描述
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
			// 这里不返回错误，因为发票已经创建成功，只是更新附件关联失败
		}

		log.Printf("OCR任务处理成功，发票ID: %d，文件ID: %d", invoice.ID, int64(fileID))

		c.JSON(http.StatusOK, Response{
			Code:    200,
			Message: "OCR结果处理成功",
		})
	} else {
		// OCR处理失败，更新附件状态为失败
		fileID, ok := result.Data["file_id"].(float64)
		if ok {
			// 尝试获取附件并记录失败信息
			attachment, err := h.service.GetAttachmentByID(int64(fileID))
			if err == nil {
				// 可以考虑在这里更新附件的状态或添加错误信息
				log.Printf("OCR任务处理失败，文件ID: %d, 文件路径: %s, 错误: %s",
					int64(fileID), attachment.FileURL, result.ErrorMsg)
			} else {
				log.Printf("OCR任务处理失败，文件ID: %d, 但无法获取附件信息: %v", int64(fileID), err)
			}
		} else {
			log.Printf("OCR任务处理失败，但回调数据中缺少文件ID，任务ID: %s, 错误: %s",
				result.TaskID, result.ErrorMsg)
		}

		log.Printf("OCR任务处理失败: %s, 错误: %s", result.TaskID, result.ErrorMsg)
		c.JSON(http.StatusOK, Response{
			Code:    200,
			Message: "OCR结果处理失败，已记录错误状态",
		})
	}
}
