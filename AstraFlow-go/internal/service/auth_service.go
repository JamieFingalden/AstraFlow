package service

import (
	"AstraFlow-go/internal/database"
	"AstraFlow-go/internal/model"
	"AstraFlow-go/internal/repository"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// AuthService 认证服务
type AuthService struct {
	userRepo   *repository.UserRepository
	roleRepo   *repository.RoleRepository
	tenantRepo *repository.TenantRepository
}

// NewAuthService 创建认证服务实例
func NewAuthService() *AuthService {
	return &AuthService{
		userRepo:   repository.NewUserRepository(),
		roleRepo:   repository.NewRoleRepository(),
		tenantRepo: repository.NewTenantRepository(),
	}
}

// RegisterTenant 处理租户注册（企业入驻）
func (s *AuthService) RegisterTenant(companyName, adminUsername, password, adminName, phone string) (*model.User, error) {
	var user *model.User
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		var count int64
		if err := tx.Model(&model.User{}).Where("username = ?", adminUsername).Count(&count).Error; err != nil {
			return err
		}
		if count > 0 {
			return errors.New("username already exists")
		}

		if err := tx.Model(&model.Tenant{}).Where("name = ?", companyName).Count(&count).Error; err != nil {
			return err
		}
		if count > 0 {
			return errors.New("company name already exists")
		}

		var adminRole model.Role
		if err := tx.Where("`key` = ?", model.RoleKeyAdmin).First(&adminRole).Error; err != nil {
			return errors.New("system error: admin role not found")
		}

		tenant := model.Tenant{
			Name:         companyName,
			ContactName:  &adminName,
			ContactPhone: &phone,
		}
		if err := tx.Create(&tenant).Error; err != nil {
			return err
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		user = &model.User{
			TenantID: &tenant.ID,
			RoleID:   adminRole.ID,
			Username: adminUsername,
			Password: string(hashedPassword),
			Name:     adminName,
			Phone:    &phone,
			Status:   model.UserStatusActive,
		}

		if err := tx.Create(user).Error; err != nil {
			return err
		}

		user.Role = &adminRole
		return nil
	})

	if err != nil {
		return nil, err
	}

	user.Password = ""
	return user, nil
}

// Login 用户登录
// 支持用户名或邮箱登录
func (s *AuthService) Login(username, email, password string) (*model.User, error) {
	var user *model.User
	var err error

	// Strategy:
	// 1. If username provided -> Try finding by username
	// 2. If email provided -> Try finding by email
	// 3. (Fallback) If email provided but not found in email col -> Try finding by username (in case email is used as username)
	// 4. (Fallback) If username provided but not found in username col -> Try finding by email (in case username input is actually an email)

	// Attempt 1: Explicit match
	if username != "" {
		user, err = s.userRepo.FindByUsername(username)
	}
	
	if (user == nil || err != nil) && email != "" {
		user, err = s.userRepo.FindByEmail(email)
	}

	// Attempt 2: Cross match (Flexible Login)
	// If still not found, check if the input 'email' is actually stored as 'username'
	if (user == nil || err != nil) && email != "" {
		user, err = s.userRepo.FindByUsername(email)
	}
	// If still not found, check if the input 'username' is actually stored as 'email'
	if (user == nil || err != nil) && username != "" {
		user, err = s.userRepo.FindByEmail(username)
	}

	if err != nil || user == nil {
		return nil, errors.New("invalid username/email or password")
	}

	if user.Status == model.UserStatusDisabled {
		return nil, errors.New("account is disabled")
	}

	// 验证密码
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid username/email or password")
	}

	// 确保 Role 被加载
	if user.Role == nil {
		if user.RoleID != 0 {
			var role model.Role
			if dbErr := database.DB.First(&role, user.RoleID).Error; dbErr == nil {
				user.Role = &role
			}
		}
		if user.Role == nil {
			return nil, fmt.Errorf("user has no role assigned, ID: %d", user.ID)
		}
	}

	user.Password = ""
	return user, nil
}

// GetUserByID 根据ID获取用户
func (s *AuthService) GetUserByID(id int64) (*model.User, error) {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if user != nil {
		user.Password = ""
	}
	return user, nil
}
