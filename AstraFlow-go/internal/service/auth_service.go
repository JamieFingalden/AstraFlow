package service

import (
	"AstraFlow-go/internal/model"
	"AstraFlow-go/internal/repository"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// AuthService 认证服务
// 处理用户注册、登录等认证相关业务逻辑
type AuthService struct {
	userRepo *repository.UserRepository // 用户数据访问仓库
}

// NewAuthService 创建认证服务实例
// 初始化认证服务并创建用户仓库依赖
func NewAuthService() *AuthService {
	return &AuthService{
		userRepo: repository.NewUserRepository(),
	}
}

// Register 用户注册
// 验证用户名唯一性，加密密码，并创建新用户
// 默认创建个人用户（personal role），当tenantID为nil时
func (s *AuthService) Register(username, password, email, phone string, tenantID *int64) (*model.User, error) {
	// 检查用户名是否已存在
	existingUser, err := s.userRepo.FindByUsername(username)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("用户名已存在")
	}

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("密码加密失败")
	}

	// 创建用户
	user := &model.User{
		Username: username,
		Password: string(hashedPassword),
		Email:    &email,
		Phone:    &phone,
		TenantID: tenantID,
	}

	// 如果没有指定租户ID，则为个人用户，设置角色为"personal"
	if tenantID == nil {
		user.RoleName = "personal"
	} else {
		// 如果指定了租户ID，则为租户用户，设置角色为"normal"
		user.RoleName = "normal"
	}

	err = s.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	// 返回用户信息（不包含密码）
	user.Password = "" // 不返回密码
	return user, nil
}

// Login 用户登录
// 验证用户名和密码，返回用户信息
func (s *AuthService) Login(username, password string) (*model.User, error) {
	user, err := s.userRepo.FindByUsername(username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("用户不存在")
	}

	// 验证密码
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("密码错误")
	}

	// 返回用户信息（不包含密码）
	user.Password = "" // 不返回密码
	return user, nil
}

// GetUserByID 根据ID获取用户
// 根据用户ID查询用户信息，不返回密码字段
func (s *AuthService) GetUserByID(id int64) (*model.User, error) {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if user != nil {
		user.Password = "" // 不返回密码
	}
	return user, nil
}
