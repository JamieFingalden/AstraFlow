package model

import (
	"time"

	"gorm.io/gorm"
)

// Attachment 附件模型
// 定义了发票照片/支付截图等附件信息
type Attachment struct {
	ID         int64          `gorm:"primaryKey;autoIncrement" json:"id"`
	TenantID   *int64         `gorm:"index" json:"tenant_id,omitempty"`       // 租户ID（个人用户为 NULL）
	UserID     int64          `gorm:"index;not null" json:"user_id"`          // 上传者 ID
	InvoiceID  *int64         `gorm:"index" json:"invoice_id,omitempty"`      // 关联发票 ID，可为空（软连接）
	FileURL    string         `gorm:"size:500;not null" json:"file_url"`      // 文件存储路径
	FileType   *string        `gorm:"size:50" json:"file_type,omitempty"`     // 文件类型：jpg/png/pdf 等
	FileSize   *int64         `json:"file_size,omitempty"`                    // 大小（字节）
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"` // 软删除时间戳
}