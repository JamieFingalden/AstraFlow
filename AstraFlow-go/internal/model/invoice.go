package model

import (
	"time"

	"gorm.io/gorm"
)

// Invoice 发票模型
// 定义了发票的基本信息和状态
type InvoiceStatus string

const (
	StatusPending  InvoiceStatus = "pending"  // 待审核 (AI识别完生成记录)
	StatusApproved InvoiceStatus = "approved" // 审核通过 (待报销)
	StatusRejected InvoiceStatus = "rejected" // 已驳回
	StatusPaid     InvoiceStatus = "paid"     // 已报销 (归档)
)

type Invoice struct {
	ID             int64          `gorm:"primaryKey;autoIncrement" json:"id"`
	TenantID       *int64         `gorm:"index" json:"tenant_id,omitempty"`                   // 租户ID，NULL 表示个人用户
	UserID         int64          `gorm:"index;not null" json:"user_id"`                      // 上传用户ID
	InvoiceNumber  *string        `gorm:"size:100" json:"invoice_number,omitempty"`           // 发票编号
	InvoiceDate    *time.Time     `json:"invoice_date,omitempty"`                             // 开票日期
	Amount         *float64       `json:"amount,omitempty"`                                   // 金额
	Vendor         string         `gorm:"size:255;not null" json:"vendor"`                    // 商户名称，必填
	ImageURL       *string        `gorm:"size:500" json:"image_url,omitempty"`                // 发票图片URL
	Description    *string        `gorm:"size:500" json:"description,omitempty"`              // 发票描述/备注
	TaxID          *string        `gorm:"size:100" json:"tax_id,omitempty"`                   // 税号
	Category       *string        `gorm:"size:100" json:"category,omitempty"`                 // 分类：餐饮、交通、购物、住宿、办公用品等
	PaymentSource  *string        `gorm:"size:50" json:"payment_source,omitempty"`            // 支付方式：微信/支付宝/现金/银行卡 等
	Status         InvoiceStatus  `gorm:"size:50;default:pending;index" json:"status"`        // 状态
	ReviewerID     *int64         `gorm:"index" json:"reviewer_id,omitempty"`                 // 审核人ID (关联User)
	ReviewRemarks  *string        `gorm:"size:255" json:"review_remarks,omitempty"`           // 审核备注/驳回原因
	PaidAt         *time.Time     `json:"paid_at,omitempty"`
	IdentifyStatus string         `gorm:"size:50;default:unnecessary" json:"identify_status"` // 识别状态：unnecessary/Identifying/failures/identified
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"` // 软删除时间戳
}
