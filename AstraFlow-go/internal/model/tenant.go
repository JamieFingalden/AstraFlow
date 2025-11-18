package model

import (
	"gorm.io/gorm"
	"time"
)

// Tenant 租户模型
// 定义了企业租户的基本信息
type Tenant struct {
	ID          int64          `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string         `gorm:"size:100;not null" json:"name" binding:"required,min=1,max=100"` // 租户名称（公司名称）
	Industry    *string        `gorm:"size:100" json:"industry,omitempty"`                            // 行业
	ContactName *string        `gorm:"size:100" json:"contact_name,omitempty"`                        // 联系人姓名
	ContactPhone *string       `gorm:"size:50" json:"contact_phone,omitempty"`                        // 联系电话
	ContactEmail *string       `gorm:"size:100" json:"contact_email,omitempty"`                       // 联系邮箱
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"` // 软删除时间戳
}