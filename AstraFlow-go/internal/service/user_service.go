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

// FindUserRoleByUserId 根据用户ID查找用户角色
func (s *UserService) FindUserRoleByUserId(userId int64) (string, error) {
	return s.userRepo.FindUserRoleByUserId(userId)
}

// GetUserList 获取用户列表
func (s *UserService) GetUserList(tenantId int64) ([]*model.User, error) {
	return s.userRepo.GetUserList(tenantId)
}

// CreateUser 创建用户
func (s *UserService) CreateUser(tenantId int64, username, password, email, phone, role string) (*model.User, error) {

	user := &model.User{
		TenantID: &tenantId,
		Username: username,
		Password: password,
		Email:    &email,
		Phone:    &phone,
		Role:     role,
	}
	return user, s.userRepo.CreateUserReturnIt(user)
}

// UpdateUser 更新用户
func (s *UserService) UpdateUser(userId int64, username, password, email, phone, role string) error {
	user, err := s.userRepo.FindByID(userId)
	if err != nil {
		return err
	}

	if email != "" {
		user.Email = &email
	}
	if phone != "" {
		user.Phone = &phone
	}
	return s.userRepo.Update(user)
}

// DeleteUser 删除用户
func (s *UserService) DeleteUser(userId int64) error {
	return s.userRepo.Delete(userId)
}
