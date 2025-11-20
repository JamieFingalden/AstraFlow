package model

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

// OCRResult OCR识别结果模型
// 定义了OCR识别结果的结构化和非结构化数据
type OCRResult struct {
	ID            int64           `gorm:"primaryKey;autoIncrement" json:"id"`
	TenantID      *int64          `gorm:"index" json:"tenant_id,omitempty"` // 租户ID / 个人用户 NULL
	UserID        int64           `gorm:"index;not null" json:"user_id"`    // 关联用户ID
	InvoiceID     int64           `gorm:"index;not null" json:"invoice_id"` // 关联发票ID（软连接）
	Amount        *float64        `json:"amount,omitempty"`                 // 金额
	InvoiceDate   *time.Time      `json:"invoice_date,omitempty"`           // 开票日期
	Vendor        *string         `gorm:"size:255" json:"vendor,omitempty"` // 商户
	Category      *string         `gorm:"size:100" json:"category,omitempty"` // 类别
	InvoiceNumber *string         `gorm:"size:100" json:"invoice_number,omitempty"` // 发票号
	Confidence    *float64        `json:"confidence,omitempty"`                     // 识别置信度，例如 0.97
	RawText       string          `gorm:"type:longtext" json:"raw_text"`            // OCR 原始文本
	ParsedJSON    json.RawMessage `gorm:"type:json" json:"parsed_json,omitempty"`   // OCR 解析后的 JSON
	CreatedAt     time.Time       `json:"created_at"`
	UpdatedAt     time.Time       `json:"updated_at"`
	DeletedAt     gorm.DeletedAt  `gorm:"index" json:"-"` // 软删除时间戳
}