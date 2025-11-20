package model

import (
	"time"

	"gorm.io/gorm"
)

// User 用户模型
// 定义了用户的基本信息和关联关系
type User struct {
	ID        int64          `gorm:"primaryKey;autoIncrement" json:"id"`
	TenantID  *int64         `gorm:"index" json:"tenant_id,omitempty"` // 租户ID，NULL 表示个人用户
	RoleID    *int64         `gorm:"index" json:"role_id,omitempty"`   // 角色ID，引用role表，NULL表示默认普通用户
	Username  string         `gorm:"uniqueIndex;size:50;not null" json:"username"`
	Password  string         `gorm:"size:255;not null" json:"-"` // 密码字段，JSON序列化时忽略
	Email     *string        `gorm:"size:100" json:"email,omitempty"`
	Phone     *string        `gorm:"size:50" json:"phone,omitempty"`
	Role      string         `gorm:"size:30;default:personal" json:"role,omitempty"` // 用户角色（保持兼容性）
	RoleName  string         `gorm:"-" json:"-"`                                     // 角色名称（内部使用）
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 软删除时间戳
}
