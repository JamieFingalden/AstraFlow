package handler

import (
	"AstraFlow-go/internal/service"
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

// UserResponse 用户响应结构体
type UserResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// 获取用户列表
func (h *UserHandler) GetUserList(c *gin.Context) {
	tenantId, exists := c.Get("tenantId")
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

}

// 更新用户
func (h *UserHandler) UpdateUser(c *gin.Context) {

}

// 删除用户
func (h *UserHandler) DeleteUser(c *gin.Context) {

}
