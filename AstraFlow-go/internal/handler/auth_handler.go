package handler

import (
	"AstraFlow-go/internal/service"
	"AstraFlow-go/pkg/jwt"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		authService: service.NewAuthService(),
	}
}

// RegisterRequest 租户注册请求体
type RegisterRequest struct {
	CompanyName   string `json:"company_name" binding:"required,min=2,max=100"`
	AdminUsername string `json:"admin_username" binding:"required,min=3,max=50"`
	Password      string `json:"password" binding:"required,min=6,max=100"`
	AdminName     string `json:"admin_name" binding:"required"`
	Phone         string `json:"phone" binding:"required"`
}

// LoginRequest 登录请求体
// Username 和 Email 至少需要一个，但在 Bind 时我们允许为空，由 Service 层处理逻辑
type LoginRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password" binding:"required"`
}

// Register 租户注册接口 (企业入驻)
func (h *AuthHandler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "请求参数错误: " + err.Error()})
		return
	}

	user, err := h.authService.RegisterTenant(
		req.CompanyName,
		req.AdminUsername,
		req.Password,
		req.AdminName,
		req.Phone,
	)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"code": 409, "message": err.Error()})
		return
	}

	token, err := jwt.GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "Token生成失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "企业注册成功",
		"data": map[string]interface{}{
			"user":  user,
			"token": token,
		},
	})
}

// Login 用户登录接口
func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "请求参数错误: " + err.Error()})
		return
	}

	if req.Username == "" && req.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "必须提供用户名或邮箱"})
		return
	}

	user, err := h.authService.Login(req.Username, req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": err.Error()})
		return
	}

	token, err := jwt.GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "Token生成失败"})
		return
	}

	refreshToken, err := jwt.GenerateRefreshToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "Refresh Token生成失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "登录成功",
		"data": map[string]interface{}{
			"user":          user,
			"token":         token,
			"refresh_token": refreshToken,
		},
	})
}

// RefreshToken 刷新Token接口
func (h *AuthHandler) RefreshToken(c *gin.Context) {
	refreshToken := c.GetHeader("Authorization")
	if refreshToken == "" {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "缺少刷新Token"})
		return
	}

	if len(refreshToken) > 7 && refreshToken[:7] == "Bearer " {
		refreshToken = refreshToken[7:]
	}

	claims, err := jwt.ParseToken(refreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "刷新Token无效"})
		return
	}

	user, err := h.authService.GetUserByID(claims.UserID)
	if err != nil || user == nil {
		c.JSON(http.StatusNotFound, gin.H{"code": 404, "message": "用户不存在"})
		return
	}

	newToken, err := jwt.GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "Token生成失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Token刷新成功",
		"data": map[string]interface{}{
			"token": newToken,
		},
	})
}

// GetCurrentUser 获取当前用户信息
func (h *AuthHandler) GetCurrentUser(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "未授权"})
		return
	}

	var userIDInt int64
	switch v := userID.(type) {
	case int64:
		userIDInt = v
	case float64:
		userIDInt = int64(v)
	default:
		strVal := fmt.Sprintf("%v", v)
		parsed, _ := strconv.ParseInt(strVal, 10, 64)
		userIDInt = parsed
	}

	user, err := h.authService.GetUserByID(userIDInt)
	if err != nil || user == nil {
		c.JSON(http.StatusNotFound, gin.H{"code": 404, "message": "用户不存在"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取用户信息成功",
		"data": map[string]interface{}{
			"user": user,
		},
	})
}
