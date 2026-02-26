package repository

import (
	"AstraFlow-go/internal/database"
	"AstraFlow-go/internal/model"
	"errors"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var (
	ErrInvoiceAlreadyProcessed = errors.New("invoice already processed")
	ErrFinalAmountRequired     = errors.New("final amount required")
	ErrRejectRemarksRequired   = errors.New("reject remarks required")
	ErrInvalidReviewAction     = errors.New("invalid review action")
)

// PendingInvoiceItem 待审核单据列表项（仓储层查询结果）
type PendingInvoiceItem struct {
	ID            int64               `json:"id"`
	TenantID      *int64              `json:"tenant_id,omitempty"`
	UserID        int64               `json:"user_id"`
	UserName      string              `json:"user_name"`
	AttachmentID  int64               `json:"attachment_id"`
	InvoiceNumber string              `json:"invoice_number"`
	InvoiceDate   *string             `json:"invoice_date,omitempty"`
	Amount        float64             `json:"amount"`
	Vendor        string              `json:"vendor"`
	Category      string              `json:"category"`
	Description   string              `json:"description"`
	Status        model.InvoiceStatus `json:"status"`
	CreatedAt     string              `json:"created_at"`
	UpdatedAt     string              `json:"updated_at"`
}

// InvoiceDetail 单据详情（仓储层查询结果）
type InvoiceDetail struct {
	ID            int64               `json:"id"`
	TenantID      *int64              `json:"tenant_id,omitempty"`
	UserID        int64               `json:"user_id"`
	UserName      string              `json:"user_name"`
	UserUsername  string              `json:"user_username"`
	AttachmentID  int64               `json:"attachment_id"`
	ImageURL      string              `json:"image_url"`
	InvoiceNumber string              `json:"invoice_number"`
	InvoiceDate   *string             `json:"invoice_date,omitempty"`
	Amount        float64             `json:"amount"`
	Vendor        string              `json:"vendor"`
	Category      string              `json:"category"`
	Description   string              `json:"description"`
	Status        model.InvoiceStatus `json:"status"`
	ReviewerID    *int64              `json:"reviewer_id,omitempty"`
	ReviewRemarks string              `json:"review_remarks"`
	PaidAt        *string             `json:"paid_at,omitempty"`
	CreatedAt     string              `json:"created_at"`
	UpdatedAt     string              `json:"updated_at"`
}

// ArchiveSearchParams 历史归档搜索参数
type ArchiveSearchParams struct {
	Page      int
	Size      int
	Status    string
	StartDate *time.Time
	EndDate   *time.Time
	Keyword   string
	UserID    *int64
	TenantID  *int64
}

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
func (r *InvoiceRepository) FindAllPage(limit, offset int) ([]*model.Invoice, int64, error) {
	var invoices []*model.Invoice
	var total int64

	err := r.db.Model(&model.Invoice{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.Limit(limit).Offset(offset).Find(&invoices).Error
	return invoices, total, err
}

// FindAllPageByUserId 根据用户id分页查询所有发票
func (r *InvoiceRepository) FindAllPageByUserId(limit, offset int, userId int64) ([]*model.Invoice, int64, error) {
	var invoices []*model.Invoice
	var total int64
	var err error

	err = r.db.Where("user_id = ?", userId).Limit(limit).Offset(offset).Find(&invoices).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.Model(&model.Invoice{}).Where("user_id = ?", userId).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	return invoices, total, err
}

// FindAllPageByTenantId 根据租户id分页查询所有发票
func (r *InvoiceRepository) FindAllPageByTenantId(limit, offset int, tenantId int64) ([]*model.Invoice, int64, error) {
	var invoices []*model.Invoice
	var total int64

	err := r.db.Model(&model.Invoice{}).Where("tenant_id = ?", tenantId).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.Where("tenant_id = ?", tenantId).Limit(limit).Offset(offset).Find(&invoices).Error
	return invoices, total, err
}

// FindAllPageByUserIdAndStatus 根据用户ID和状态分页查询发票
func (r *InvoiceRepository) FindAllPageByUserIdAndStatus(limit, offset int, userId int64, status string) ([]*model.Invoice, int64, error) {
	var invoices []*model.Invoice
	var total int64

	// 构建基础查询
	query := r.db.Model(&model.Invoice{}).Where("user_id = ?", userId)

	// 根据状态进行筛选
	if status != "" {
		if status == "submitted" {
			// "submitted" 是一个前端概念，代表所有已进入正式流程的状态
			submittedStatuses := []string{
				string(model.StatusPending),
				string(model.StatusApproved),
				string(model.StatusRejected),
				string(model.StatusPaid),
			}
			query = query.Where("status IN ?", submittedStatuses)
		} else {
			// 查询特定状态
			query = query.Where("status = ?", status)
		}
	}

	// 计算总数
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 获取分页数据
	err = query.Limit(limit).Offset(offset).Order("created_at DESC").Find(&invoices).Error
	if err != nil {
		return nil, 0, err
	}

	return invoices, total, err
}

func (r *InvoiceRepository) Update(invoice *model.Invoice) error {
	return r.db.Save(invoice).Error
}

func (r *InvoiceRepository) Delete(id int64) error {
	return r.db.Delete(&model.Invoice{}, id).Error
}

func (r *InvoiceRepository) FindById(id int64) (*model.Invoice, error) {
	var invoice model.Invoice
	err := r.db.Where("id = ?", id).Find(&invoice).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &invoice, err
}

// BeginTx 开启数据库事务
func (r *InvoiceRepository) BeginTx() *gorm.DB {
	return r.db.Begin()
}

// FindApprovedInvoices 分页查询待打款单据，预加载提交人信息并返回总数
func (r *InvoiceRepository) FindApprovedInvoices(limit, offset int) ([]*model.Invoice, int64, error) {
	items := make([]*model.Invoice, 0)
	var total int64

	err := r.db.Model(&model.Invoice{}).
		Where("status = ?", model.StatusApproved).
		Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.Model(&model.Invoice{}).
		Preload("User").
		Where("status = ?", model.StatusApproved).
		Order("updated_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&items).Error
	if err != nil {
		return nil, 0, err
	}

	return items, total, nil
}

// BatchPayInvoices 在事务中批量打款：approved -> paid
func (r *InvoiceRepository) BatchPayInvoices(tx *gorm.DB, invoiceIDs []int64) (int64, error) {
	result := tx.Model(&model.Invoice{}).
		Where("id IN ? AND status = ?", invoiceIDs, model.StatusApproved).
		Updates(map[string]interface{}{
			"status":  model.StatusPaid,
			"paid_at": time.Now(),
		})

	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, nil
}

// SearchArchiveInvoices 按条件搜索归档单据并返回分页结果
func (r *InvoiceRepository) SearchArchiveInvoices(params ArchiveSearchParams) ([]*model.Invoice, int64, error) {
	items := make([]*model.Invoice, 0)
	var total int64

	query := r.db.Model(&model.Invoice{}).Preload("User")

	if params.TenantID != nil && *params.TenantID != 0 {
		query = query.Where("tenant_id = ?", *params.TenantID)
	} else {
		query = query.Where("tenant_id IS NULL OR tenant_id = 0")
	}

	if params.UserID != nil {
		query = query.Where("user_id = ?", *params.UserID)
	}

	if params.Status != "" {
		query = query.Where("status = ?", params.Status)
	}

	if params.StartDate != nil && params.EndDate != nil {
		query = query.Where("invoice_date BETWEEN ? AND ?", *params.StartDate, *params.EndDate)
	} else if params.StartDate != nil {
		query = query.Where("invoice_date >= ?", *params.StartDate)
	} else if params.EndDate != nil {
		query = query.Where("invoice_date <= ?", *params.EndDate)
	}

	if params.Keyword != "" {
		query = query.Where("description LIKE ? OR category LIKE ?", "%"+params.Keyword+"%", "%"+params.Keyword+"%")
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (params.Page - 1) * params.Size
	err = query.Order("created_at DESC").Offset(offset).Limit(params.Size).Find(&items).Error
	if err != nil {
		return nil, 0, err
	}

	return items, total, nil
}

// CountByStatus 按状态统计单据数量
func (r *InvoiceRepository) CountByStatus(tenantID *int64, userID *int64, status model.InvoiceStatus) (int64, error) {
	var total int64
	query := r.db.Model(&model.Invoice{}).Where("status = ?", status)

	if tenantID != nil && *tenantID != 0 {
		query = query.Where("tenant_id = ?", *tenantID)
	} else {
		query = query.Where("tenant_id IS NULL OR tenant_id = 0")
	}

	if userID != nil {
		query = query.Where("user_id = ?", *userID)
	}

	err := query.Count(&total).Error
	if err != nil {
		return 0, err
	}

	return total, nil
}

// SumAmountByStatuses 按状态集合汇总金额
func (r *InvoiceRepository) SumAmountByStatuses(tenantID *int64, userID *int64, statuses []model.InvoiceStatus) (float64, error) {
	var sum float64
	query := r.db.Model(&model.Invoice{}).Select("COALESCE(SUM(amount), 0)").Where("status IN ?", statuses)

	if tenantID != nil && *tenantID != 0 {
		query = query.Where("tenant_id = ?", *tenantID)
	} else {
		query = query.Where("tenant_id IS NULL OR tenant_id = 0")
	}

	if userID != nil {
		query = query.Where("user_id = ?", *userID)
	}

	err := query.Scan(&sum).Error
	if err != nil {
		return 0, err
	}

	return sum, nil
}

// SumAmountByStatusAndUpdatedRange 按状态与更新时间范围汇总金额
func (r *InvoiceRepository) SumAmountByStatusAndUpdatedRange(tenantID *int64, userID *int64, status model.InvoiceStatus, startTime, endTime time.Time) (float64, error) {
	var sum float64
	query := r.db.Model(&model.Invoice{}).
		Select("COALESCE(SUM(amount), 0)").
		Where("status = ?", status).
		Where("updated_at BETWEEN ? AND ?", startTime, endTime)

	if tenantID != nil && *tenantID != 0 {
		query = query.Where("tenant_id = ?", *tenantID)
	} else {
		query = query.Where("tenant_id IS NULL OR tenant_id = 0")
	}

	if userID != nil {
		query = query.Where("user_id = ?", *userID)
	}

	err := query.Scan(&sum).Error
	if err != nil {
		return 0, err
	}

	return sum, nil
}

// CountPendingInvoices 统计待审核单据总数
func (r *InvoiceRepository) CountPendingInvoices() (int64, error) {
	var total int64
	err := r.db.Model(&model.Invoice{}).Where("status = ?", model.StatusPending).Count(&total).Error
	if err != nil {
		return 0, err
	}
	return total, nil
}

// FindPendingInvoices 分页查询待审核单据，包含提交人姓名
func (r *InvoiceRepository) FindPendingInvoices(limit, offset int) ([]PendingInvoiceItem, error) {
	items := make([]PendingInvoiceItem, 0)
	err := r.db.Table("invoices AS i").
		Select(`
			i.id,
			i.tenant_id,
			i.user_id,
			COALESCE(NULLIF(u.name, ''), u.username) AS user_name,
			i.attachment_id,
			i.invoice_number,
			DATE_FORMAT(i.invoice_date, '%Y-%m-%d') AS invoice_date,
			i.amount,
			i.vendor,
			i.category,
			i.description,
			i.status,
			DATE_FORMAT(i.created_at, '%Y-%m-%d %H:%i:%s') AS created_at,
			DATE_FORMAT(i.updated_at, '%Y-%m-%d %H:%i:%s') AS updated_at
		`).
		Joins("LEFT JOIN users AS u ON u.id = i.user_id").
		Where("i.status = ?", model.StatusPending).
		Order("i.created_at DESC").
		Offset(offset).
		Limit(limit).
		Scan(&items).Error
	if err != nil {
		return nil, err
	}
	return items, nil
}

// FindInvoiceDetailByID 查询单据详情（状态不限）
func (r *InvoiceRepository) FindInvoiceDetailByID(id int64) (*InvoiceDetail, error) {
	var detail InvoiceDetail
	err := r.db.Table("invoices AS i").
		Select(`
			i.id,
			i.tenant_id,
			i.user_id,
			COALESCE(NULLIF(u.name, ''), u.username) AS user_name,
			u.username AS user_username,
			i.attachment_id,
			a.file_url AS image_url,
			i.invoice_number,
			DATE_FORMAT(i.invoice_date, '%Y-%m-%d') AS invoice_date,
			i.amount,
			i.vendor,
			i.category,
			i.description,
			i.status,
			i.reviewer_id,
			i.review_remarks,
			DATE_FORMAT(i.paid_at, '%Y-%m-%d %H:%i:%s') AS paid_at,
			DATE_FORMAT(i.created_at, '%Y-%m-%d %H:%i:%s') AS created_at,
			DATE_FORMAT(i.updated_at, '%Y-%m-%d %H:%i:%s') AS updated_at
		`).
		Joins("LEFT JOIN users AS u ON u.id = i.user_id").
		Joins("LEFT JOIN attachments AS a ON a.id = i.attachment_id").
		Where("i.id = ?", id).
		First(&detail).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &detail, nil
}

// ReviewInvoice 审核单据（事务 + 行锁）
func (r *InvoiceRepository) ReviewInvoice(id, reviewerID int64, action string, finalAmount *float64, remarks string) (*model.Invoice, error) {
	var reviewed *model.Invoice
	err := r.db.Transaction(func(tx *gorm.DB) error {
		var invoice model.Invoice
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&invoice, id).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return err
			}
			return err
		}

		if invoice.Status != model.StatusPending {
			return ErrInvoiceAlreadyProcessed
		}

		updates := map[string]interface{}{
			"reviewer_id":    reviewerID,
			"review_remarks": remarks,
		}

		switch action {
		case "approve":
			if finalAmount == nil {
				return ErrFinalAmountRequired
			}
			updates["status"] = model.StatusApproved
			updates["amount"] = *finalAmount
		case "reject":
			if remarks == "" {
				return ErrRejectRemarksRequired
			}
			updates["status"] = model.StatusRejected
		default:
			return ErrInvalidReviewAction
		}

		if err := tx.Model(&model.Invoice{}).Where("id = ?", id).Updates(updates).Error; err != nil {
			return err
		}

		if err := tx.First(&invoice, id).Error; err != nil {
			return err
		}

		reviewed = &invoice
		return nil
	})
	if err != nil {
		return nil, err
	}

	return reviewed, nil
}
