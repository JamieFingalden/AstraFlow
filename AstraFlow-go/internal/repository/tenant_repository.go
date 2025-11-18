package repository

import (
	"AstraFlow-go/internal/database"
	"AstraFlow-go/internal/model"

	"gorm.io/gorm"
)

// TenantRepository 租户数据访问仓库
// 封装对租户表的CRUD操作
type TenantRepository struct {
	db *gorm.DB // 数据库连接实例
}

// NewTenantRepository 创建租户仓库实例
// 初始化时使用全局数据库连接
func NewTenantRepository() *TenantRepository {
	return &TenantRepository{
		db: database.DB,
	}
}

// Create 创建租户
// 将新租户信息插入数据库
func (r *TenantRepository) Create(tenant *model.Tenant) error {
	return r.db.Create(tenant).Error
}

// FindByID 根据ID查找租户
// 返回匹配的租户对象，如果未找到则返回nil
func (r *TenantRepository) FindByID(id int64) (*model.Tenant, error) {
	var tenant model.Tenant
	err := r.db.Where("id = ?", id).First(&tenant).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &tenant, nil
}

// FindAll 获取所有租户列表
// 返回租户列表和错误信息
func (r *TenantRepository) FindAll(limit, offset int) ([]*model.Tenant, error) {
	var tenants []*model.Tenant
	err := r.db.Limit(limit).Offset(offset).Find(&tenants).Error
	return tenants, err
}

// Update 更新租户
// 更新租户信息到数据库
func (r *TenantRepository) Update(tenant *model.Tenant) error {
	return r.db.Save(tenant).Error
}

// Delete 删除租户
// 根据租户ID从数据库中删除租户记录
func (r *TenantRepository) Delete(id int64) error {
	return r.db.Delete(&model.Tenant{}, id).Error
}

// FindByName 根据名称查找租户
// 返回匹配的租户对象，如果未找到则返回nil
func (r *TenantRepository) FindByName(name string) (*model.Tenant, error) {
	var tenant model.Tenant
	err := r.db.Where("name = ?", name).First(&tenant).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &tenant, nil
}
