package model

import (
	"time"

	"gorm.io/gorm"
)

// Reimbursement 报销单模型
// 定义了报销单的基本信息和状态
type Reimbursement struct {
	ID          int64          `gorm:"primaryKey;autoIncrement" json:"id"`
	TenantID    *int64         `gorm:"index" json:"tenant_id,omitempty"`    // 租户ID（个人用户为 NULL）
	UserID      int64          `gorm:"index;not null" json:"user_id"`       // 申请人
	Title       string         `gorm:"size:255;not null" json:"title"`      // 报销单标题
	TotalAmount *float64       `json:"total_amount,omitempty"`              // 总金额
	Status      string         `gorm:"size:50;default:draft" json:"status"` // 状态：draft/submitted/approved/rejected/paid
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"` // 软删除时间戳
}
