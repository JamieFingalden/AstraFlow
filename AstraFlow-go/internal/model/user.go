package model

import (
	"time"

	"gorm.io/gorm"
)

// User Status Constants
const (
	UserStatusActive   = 1
	UserStatusDisabled = 2
)

// User 用户模型
// Refactored to use RoleID and Status, with Role association for Preloading
type User struct {
	ID        int64          `gorm:"primaryKey;autoIncrement" json:"id"`
	TenantID  *int64         `gorm:"index" json:"tenant_id,omitempty"`        // 租户ID，NULL 表示个人用户
	RoleID    int64          `gorm:"index;not null" json:"role_id"`           // 角色ID
	Role      *Role          `gorm:"foreignKey:RoleID" json:"role,omitempty"` // Association for easy access
	Username  string         `gorm:"uniqueIndex;size:50;not null" json:"username"`
	Password  string         `gorm:"size:255;not null" json:"-"` // 密码字段，JSON序列化时忽略
	Name      string         `gorm:"size:100" json:"name"`
	Email     *string        `gorm:"size:100" json:"email,omitempty"`
	Phone     *string        `gorm:"size:50" json:"phone,omitempty"`
	Status    int            `gorm:"default:1" json:"status"` // 1: Active, 2: Disabled
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 软删除时间戳
}
