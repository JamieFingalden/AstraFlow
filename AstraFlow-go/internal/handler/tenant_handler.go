package handler

import (
	"AstraFlow-go/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// TenantHandler 租户处理器
// 处理租户相关的HTTP请求
type TenantHandler struct {
	tenantService *service.TenantService // 租户服务实例
}

// NewTenantHandler 创建租户处理器实例
// 初始化租户处理器并创建租户服务依赖
func NewTenantHandler() *TenantHandler {
	return &TenantHandler{
		tenantService: service.NewTenantService(),
	}
}

// CreateTenantRequest 创建租户请求体
type CreateTenantRequest struct {
	Name         string  `json:"name" binding:"required,min=1,max=100"`
	Industry     *string `json:"industry,omitempty"`
	ContactName  *string `json:"contact_name,omitempty"`
	ContactPhone *string `json:"contact_phone,omitempty"`
	ContactEmail *string `json:"contact_email,omitempty"`
}

// UpdateTenantRequest 更新租户请求体
type UpdateTenantRequest struct {
	Name         string  `json:"name" binding:"required,min=1,max=100"`
	Industry     *string `json:"industry,omitempty"`
	ContactName  *string `json:"contact_name,omitempty"`
	ContactPhone *string `json:"contact_phone,omitempty"`
	ContactEmail *string `json:"contact_email,omitempty"`
}

// TenantResponse 租户响应体
type TenantResponse struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data,omitempty"`
}

// CreateTenant 创建租户接口
// 处理创建租户请求，验证参数，调用服务层创建租户
func (h *TenantHandler) CreateTenant(c *gin.Context) {
	var req CreateTenantRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, TenantResponse{
			Code:    400,
			Message: "请求参数错误: " + err.Error(),
		})
		return
	}

	// 设置默认值
	industry := ""
	if req.Industry != nil {
		industry = *req.Industry
	}
	contactName := ""
	if req.ContactName != nil {
		contactName = *req.ContactName
	}
	contactPhone := ""
	if req.ContactPhone != nil {
		contactPhone = *req.ContactPhone
	}
	contactEmail := ""
	if req.ContactEmail != nil {
		contactEmail = *req.ContactEmail
	}

	// 调用服务层创建租户
	tenant, err := h.tenantService.CreateTenant(req.Name, industry, contactName, contactPhone, contactEmail)
	if err != nil {
		c.JSON(http.StatusConflict, TenantResponse{
			Code:    409,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, TenantResponse{
		Code:    200,
		Message: "创建租户成功",
		Data: map[string]interface{}{
			"tenant": tenant,
		},
	})
}

// GetTenant 获取租户详情接口
// 根据租户ID查询并返回租户详情
func (h *TenantHandler) GetTenant(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, TenantResponse{
			Code:    400,
			Message: "租户ID格式错误",
		})
		return
	}

	// 调用服务层获取租户
	tenant, err := h.tenantService.GetTenantByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, TenantResponse{
			Code:    500,
			Message: "获取租户失败: " + err.Error(),
		})
		return
	}

	if tenant == nil {
		c.JSON(http.StatusNotFound, TenantResponse{
			Code:    404,
			Message: "租户不存在",
		})
		return
	}

	c.JSON(http.StatusOK, TenantResponse{
		Code:    200,
		Message: "获取租户成功",
		Data: map[string]interface{}{
			"tenant": tenant,
		},
	})
}

// GetTenantList 获取租户列表接口
// 查询并返回租户列表，支持分页
func (h *TenantHandler) GetTenantList(c *gin.Context) {
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

	// 调用服务层获取租户列表
	tenants, err := h.tenantService.GetTenantList(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, TenantResponse{
			Code:    500,
			Message: "获取租户列表失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, TenantResponse{
		Code:    200,
		Message: "获取租户列表成功",
		Data: map[string]interface{}{
			"tenants": tenants,
			"page":    page,
			"size":    len(tenants),
		},
	})
}

// UpdateTenant 更新租户接口
// 根据租户ID更新租户信息
func (h *TenantHandler) UpdateTenant(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, TenantResponse{
			Code:    400,
			Message: "租户ID格式错误",
		})
		return
	}

	var req UpdateTenantRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, TenantResponse{
			Code:    400,
			Message: "请求参数错误: " + err.Error(),
		})
		return
	}

	// 设置默认值
	industry := ""
	if req.Industry != nil {
		industry = *req.Industry
	}
	contactName := ""
	if req.ContactName != nil {
		contactName = *req.ContactName
	}
	contactPhone := ""
	if req.ContactPhone != nil {
		contactPhone = *req.ContactPhone
	}
	contactEmail := ""
	if req.ContactEmail != nil {
		contactEmail = *req.ContactEmail
	}

	// 调用服务层更新租户
	tenant, err := h.tenantService.UpdateTenant(id, req.Name, industry, contactName, contactPhone, contactEmail)
	if err != nil {
		c.JSON(http.StatusInternalServerError, TenantResponse{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, TenantResponse{
		Code:    200,
		Message: "更新租户成功",
		Data: map[string]interface{}{
			"tenant": tenant,
		},
	})
}

// DeleteTenant 删除租户接口
// 根据租户ID删除租户
func (h *TenantHandler) DeleteTenant(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, TenantResponse{
			Code:    400,
			Message: "租户ID格式错误",
		})
		return
	}

	err = h.tenantService.DeleteTenant(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, TenantResponse{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, TenantResponse{
		Code:    200,
		Message: "删除租户成功",
	})
}
