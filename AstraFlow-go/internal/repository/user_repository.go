package repository

import (
	"AstraFlow-go/internal/database"
	"AstraFlow-go/internal/model"
	"gorm.io/gorm"
)

// UserRepository 用户数据访问仓库
// 封装对用户表的CRUD操作
type UserRepository struct {
	db *gorm.DB // 数据库连接实例
}

// NewUserRepository 创建用户仓库实例
// 初始化时使用全局数据库连接
func NewUserRepository() *UserRepository {
	return &UserRepository{
		db: database.DB,
	}
}

// Create 创建用户
// 将新用户信息插入数据库
func (r *UserRepository) Create(user *model.User) error {
	return r.db.Create(user).Error
}

// FindByUsername 根据用户名查找用户
// 返回匹配的用户对象，如果未找到则返回nil
func (r *UserRepository) FindByUsername(username string) (*model.User, error) {
	var user model.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// FindByID 根据ID查找用户
// 返回匹配的用户对象，如果未找到则返回nil
func (r *UserRepository) FindByID(id int64) (*model.User, error) {
	var user model.User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// Update 更新用户
// 更新用户信息到数据库
func (r *UserRepository) Update(user *model.User) error {
	return r.db.Save(user).Error
}

// Delete 删除用户
// 根据用户ID从数据库中删除用户记录
func (r *UserRepository) Delete(id int64) error {
	return r.db.Delete(&model.User{}, id).Error
}