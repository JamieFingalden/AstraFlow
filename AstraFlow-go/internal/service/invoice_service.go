package service

import (
	"AstraFlow-go/internal/model"
	"AstraFlow-go/internal/repository"
	"errors"
	"time"
)

// InvoiceService 发票服务
type InvoiceService struct {
	invoiceRepo *repository.InvoiceRepository
}

// NewInvoiceService 创建发票服务实例
func NewInvoiceService() *InvoiceService {
	return &InvoiceService{invoiceRepo: repository.NewInvoiceRepository()}
}

// CreateInvoice 创建发票
func (s *InvoiceService) CreateInvoice(tenantId, userId int64, invoiceDate time.Time, amount float64, invoiceNumber, vendor, taxId, category, paymentSource, status string) (*model.Invoice, error) {
	existingInvoice, err := s.invoiceRepo.FindByInvoiceNumber(invoiceNumber)
	if err != nil {
		return nil, err
	}
	if existingInvoice != nil {
		return nil, errors.New("该发票已存在, 请勿重复上传")
	}

	invoice := &model.Invoice{
		TenantID:      &tenantId,
		UserID:        userId,
		InvoiceNumber: &invoiceNumber,
		InvoiceDate:   &invoiceDate,
		Amount:        &amount,
		Vendor:        &vendor,
		TaxID:         &taxId,
		Category:      &category,
		PaymentSource: &paymentSource,
		Status:        status,
	}

	err = s.invoiceRepo.Create(invoice)
	if err != nil {
		return nil, err
	}

	return invoice, nil
}

// GetAllInvoicePage 分页获取所有发票
func (s *InvoiceService) GetAllInvoicePage(page, pageSize int) ([]*model.Invoice, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize
	limit := pageSize

	invoices, err := s.invoiceRepo.FindAllPage(limit, offset)
	if err != nil {
		return nil, err
	}

	return invoices, nil
}

// GetAllInvoicePageByUserId 通过用户id分页获取所有发票
func (s *InvoiceService) GetAllInvoicePageByUserId(page, pageSize int, userId int64) ([]*model.Invoice, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize
	limit := pageSize

	invoices, err := s.invoiceRepo.FindAllPageByUserId(limit, offset, userId)
	if err != nil {
		return nil, err
	}

	return invoices, nil
}

// GetAllInvoicePageByTenantId 通过租户id分页获取所有发票
func (s *InvoiceService) GetAllInvoicePageByTenantId(page, pageSize int, tenantId int64) ([]*model.Invoice, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize
	limit := pageSize

	invoices, err := s.invoiceRepo.FindAllPageByUserId(limit, offset, tenantId)
	if err != nil {
		return nil, err
	}

	return invoices, nil
}
