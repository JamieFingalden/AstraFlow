package repository

import (
	"AstraFlow-go/internal/database"
	"AstraFlow-go/internal/model"

	"gorm.io/gorm"
)

type RoleRepository struct {
	db *gorm.DB
}

func NewRoleRepository() *RoleRepository {
	return &RoleRepository{
		db: database.DB,
	}
}

func (r *RoleRepository) FindByKey(key string) (*model.Role, error) {
	var role model.Role
	err := r.db.Where("`key` = ?", key).First(&role).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}
