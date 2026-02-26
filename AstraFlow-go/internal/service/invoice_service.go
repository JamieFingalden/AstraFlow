package service

import (
	"AstraFlow-go/internal/model"
	"AstraFlow-go/internal/repository"
	"errors"
	"strings"
	"time"

	"gorm.io/gorm"
)

var (
	ErrInvalidInvoiceID      = errors.New("invalid invoice id")
	ErrInvoiceNotFound       = errors.New("invoice not found")
	ErrInvoiceAlreadyHandled = errors.New("invoice already handled")
	ErrFinalAmountRequired   = errors.New("final amount required")
	ErrRejectRemarksRequired = errors.New("reject remarks required")
	ErrInvalidReviewAction   = errors.New("invalid review action")
)

// PendingInvoicesResult 待审核单据分页结果
type PendingInvoicesResult struct {
	Items      []repository.PendingInvoiceItem `json:"items"`
	Page       int                             `json:"page"`
	Size       int                             `json:"size"`
	Total      int64                           `json:"total"`
	TotalPages int64                           `json:"total_pages"`
}

// ReviewInvoiceResult 审核结果
type ReviewInvoiceResult struct {
	InvoiceID    int64               `json:"invoice_id"`
	Status       model.InvoiceStatus `json:"status"`
	Amount       float64             `json:"amount"`
	ReviewerID   *int64              `json:"reviewer_id,omitempty"`
	ReviewRemark string              `json:"review_remarks"`
}

// InvoiceService 发票服务
type InvoiceService struct {
	invoiceRepo *repository.InvoiceRepository
}

// NewInvoiceService 创建发票服务实例
func NewInvoiceService() *InvoiceService {
	return &InvoiceService{invoiceRepo: repository.NewInvoiceRepository()}
}

// CreateInvoice 创建发票
// Refactored to match new Invoice model
func (s *InvoiceService) CreateInvoice(tenantId, userId int64, attachmentID int64, invoiceDate time.Time, amount float64, invoiceNumber, vendor, category, description string) (*model.Invoice, error) {
	// Check for duplicate invoice number if provided
	if invoiceNumber != "" {
		existingInvoice, err := s.invoiceRepo.FindByInvoiceNumber(invoiceNumber)
		if err != nil {
			return nil, err
		}
		if existingInvoice != nil {
			return nil, errors.New("该发票已存在, 请勿重复上传")
		}
	}

	invoice := &model.Invoice{
		TenantID:      &tenantId,
		UserID:        userId,
		AttachmentID:  attachmentID,
		InvoiceNumber: invoiceNumber,
		InvoiceDate:   &invoiceDate,
		Amount:        amount,
		Vendor:        vendor,
		Category:      category,
		Description:   description,
		Status:        model.StatusPending,
	}

	err := s.invoiceRepo.Create(invoice)
	if err != nil {
		return nil, err
	}

	return invoice, nil
}

// CreateOCRInvoice 创建 AI OCR 发票记录（待识别状态）
func (s *InvoiceService) CreateOCRInvoice(tenantId *int64, userId int64, attachmentID int64) (*model.Invoice, error) {
	invoice := &model.Invoice{
		TenantID:     tenantId,
		UserID:       userId,
		AttachmentID: attachmentID,
		Status:       model.StatusRecognizing,
		Source:       model.SourceOCR,
		// InvoiceDate is omitted here, so it will be NULL in the DB
	}

	err := s.invoiceRepo.Create(invoice)
	if err != nil {
		return nil, err
	}

	return invoice, nil
}

// CreateManualInvoice 创建手动填报的发票记录（待发布草稿状态）
func (s *InvoiceService) CreateManualInvoice(tenantId *int64, userId int64, attachmentID int64, amount float64, invoiceDate time.Time, category, description string) (*model.Invoice, error) {
	invoice := &model.Invoice{
		TenantID:     tenantId,
		UserID:       userId,
		AttachmentID: attachmentID,
		Amount:       amount,
		InvoiceDate:  &invoiceDate,
		Category:     category,
		Description:  description,
		Status:       model.StatusDraft,
		Source:       model.SourceManual,
	}

	err := s.invoiceRepo.Create(invoice)
	if err != nil {
		return nil, err
	}

	return invoice, nil
}

