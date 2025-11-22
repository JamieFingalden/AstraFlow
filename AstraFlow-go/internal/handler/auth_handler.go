package handler

import (
	"AstraFlow-go/internal/service"
	"AstraFlow-go/pkg/jwt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AuthHandler 认证处理器
// 处理用户注册、登录、Token刷新等认证相关的HTTP请求
type AuthHandler struct {
	authService *service.AuthService // 认证服务实例
}

// NewAuthHandler 创建认证处理器实例
// 初始化认证处理器并创建认证服务依赖
func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		authService: service.NewAuthService(),
	}
}

// RegisterRequest 用户注册请求体
type RegisterRequest struct {
	Username string  `json:"username" binding:"required,min=3,max=50"`
	Password string  `json:"password" binding:"required,min=6,max=100"`
	Email    *string `json:"email,omitempty"`
	Phone    *string `json:"phone,omitempty"`
	TenantID *int64  `json:"tenant_id,omitempty"`
}

// RegisterResponse 用户注册响应体
type RegisterResponse struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data,omitempty"`
}

// Register 用户注册接口
// 处理用户注册请求，验证参数，调用服务层创建用户，生成JWT Token
func (h *AuthHandler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, RegisterResponse{
			Code:    400,
			Message: "请求参数错误: " + err.Error(),
		})
		return
	}

	// 设置默认值
	email := ""
	if req.Email != nil {
		email = *req.Email
	}
	phone := ""
	if req.Phone != nil {
		phone = *req.Phone
	}

	// 调用服务层注册用户
	user, err := h.authService.Register(req.Username, req.Password, email, phone, req.TenantID)
	if err != nil {
		c.JSON(http.StatusConflict, RegisterResponse{
			Code:    409,
			Message: err.Error(),
		})
		return
	}

	// 生成JWT Token
	token, err := jwt.GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, RegisterResponse{
			Code:    500,
			Message: "Token生成失败",
		})
		return
	}

	c.JSON(http.StatusOK, RegisterResponse{
		Code:    200,
		Message: "注册成功",
		Data: map[string]interface{}{
			"user":  user,
			"token": token,
		},
	})
}

// LoginRequest 用户登录请求体
type LoginRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password" binding:"required"`
}

// Login 用户登录接口
// 处理用户登录请求，验证用户名密码，生成访问Token和刷新Token
func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, RegisterResponse{
			Code:    400,
			Message: "请求参数错误: " + err.Error(),
		})
		return
	}

	// 调用服务层登录
	user, err := h.authService.Login(req.Username, req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, RegisterResponse{
			Code:    401,
			Message: err.Error(),
		})
		return
	}

	// 生成JWT Token
	token, err := jwt.GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, RegisterResponse{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	// 生成刷新Token
	refreshToken, err := jwt.GenerateRefreshToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, RegisterResponse{
			Code:    500,
			Message: "刷新Token生成失败",
		})
		return
	}

	c.JSON(http.StatusOK, RegisterResponse{
		Code:    200,
		Message: "登录成功",
		Data: map[string]interface{}{
			"user":          user,
			"token":         token,
			"refresh_token": refreshToken,
		},
	})
}

// RefreshToken 刷新Token接口
// 使用刷新Token生成新的访问Token
func (h *AuthHandler) RefreshToken(c *gin.Context) {
	refreshToken := c.GetHeader("Authorization")
	if refreshToken == "" {
		c.JSON(http.StatusBadRequest, RegisterResponse{
			Code:    400,
			Message: "缺少刷新Token",
		})
		return
	}

	// 移除 "Bearer " 前缀
	if len(refreshToken) > 7 && refreshToken[:7] == "Bearer " {
		refreshToken = refreshToken[7:]
	}

	// 解析刷新Token
	claims, err := jwt.ParseToken(refreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, RegisterResponse{
			Code:    401,
			Message: "刷新Token无效",
		})
		return
	}

	// 根据用户ID获取用户信息
	user, err := h.authService.GetUserByID(claims.UserID)
	if err != nil || user == nil {
		c.JSON(http.StatusNotFound, RegisterResponse{
			Code:    404,
			Message: "用户不存在",
		})
		return
	}

	// 生成新的Token
	newToken, err := jwt.GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, RegisterResponse{
			Code:    500,
			Message: "Token生成失败",
		})
		return
	}

	c.JSON(http.StatusOK, RegisterResponse{
		Code:    200,
		Message: "Token刷新成功",
		Data: map[string]interface{}{
			"token": newToken,
		},
	})
}

// GetCurrentUser 获取当前用户信息接口
// 从上下文中获取用户ID，查询并返回当前登录用户的信息
func (h *AuthHandler) GetCurrentUser(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, RegisterResponse{
			Code:    401,
			Message: "未授权",
		})
		return
	}

	// 将接口类型转换为int64
	var userIDInt int64
	switch v := userID.(type) {
	case int64:
		userIDInt = v
	case float64:
		userIDInt = int64(v)
	case string:
		parsedID, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			c.JSON(http.StatusInternalServerError, RegisterResponse{
				Code:    500,
				Message: "用户ID转换错误",
			})
			return
		}
		userIDInt = parsedID
	default:
		c.JSON(http.StatusInternalServerError, RegisterResponse{
			Code:    500,
			Message: "用户ID类型错误",
		})
		return
	}

	// 根据用户ID获取用户信息
	user, err := h.authService.GetUserByID(userIDInt)
	if err != nil || user == nil {
		c.JSON(http.StatusNotFound, RegisterResponse{
			Code:    404,
			Message: "用户不存在",
		})
		return
	}

	c.JSON(http.StatusOK, RegisterResponse{
		Code:    200,
		Message: "获取用户信息成功",
		Data: map[string]interface{}{
			"user": user,
		},
	})
}
