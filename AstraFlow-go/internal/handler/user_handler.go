package handler

import (
	"AstraFlow-go/internal/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type UserHandler struct {
	userService *service.UserService // 用户服务实例
}

// NewUserHandler 创建用户处理器实例
// 初始化用户处理器
func NewUserHandler() *UserHandler {
	return &UserHandler{
		userService: service.NewUserService(),
	}
}

// UserRequest 用户请求结构体
type UserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Role     bool   `json:"role"`
}

// UserResponse 用户响应结构体
type UserResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// 获取用户列表
func (h *UserHandler) GetUserList(c *gin.Context) {
	tenantId, exists := c.Get("tenant_id")
	fmt.Println(tenantId, exists)
	if !exists {
		c.JSON(http.StatusUnauthorized, UserResponse{
			Code:    401,
			Message: "无权限",
		})
		return
	}

	tenantIdInt64 := cast.ToInt64(tenantId)

	userList, err := h.userService.GetUserList(tenantIdInt64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, UserResponse{
			Code:    500,
			Message: "获取用户列表失败：" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, UserResponse{
		Code:    200,
		Message: "获取用户列表成功",
		Data:    userList,
	})
}

// 添加用户
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, UserResponse{
			Code:    400,
			Message: "请求参数错误：" + err.Error(),
		})
		return
	}

	tenantId, exists := c.Get("tenant_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, UserResponse{
			Code:    401,
			Message: "无权限",
		})
		return
	}

	userId, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, UserResponse{
			Code:    401,
			Message: "无权限",
		})
		return
	}

	userIdInt64 := cast.ToInt64(userId)
	role, err := h.userService.FindUserRoleByUserId(userIdInt64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, UserResponse{
			Code:    500,
			Message: "获取用户角色失败：" + err.Error(),
		})
		return
	}
	if role != "admin" {
		c.JSON(http.StatusUnauthorized, UserResponse{
			Code:    401,
			Message: "无权限",
		})
		return
	}

	tenantIdInt64 := cast.ToInt64(tenantId)
	role = "normal"
	if req.Role {
		role = "admin"
	}

	if req.Password == "" {
		req.Password = "123456"
	}

	user, err := h.userService.CreateUser(tenantIdInt64, req.Username, req.Password, req.Email, req.Phone, role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, UserResponse{
			Code:    500,
			Message: "创建用户失败：" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, UserResponse{
		Code:    200,
		Message: "创建用户成功",
		Data:    user,
	})
}

// 更新用户
func (h *UserHandler) UpdateUser(c *gin.Context) {
	var req UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, UserResponse{
			Code:    400,
			Message: "请求参数错误：" + err.Error(),
		})
		return
	}

	userId, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, UserResponse{
			Code:    401,
			Message: "无权限",
		})
		return
	}

	userIdInt64 := cast.ToInt64(userId)
	role, err := h.userService.FindUserRoleByUserId(userIdInt64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, UserResponse{
			Code:    500,
			Message: "获取用户角色失败：" + err.Error(),
		})
		return
	}
	if role != "admin" {
		c.JSON(http.StatusUnauthorized, UserResponse{
			Code:    401,
			Message: "无权限",
		})
		return
	}

	role = "normal"
	if req.Role {
		role = "admin"
	}

	idStr := c.Param("id")
	idInt64 := cast.ToInt64(idStr)

	err = h.userService.UpdateUser(idInt64, req.Username, req.Password, req.Email, req.Phone, role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, UserResponse{
			Code:    500,
			Message: "更新用户失败：" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, UserResponse{
		Code:    200,
		Message: "更新用户成功",
	})
}

// 删除用户
func (h *UserHandler) DeleteUser(c *gin.Context) {
	userId, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, UserResponse{
			Code:    401,
			Message: "无权限",
		})
		return
	}

	userIdInt64 := cast.ToInt64(userId)
	role, err := h.userService.FindUserRoleByUserId(userIdInt64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, UserResponse{
			Code:    500,
			Message: "获取用户角色失败：" + err.Error(),
		})
		return
	}
	if role != "admin" {
		c.JSON(http.StatusUnauthorized, UserResponse{
			Code:    401,
			Message: "无权限",
		})
		return
	}

	idStr := c.Param("id")
	idInt64 := cast.ToInt64(idStr)

	err = h.userService.DeleteUser(idInt64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, UserResponse{
			Code:    500,
			Message: "删除用户失败：" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, UserResponse{
		Code:    200,
		Message: "删除用户成功",
	})
}

// 获取用户详情
func (h *UserHandler) GetUserById(c *gin.Context) {
	// Get the authenticated user ID (current user)
	authUserId, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, UserResponse{
			Code:    401,
			Message: "无权限",
		})
		return
	}

	authUserIdInt64 := cast.ToInt64(authUserId)

	// Get the requested user ID from the URL parameter
	requestedUserIdStr := c.Param("id")
	requestedUserIdInt64 := cast.ToInt64(requestedUserIdStr)

	// Check if the authenticated user has permission to access the requested user
	authUserRole, err := h.userService.FindUserRoleByUserId(authUserIdInt64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, UserResponse{
			Code:    500,
			Message: "获取当前用户角色失败：" + err.Error(),
		})
		return
	}

	// Allow access if it's the same user or if current user is admin
	if requestedUserIdInt64 != authUserIdInt64 && authUserRole != "admin" {
		c.JSON(http.StatusUnauthorized, UserResponse{
			Code:    401,
			Message: "无权限",
		})
		return
	}

	user, err := h.userService.FindUserById(requestedUserIdInt64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, UserResponse{
			Code:    500,
			Message: "获取用户详情失败：" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, UserResponse{
		Code:    200,
		Message: "获取用户详情成功",
		Data:    user,
	})
}


