package model

import (
	"time"

	"gorm.io/gorm"
)

// InvoiceStatus 发票状态枚举
type InvoiceStatus string

const (
	StatusPending  InvoiceStatus = "pending"  // Waiting for audit
	StatusApproved InvoiceStatus = "approved" // Ready for reimbursement
	StatusRejected InvoiceStatus = "rejected" // Sent back to user
	StatusPaid     InvoiceStatus = "paid"     // Archived
)

// Invoice 发票模型
// Simplified: stores data AFTER AI processing or manual entry.
type Invoice struct {
	ID            int64          `gorm:"primaryKey;autoIncrement" json:"id"`
	TenantID      *int64         `gorm:"index" json:"tenant_id,omitempty"`
	UserID        int64          `gorm:"index;not null" json:"user_id"`
	AttachmentID  int64          `gorm:"index;" json:"attachment_id"` // Linked Source File
	InvoiceNumber string         `gorm:"size:100;index" json:"invoice_number"`
	InvoiceDate   time.Time      `json:"invoice_date"`
	Amount        float64        `json:"amount"`
	Vendor        string         `gorm:"size:255" json:"vendor"`
	Category      string         `gorm:"size:100" json:"category"`
	Description   string         `gorm:"size:500" json:"description"`
	Status        InvoiceStatus  `gorm:"size:50;default:'pending';index" json:"status"`
	ReviewerID    *int64         `gorm:"index" json:"reviewer_id,omitempty"`
	ReviewRemarks string         `gorm:"size:255" json:"review_remarks"`
	PaidAt        *time.Time     `json:"paid_at,omitempty"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}
