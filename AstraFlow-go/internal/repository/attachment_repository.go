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
	GetByFilePath(filePath string) (*model.Attachment, error)
	UpdateInvoiceID(fileID, invoiceID int64) error
	GetByStatus(status model.AttachmentStatus) ([]*model.Attachment, error)
	UpdateStatus(fileID int64, status model.AttachmentStatus) error
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
	err := r.db.Where("tenant_id = ?", tenantID).
		Order("status DESC").
		Order("created_at DESC").
		Find(&attachments).Error
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

func (r *attachmentRepository) GetByFilePath(filePath string) (*model.Attachment, error) {
	var attachment model.Attachment
	err := r.db.Where("file_url = ?", filePath).First(&attachment).Error
	if err != nil {
		return nil, err
	}
	return &attachment, nil
}

func (r *attachmentRepository) UpdateInvoiceID(fileID, invoiceID int64) error {
	return r.db.Model(&model.Attachment{}).Where("id = ?", fileID).Update("invoice_id", invoiceID).Error
}

func (r *attachmentRepository) GetByStatus(status model.AttachmentStatus) ([]*model.Attachment, error) {
	var attachments []*model.Attachment
	err := r.db.Where("status = ?", status).Find(&attachments).Error
	if err != nil {
		return nil, err
	}
	return attachments, nil
}

func (r *attachmentRepository) UpdateStatus(fileID int64, status model.AttachmentStatus) error {
	return r.db.Model(&model.Attachment{}).Where("id = ?", fileID).Update("status", status).Error
}