// CreateEmptyInvoice 创建一个初始的空发票记录
func (s *InvoiceService) CreateEmptyInvoice(tenantId *int64, userId int64, attachmentID int64) (*model.Invoice, error) {
	invoice := &model.Invoice{
		TenantID:     tenantId,
		UserID:       userId,
		AttachmentID: attachmentID,
		Status:       model.StatusPending,
		// 其他字段留空或默认值
	}

	err := s.invoiceRepo.Create(invoice)
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

// GetInvoicesByUserIDAndStatus 根据用户ID和状态分页获取发票
func (s *InvoiceService) GetInvoicesByUserIDAndStatus(page, pageSize int, userId int64, status string) ([]*model.Invoice, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	offset := (page - 1) * pageSize
	return s.invoiceRepo.FindAllPageByUserIdAndStatus(pageSize, offset, userId, status)
}

func (s *InvoiceService) UpdateInvoice(id int64, invoiceDate time.Time, amount float64, invoiceNumber, vendor, category, description, status string) (*model.Invoice, error) {
	existingInvoice, err := s.invoiceRepo.FindById(id)
	if err != nil {
		return nil, err
	}

	if invoiceNumber != "" && invoiceNumber != existingInvoice.InvoiceNumber {
		found, err := s.invoiceRepo.FindByInvoiceNumber(invoiceNumber)
		if err != nil && err.Error() != "record not found" {
			return nil, err
		}
		if found != nil && found.ID != id {
			return nil, errors.New("该发票号已存在")
		}
	}

	existingInvoice.InvoiceNumber = invoiceNumber
	if !invoiceDate.IsZero() {
		existingInvoice.InvoiceDate = &invoiceDate
	}
	existingInvoice.Amount = amount
	existingInvoice.Vendor = vendor
	existingInvoice.Category = category
	existingInvoice.Description = description

	if status != "" {
		existingInvoice.Status = model.InvoiceStatus(status)
	}

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

func (s *InvoiceService) ConfirmInvoice(id int64, invoiceDate time.Time, amount float64, invoiceNumber, vendor, category, description string) (*model.Invoice, error) {
	existingInvoice, err := s.invoiceRepo.FindById(id)
	if err != nil {
		return nil, err
	}
	if existingInvoice == nil {
		return nil, errors.New("发票不存在")
	}

	if invoiceNumber != "" && invoiceNumber != existingInvoice.InvoiceNumber {
		found, err := s.invoiceRepo.FindByInvoiceNumber(invoiceNumber)
		if err != nil && err.Error() != "record not found" {
			return nil, err
		}
		if found != nil && found.ID != id {
			return nil, errors.New("该发票号已存在")
		}
	}

	existingInvoice.InvoiceNumber = invoiceNumber
	if !invoiceDate.IsZero() {
		existingInvoice.InvoiceDate = &invoiceDate
	}
	existingInvoice.Amount = amount
	existingInvoice.Vendor = vendor
	existingInvoice.Category = category
	existingInvoice.Description = description

	// Update status from unconfirmed to draft
	if existingInvoice.Status == model.StatusUnconfirmed {
		existingInvoice.Status = model.StatusDraft
	}

	err = s.invoiceRepo.Update(existingInvoice)
	if err != nil {
		return nil, err
	}

	return existingInvoice, nil
}

func (s *InvoiceService) PublishInvoices(ids []int64) error {
	for _, id := range ids {
		invoice, err := s.invoiceRepo.FindById(id)
		if err != nil || invoice == nil {
			continue
		}
		if invoice.Status == model.StatusDraft {
			invoice.Status = model.StatusPending
			_ = s.invoiceRepo.Update(invoice)
		}
	}
	return nil
}

// GetPendingInvoices 分页查询待审核任务池
func (s *InvoiceService) GetPendingInvoices(page, size int) (*PendingInvoicesResult, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}
	if size > 100 {
		size = 100
	}

	offset := (page - 1) * size

	total, err := s.invoiceRepo.CountPendingInvoices()
	if err != nil {
		return nil, err
	}

	items, err := s.invoiceRepo.FindPendingInvoices(size, offset)
	if err != nil {
		return nil, err
	}

	totalPages := int64(0)
	if size > 0 {
		totalPages = (total + int64(size) - 1) / int64(size)
	}

	return &PendingInvoicesResult{
		Items:      items,
		Page:       page,
		Size:       size,
		Total:      total,
		TotalPages: totalPages,
	}, nil
}

// GetInvoiceDetailForAudit 获取审核详情（只读，状态不限）
func (s *InvoiceService) GetInvoiceDetailForAudit(id int64) (*repository.InvoiceDetail, error) {
	if id <= 0 {
		return nil, ErrInvalidInvoiceID
	}

	detail, err := s.invoiceRepo.FindInvoiceDetailByID(id)
	if err != nil {
		return nil, err
	}
	if detail == nil {
		return nil, ErrInvoiceNotFound
	}

	return detail, nil
}

// ReviewInvoice 审核批复
func (s *InvoiceService) ReviewInvoice(id, reviewerID int64, action string, finalAmount *float64, remarks string) (*ReviewInvoiceResult, error) {
	if id <= 0 {
		return nil, errors.New("单据ID不合法")
	}
	if reviewerID <= 0 {
		return nil, errors.New("审核人信息无效")
	}

	action = strings.TrimSpace(action)
	remarks = strings.TrimSpace(remarks)

	invoice, err := s.invoiceRepo.ReviewInvoice(id, reviewerID, action, finalAmount, remarks)
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return nil, ErrInvoiceNotFound
		case errors.Is(err, repository.ErrInvoiceAlreadyProcessed):
			return nil, ErrInvoiceAlreadyHandled
		case errors.Is(err, repository.ErrFinalAmountRequired):
			return nil, ErrFinalAmountRequired
		case errors.Is(err, repository.ErrRejectRemarksRequired):
			return nil, ErrRejectRemarksRequired
		case errors.Is(err, repository.ErrInvalidReviewAction):
			return nil, ErrInvalidReviewAction
		default:
			return nil, err
		}
	}

	if invoice == nil {
		return nil, ErrInvoiceNotFound
	}

	return &ReviewInvoiceResult{
		InvoiceID:    invoice.ID,
		Status:       invoice.Status,
		Amount:       invoice.Amount,
		ReviewerID:   invoice.ReviewerID,
		ReviewRemark: invoice.ReviewRemarks,
	}, nil
}
