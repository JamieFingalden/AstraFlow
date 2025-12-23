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
func (s *InvoiceService) CreateInvoice(tenantId, userId int64, invoiceDate time.Time, amount float64, invoiceNumber, vendor, taxId, category, paymentSource, status string, imageURL, description *string) (*model.Invoice, error) {
	existingInvoice, err := s.invoiceRepo.FindByInvoiceNumber(invoiceNumber)
	if err != nil {
		return nil, err
	}
	if existingInvoice != nil {
		return nil, errors.New("该发票已存在, 请勿重复上传")
	}

	invoice := &model.Invoice{
		TenantID:       &tenantId,
		UserID:         userId,
		InvoiceNumber:  &invoiceNumber,
		InvoiceDate:    &invoiceDate,
		Amount:         &amount,
		Vendor:         vendor,
		TaxID:          &taxId,
		Category:       &category,
		PaymentSource:  &paymentSource,
		Status:         status,
		ImageURL:       imageURL,
		Description:    description,
		IdentifyStatus: "unnecessary", // 默认值
	}

	err = s.invoiceRepo.Create(invoice)
	if err != nil {
		return nil, err
	}

	return invoice, nil
}

// GetAllInvoicePage 分页获取所有发票
func (s *InvoiceService) GetAllInvoicePage(page, pageSize int) ([]*model.Invoice, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize
	limit := pageSize

	invoices, total, err := s.invoiceRepo.FindAllPage(limit, offset)
	if err != nil {
		return nil, 0, err
	}

	return invoices, total, nil
}

// GetAllInvoicePageByUserId 通过用户id分页获取所有发票
func (s *InvoiceService) GetAllInvoicePageByUserId(page, pageSize int, userId, tenantId int64) ([]*model.Invoice, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize
	limit := pageSize

	if tenantId != 0 {
		invoices, total, err := s.invoiceRepo.FindAllPageByTenantId(limit, offset, tenantId)
		if err != nil {
			return nil, 0, err
		}

		return invoices, total, nil
	}
	invoices, total, err := s.invoiceRepo.FindAllPageByUserId(limit, offset, userId)
	if err != nil {
		return nil, 0, err
	}

	return invoices, total, nil
}

// GetAllInvoicePageByTenantId 通过租户id分页获取所有发票
func (s *InvoiceService) GetAllInvoicePageByTenantId(page, pageSize int, tenantId int64) ([]*model.Invoice, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize
	limit := pageSize

	invoices, total, err := s.invoiceRepo.FindAllPageByTenantId(limit, offset, tenantId)
	if err != nil {
		return nil, 0, err
	}

	return invoices, total, nil
}

func (s *InvoiceService) UpdateInvoice(id int64, invoiceDate time.Time, amount float64, invoiceNumber, vendor, taxId, category, paymentSource, status string, imageURL, description *string) (*model.Invoice, error) {
	existingInvoice, err := s.invoiceRepo.FindById(id)
	if err != nil {
		return nil, err
	}

	if invoiceNumber != *existingInvoice.InvoiceNumber {
		_, err := s.invoiceRepo.FindByInvoiceNumber(invoiceNumber)
		if err != nil && err.Error() != "record not found" { // Only error if there's a different issue than record not found
			return nil, errors.New("该发票号已存在")
		}
	}

	existingInvoice.InvoiceNumber = &invoiceNumber
	existingInvoice.InvoiceDate = &invoiceDate
	existingInvoice.Amount = &amount
	existingInvoice.Vendor = vendor // Vendor is now a required string, not a pointer
	existingInvoice.TaxID = &taxId
	existingInvoice.Category = &category
	existingInvoice.PaymentSource = &paymentSource
	existingInvoice.Status = status
	existingInvoice.ImageURL = imageURL
	existingInvoice.Description = description

	err = s.invoiceRepo.Update(existingInvoice)
	if err != nil {
		return nil, err
	}

	return existingInvoice, nil
}

func (s *InvoiceService) DeleteInvoice(id int64) error {
	invoice, err := s.invoiceRepo.FindById(id)
	if err != nil {
		return err
	}

	if invoice == nil {
		return errors.New("发票不存在")
	}

	err = s.invoiceRepo.Delete(id)

	if err != nil {
		return err
	}

	return nil
}

func (s *InvoiceService) FindInvoiceInfoById(id int64) (*model.Invoice, error) {
	invoice, err := s.invoiceRepo.FindById(id)
	if err != nil {
		return nil, err
	}

	if invoice == nil {
		return nil, errors.New("发票不存在")
	}

	return invoice, nil
}
