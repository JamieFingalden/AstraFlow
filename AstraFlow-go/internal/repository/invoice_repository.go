package repository

import (
	"AstraFlow-go/internal/database"
	"AstraFlow-go/internal/model"
	"errors"

	"gorm.io/gorm"
)

// InvoiceRepository 封装发票的 CRUD 操作
type InvoiceRepository struct {
	db *gorm.DB
}

// NewInvoiceRepository 初始化实例
func NewInvoiceRepository() *InvoiceRepository {
	return &InvoiceRepository{db: database.DB}
}

// Create 创建发票
func (r *InvoiceRepository) Create(invoice *model.Invoice) error {
	return r.db.Create(invoice).Error
}

// FindByInvoiceNumber 根据发票编号查询发票
func (r *InvoiceRepository) FindByInvoiceNumber(InvoiceNumber string) (*model.Invoice, error) {
	var invoice model.Invoice
	err := r.db.Where("Invoice_number = ?", InvoiceNumber).First(&invoice).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &invoice, nil
}

// FindAllPage 分页查询所有发票
func (r *InvoiceRepository) FindAllPage(limit, offset int) ([]*model.Invoice, error) {
	var invoices []*model.Invoice
	err := r.db.Limit(limit).Offset(offset).Find(&invoices).Error
	return invoices, err
}
