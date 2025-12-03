package service

import (
	"AstraFlow-go/internal/model"
	"AstraFlow-go/internal/repository"
)

type UserService struct {
	userRepo *repository.UserRepository // 用户仓库实例
}

// NewUserService 创建用户服务实例
// 初始化用户服务
func NewUserService() *UserService {
	return &UserService{
		userRepo: repository.NewUserRepository(),
	}
}

// GetUserList 获取用户列表
func (s *UserService) GetUserList(tenantId int64) ([]*model.User, error) {
	return s.userRepo.GetUserList(tenantId)
}
