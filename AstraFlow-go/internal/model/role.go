package model

import (
	"time"

	"gorm.io/gorm"
)

// Role 角色模型
// 定义了用户角色的基本信息
type Role struct {
	ID          int64          `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string         `gorm:"size:50;not null;uniqueIndex" json:"name"`        // 角色名称
	DisplayName string         `gorm:"size:100;not null" json:"display_name"`          // 角色显示名称
	Description string         `gorm:"size:255" json:"description"`                     // 角色描述
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}