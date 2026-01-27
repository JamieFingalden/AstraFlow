package model

import (
	"time"

	"gorm.io/gorm"
)

// Define fixed roles as constants
const (
	RoleKeyAdmin    = "admin"
	RoleKeyAuditor  = "auditor"
	RoleKeyEmployee = "employee"
)

// Role 角色模型
// The system has 3 FIXED roles. Do not allow dynamic role creation.
type Role struct {
	ID          int64          `gorm:"primaryKey;autoIncrement" json:"id"`
	Key         string         `gorm:"size:50;uniqueIndex;not null" json:"key"` // e.g., "admin"
	Name        string         `gorm:"size:50;not null" json:"name"`            // e.g., "HR Administrator"
	Description string         `gorm:"size:255" json:"description"`             // 角色描述
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
