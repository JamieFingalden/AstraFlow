package handler

import (
	"AstraFlow-go/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Response 定义通用响应体
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type AttachmentHandler struct {
	service service.AttachmentService
}

func NewAttachmentHandler() *AttachmentHandler {
	return &AttachmentHandler{
		service: service.NewAttachmentService(),
	}
}

// UploadFile 上传文件
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
	if err := h.service.ValidateFile(file); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    400,
			Message: "文件格式错误: " + err.Error(),
		})
		return
	}

	// TODO Use Flask to get invoice image information and save it to the database via an HTTP request

	// TODO Save the invoice information first to obtain the invoice ID
	var invoiceID *int64

	// 调用服务层上传文件
	attachment, err := h.service.UploadFile(file, userID.(int64), tenantID.(*int64), invoiceID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "文件上传成功",
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
