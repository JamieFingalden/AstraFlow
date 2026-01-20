package model

import (
	"time"

	"gorm.io/gorm"
)

// AttachmentStatus 附件识别状态
type AttachmentStatus string

const (
	AttachmentStatusPending   AttachmentStatus = "pending"   // 识别中
	AttachmentStatusSuccess   AttachmentStatus = "success"   // 识别成功
	AttachmentStatusFailed    AttachmentStatus = "failed"    // 识别失败
	AttachmentStatusDuplicate AttachmentStatus = "duplicate" // 重复识别
)

// Attachment 附件模型
// 定义了发票照片/支付截图等附件信息
type Attachment struct {
	ID        int64            `gorm:"primaryKey;autoIncrement" json:"id"`
	TenantID  *int64           `gorm:"index" json:"tenant_id,omitempty"`        // 租户ID（个人用户为 NULL）
	UserID    int64            `gorm:"index;not null" json:"user_id"`           // 上传者 ID
	InvoiceID *int64           `gorm:"index" json:"invoice_id,omitempty"`       // 关联发票 ID，可为空（软连接）
	FileURL   string           `gorm:"size:500;not null" json:"file_url"`       // 文件存储路径
	FileType  *string          `gorm:"size:50" json:"file_type,omitempty"`      // 文件类型：jpg/png/pdf 等
	FileSize  *int64           `json:"file_size,omitempty"`                     // 大小（字节）
	Status    AttachmentStatus `gorm:"size:20;default:'pending'" json:"status"` // 识别状态
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
	DeletedAt gorm.DeletedAt   `gorm:"index" json:"-"` // 软删除时间戳
}
