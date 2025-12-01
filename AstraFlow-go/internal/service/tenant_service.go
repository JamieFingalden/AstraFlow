package service

import (
	"AstraFlow-go/internal/model"
	"AstraFlow-go/internal/repository"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// TenantService 租户服务
// 处理租户相关的业务逻辑
type TenantService struct {
	tenantRepo *repository.TenantRepository // 租户数据访问仓库
	userRepo   *repository.UserRepository   // 用户数据访问仓库
}

// NewTenantService 创建租户服务实例
// 初始化租户服务并创建租户仓库依赖
func NewTenantService() *TenantService {
	return &TenantService{
		tenantRepo: repository.NewTenantRepository(),
		userRepo:   repository.NewUserRepository(),
	}
}

// CreateTenant 创建租户
// 验证租户名称唯一性，并创建新租户
func (s *TenantService) CreateTenant(tenantName, username, email, password string) (*model.Tenant, *model.User, error) {
	// 检查租户名称是否已存在
	existingTenant, err := s.tenantRepo.FindByName(tenantName)
	if err != nil {
		return nil, nil, err
	}
	if existingTenant != nil {
		return nil, nil, errors.New("租户名称已存在")
	}

	existingUser, err := s.userRepo.FindByUsername(username)
	if err != nil {
		return nil, nil, err
	}
	if existingUser != nil {
		return nil, nil, errors.New("用户名已存在")
	}

	// 创建租户
	tenant := &model.Tenant{
		Name:         tenantName,
		ContactName:  &username,
		ContactEmail: &email,
	}

	err = s.tenantRepo.Create(tenant)
	if err != nil {
		return nil, nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, nil, err
	}

	user := &model.User{
		TenantID: &tenant.ID,
		Username: username,
		Password: string(hashedPassword),
		Role:     "admin",
	}

	err = s.userRepo.Create(user)
	if err != nil {
		return nil, nil, err
	}

	return tenant, user, nil
}

// GetTenantByID 根据ID获取租户
// 根据租户ID查询租户信息
func (s *TenantService) GetTenantByID(id int64) (*model.Tenant, error) {
	tenant, err := s.tenantRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return tenant, nil
}

// GetTenantList 获取租户列表
// 支持分页查询租户列表
func (s *TenantService) GetTenantList(page, pageSize int) ([]*model.Tenant, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize
	limit := pageSize

	tenants, err := s.tenantRepo.FindAll(limit, offset)
	if err != nil {
		return nil, err
	}

	return tenants, nil
}

// UpdateTenant 更新租户信息
// 根据租户ID更新租户信息
func (s *TenantService) UpdateTenant(id int64, name, industry, contactName, contactPhone, contactEmail string) (*model.Tenant, error) {
	// 检查租户是否存在
	existingTenant, err := s.tenantRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if existingTenant == nil {
		return nil, errors.New("租户不存在")
	}

	// 如果更新了名称，检查新名称是否与其他租户冲突
	if name != existingTenant.Name {
		conflictingTenant, err := s.tenantRepo.FindByName(name)
		if err != nil {
			return nil, err
		}
		if conflictingTenant != nil && conflictingTenant.ID != id {
			return nil, errors.New("租户名称已存在")
		}
	}

	// 更新租户信息
	existingTenant.Name = name
	existingTenant.Industry = &industry
	existingTenant.ContactName = &contactName
	existingTenant.ContactPhone = &contactPhone
	existingTenant.ContactEmail = &contactEmail

	err = s.tenantRepo.Update(existingTenant)
	if err != nil {
		return nil, err
	}

	return existingTenant, nil
}

// DeleteTenant 删除租户
// 根据租户ID删除租户
func (s *TenantService) DeleteTenant(id int64) error {
	tenant, err := s.tenantRepo.FindByID(id)
	if err != nil {
		return err
	}
	if tenant == nil {
		return errors.New("租户不存在")
	}

	err = s.tenantRepo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
