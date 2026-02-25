package model

import (
	"time"

	"gorm.io/gorm"
)

// InvoiceStatus 发票状态枚举
type InvoiceStatus string

const (
	StatusRecognizing InvoiceStatus = "recognizing" // AI OCR in progress
	StatusUnconfirmed InvoiceStatus = "unconfirmed" // AI OCR finished, waiting for user confirmation
	StatusDraft       InvoiceStatus = "draft"       // Manual entry or confirmed OCR, ready to publish
	StatusPending     InvoiceStatus = "pending"     // Waiting for audit
	StatusApproved    InvoiceStatus = "approved"    // Ready for reimbursement
	StatusRejected    InvoiceStatus = "rejected"    // Sent back to user
	StatusPaid        InvoiceStatus = "paid"        // Archived
)

// InvoiceSource 发票来源
type InvoiceSource string

const (
	SourceOCR    InvoiceSource = "ocr"
	SourceManual InvoiceSource = "manual"
)

// Invoice 发票模型
// Simplified: stores data AFTER AI processing or manual entry.
type Invoice struct {
	ID            int64          `gorm:"primaryKey;autoIncrement" json:"id"`            // 发票ID，主键，自增
	TenantID      *int64         `gorm:"index" json:"tenant_id,omitempty"`              // 租户ID，索引，可为空
	UserID        int64          `gorm:"index;not null" json:"user_id"`                 // 用户ID，索引，非空
	AttachmentID  int64          `gorm:"index;" json:"attachment_id"`                   // 附件ID，索引，关联源文件
	InvoiceNumber string         `gorm:"size:100;index" json:"invoice_number"`          // 发票号码，大小100，索引
	InvoiceDate   time.Time      `json:"invoice_date"`                                  // 发票日期
	Amount        float64        `json:"amount"`                                        // 金额
	Vendor        string         `gorm:"size:255" json:"vendor"`                        // 供应商，大小255
	Category      string         `gorm:"size:100" json:"category"`                      // 类别，大小100
	Description   string         `gorm:"size:500" json:"description"`                   // 描述，大小500
	Status        InvoiceStatus  `gorm:"size:50;default:'pending';index" json:"status"` // 状态，大小50，默认'pending'，索引
	Source        InvoiceSource  `gorm:"size:20;index" json:"source"`                   // 来源：ocr 或 manual
	ReviewerID    *int64         `gorm:"index" json:"reviewer_id,omitempty"`            // 审核人ID，索引，可为空
	ReviewRemarks string         `gorm:"size:255" json:"review_remarks"`                // 审核备注，大小255
	PaidAt        *time.Time     `json:"paid_at,omitempty"`                             // 支付时间，可为空
	CreatedAt     time.Time      `json:"created_at"`                                    // 创建时间
	UpdatedAt     time.Time      `json:"updated_at"`                                    // 更新时间
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`                                // 删除时间，软删除，索引，JSON不返回
}
