package service

import (
	"AstraFlow-go/internal/model"
	"log"

	"gorm.io/gorm"
)

// EnsureSystemRoles checks if the 3 fixed roles exist and creates them if not.
func EnsureSystemRoles(db *gorm.DB) error {
	roles := []model.Role{
		{
			Key:         model.RoleKeyAdmin,
			Name:        "Administrator",
			Description: "System Administrator with full access",
		},
		{
			Key:         model.RoleKeyAuditor,
			Name:        "Auditor",
			Description: "Financial Auditor can review and approve invoices",
		},
		{
			Key:         model.RoleKeyEmployee,
			Name:        "Employee",
			Description: "Regular employee can upload invoices",
		},
	}

	for _, r := range roles {
		var count int64
		// Using `Key` to check for existence as it is unique
		if err := db.Model(&model.Role{}).Where("`key` = ?", r.Key).Count(&count).Error; err != nil {
			return err
		}

		if count == 0 {
			if err := db.Create(&r).Error; err != nil {
				return err
			}
			log.Printf("Seeded system role: %s (%s)", r.Name, r.Key)
		}
	}

	return nil
}
