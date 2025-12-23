package model

import (
	"time"

	"gorm.io/gorm"
)

// ReimbursementItem 报销单子项模型
// 定义了报销单中关联发票的子项信息
type ReimbursementItem struct {
	ID              int64          `gorm:"primaryKey;autoIncrement" json:"id"`
	ReimbursementID int64          `gorm:"index;not null" json:"reimbursement_id"` // 报销单 ID（软连接）
	InvoiceID       int64          `gorm:"index;not null" json:"invoice_id"`       // 发票 ID（软连接）
	Amount          *float64       `json:"amount,omitempty"`                       // 发票金额快照
	Note            *string        `gorm:"size:255" json:"note,omitempty"`         // 备注
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"` // 软删除时间戳
}
