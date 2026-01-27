package handler

import (
	"AstraFlow-go/internal/model"
	"AstraFlow-go/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		userService: service.NewUserService(),
	}
}

// CreateUserRequest 添加员工请求体
type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
	RoleKey  string `json:"role_key" binding:"required"` // "admin", "auditor", "employee"
}

// UpdateUserRequest 更新用户请求体
type UpdateUserRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	RoleKey string `json:"role_key"`
}

// GetUserList 获取用户列表
func (h *UserHandler) GetUserList(c *gin.Context) {
	tenantId, exists := c.Get("tenant_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "无权限"})
		return
	}

	tenantIdInt64 := cast.ToInt64(tenantId)

	userList, err := h.userService.GetUserList(tenantIdInt64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "获取用户列表失败：" + err.Error()})
		return
	}

	// Hide passwords
	for _, u := range userList {
		u.Password = ""
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取用户列表成功",
		"data":    userList,
	})
}

// CreateUser 添加员工 (Internal Invite)
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "请求参数错误：" + err.Error()})
		return
	}

	// 1. Check Permissions (Must be Admin)
	roleKey := c.GetString("role")
	if roleKey != model.RoleKeyAdmin {
		c.JSON(http.StatusForbidden, gin.H{"code": 403, "message": "只有管理员可以添加员工"})
		return
	}

	// 2. Get Tenant ID
	tenantId, exists := c.Get("tenant_id")
	if !exists || tenantId == nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "无法确定当前租户"})
		return
	}
	tenantIdInt64 := cast.ToInt64(tenantId)

	// 3. Create User via Service
	user, err := h.userService.CreateUser(
		tenantIdInt64,
		req.Username,
		req.Password,
		req.Name,
		req.RoleKey,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "创建用户失败：" + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "添加员工成功",
		"data":    user,
	})
}

// UpdateUser 更新用户
func (h *UserHandler) UpdateUser(c *gin.Context) {
	var req UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "请求参数错误：" + err.Error()})
		return
	}

	// Check Admin Perms
	roleKey := c.GetString("role")
	if roleKey != model.RoleKeyAdmin {
		c.JSON(http.StatusForbidden, gin.H{"code": 403, "message": "无权限"})
		return
	}

	idStr := c.Param("id")
	idInt64 := cast.ToInt64(idStr)

	err := h.userService.UpdateUser(idInt64, "", "", req.Email, req.Phone, req.RoleKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "更新用户失败：" + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "message": "更新用户成功"})
}

// DeleteUser 删除用户
func (h *UserHandler) DeleteUser(c *gin.Context) {
	// Check Admin Perms
	roleKey := c.GetString("role")
	if roleKey != model.RoleKeyAdmin {
		c.JSON(http.StatusForbidden, gin.H{"code": 403, "message": "无权限"})
		return
	}

	idStr := c.Param("id")
	idInt64 := cast.ToInt64(idStr)

	err := h.userService.DeleteUser(idInt64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "删除用户失败：" + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "message": "删除用户成功"})
}

// GetUserById 获取用户详情
func (h *UserHandler) GetUserById(c *gin.Context) {
	authUserId := cast.ToInt64(c.GetString("user_id")) // Wait, middleware sets int64? Check casting
	// Context Get returns interface{}. Middleware set it as int64.
	// cast.ToInt64 handles interface{} well.

	// Better check:
	val, _ := c.Get("user_id")
	authUserId = cast.ToInt64(val)

	requestedUserIdStr := c.Param("id")
	requestedUserIdInt64 := cast.ToInt64(requestedUserIdStr)

	roleKey := c.GetString("role")

	// Perms: Admin OR Self
	if requestedUserIdInt64 != authUserId && roleKey != model.RoleKeyAdmin {
		c.JSON(http.StatusForbidden, gin.H{"code": 403, "message": "无权限"})
		return
	}

	user, err := h.userService.FindUserById(requestedUserIdInt64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "获取用户详情失败：" + err.Error()})
		return
	}

	if user != nil {
		user.Password = ""
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取用户详情成功",
		"data":    user,
	})
}
