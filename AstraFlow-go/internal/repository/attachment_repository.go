package repository

import (
	"AstraFlow-go/internal/database"
	"AstraFlow-go/internal/model"

	"gorm.io/gorm"
)

type AttachmentRepository interface {
	Create(attachment *model.Attachment) error
	GetByID(id int64) (*model.Attachment, error)
	GetByUserID(userID int64) ([]*model.Attachment, error)
	GetByTenantID(tenantID int64) ([]*model.Attachment, error)
	Update(attachment *model.Attachment) error
	Delete(id int64) error
	GetByInvoiceID(invoiceID int64) ([]*model.Attachment, error)
}

type attachmentRepository struct {
	db *gorm.DB
}

func NewAttachmentRepository() AttachmentRepository {
	return &attachmentRepository{
		db: database.DB,
	}
}

func (r *attachmentRepository) Create(attachment *model.Attachment) error {
	return r.db.Create(attachment).Error
}

func (r *attachmentRepository) GetByID(id int64) (*model.Attachment, error) {
	var attachment model.Attachment
	err := r.db.Where("id = ?", id).First(&attachment).Error
	if err != nil {
		return nil, err
	}
	return &attachment, nil
}

func (r *attachmentRepository) GetByUserID(userID int64) ([]*model.Attachment, error) {
	var attachments []*model.Attachment
	err := r.db.Where("user_id = ?", userID).Find(&attachments).Error
	if err != nil {
		return nil, err
	}
	return attachments, nil
}

func (r *attachmentRepository) GetByTenantID(tenantID int64) ([]*model.Attachment, error) {
	var attachments []*model.Attachment
	err := r.db.Where("tenant_id = ?", tenantID).Find(&attachments).Error
	if err != nil {
		return nil, err
	}
	return attachments, nil
}

func (r *attachmentRepository) Update(attachment *model.Attachment) error {
	return r.db.Save(attachment).Error
}

func (r *attachmentRepository) Delete(id int64) error {
	return r.db.Delete(&model.Attachment{}, id).Error
}

func (r *attachmentRepository) GetByInvoiceID(invoiceID int64) ([]*model.Attachment, error) {
	var attachments []*model.Attachment
	err := r.db.Where("invoice_id = ?", invoiceID).Find(&attachments).Error
	if err != nil {
		return nil, err
	}
	return attachments, nil
}
