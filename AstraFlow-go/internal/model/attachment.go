package model

import (
	"time"

	"gorm.io/gorm"
)

// AttachmentStatus 附件状态枚举
type AttachmentStatus int

const (
	AttachmentStatusPending AttachmentStatus = 0 // Uploaded, waiting for AI
	AttachmentStatusSuccess AttachmentStatus = 1 // AI processed, Invoice created
	AttachmentStatusFailed  AttachmentStatus = 2 // OCR failed
)

// Attachment 附件模型
// Represents the uploaded file and holds the InvoiceID once created
type Attachment struct {
	ID        int64            `gorm:"primaryKey;autoIncrement" json:"id"`
	TenantID  *int64           `gorm:"index" json:"tenant_id,omitempty"`  // 租户ID
	UserID    int64            `gorm:"index;not null" json:"user_id"`     // 上传者 ID
	InvoiceID *int64           `gorm:"index" json:"invoice_id,omitempty"` // 关联发票 ID (Foreign Key)
	FileName  string           `gorm:"size:255" json:"file_name"`         // 原始文件名
	FileURL   string           `gorm:"size:500;not null" json:"file_url"` // 文件存储路径
	Status    AttachmentStatus `gorm:"default:0" json:"status"`           // 状态
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
	DeletedAt gorm.DeletedAt   `gorm:"index" json:"-"` // 软删除时间戳
}