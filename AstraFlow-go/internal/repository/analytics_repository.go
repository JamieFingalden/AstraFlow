package repository

import (
	"AstraFlow-go/internal/database"
	"AstraFlow-go/internal/model"
	"fmt"
	"time"

	"gorm.io/gorm"
)

// AnalyticsRepository handles analytics data operations
type AnalyticsRepository struct {
	db *gorm.DB
}

// NewAnalyticsRepository creates a new analytics repository instance
func NewAnalyticsRepository() *AnalyticsRepository {
	return &AnalyticsRepository{
		db: database.DB,
	}
}

// GetDashboardMetrics retrieves key dashboard metrics
func (r *AnalyticsRepository) GetDashboardMetrics(tenantID *int64, userID int64, startDate, endDate *time.Time) (*model.DashboardMetrics, error) {
	metrics := &model.DashboardMetrics{}

	// Calculate monthly expense
	var monthlyExpense float64
	query := r.db.Model(&model.Invoice{}).Select("COALESCE(SUM(amount), 0)")

	if tenantID != nil && *tenantID != 0 {
		query = query.Where("tenant_id = ?", *tenantID)
	} else {
		query = query.Where("tenant_id IS NULL OR tenant_id = 0")
	}

	if startDate != nil && endDate != nil {
		query = query.Where("created_at BETWEEN ? AND ?", *startDate, *endDate)
	} else if startDate != nil {
		query = query.Where("created_at >= ?", *startDate)
	} else if endDate != nil {
		query = query.Where("created_at <= ?", *endDate)
	}

	err := query.Where("user_id = ?", userID).Scan(&monthlyExpense).Error
	if err != nil {
		return nil, err
	}
	metrics.MonthlyExpense = monthlyExpense

	// Count auto-processed invoices (assuming those with recognized status)
	var autoProcessed int64
	autoQuery := r.db.Model(&model.Invoice{}).Where("user_id = ?", userID)

	if tenantID != nil && *tenantID != 0 {
		autoQuery = autoQuery.Where("tenant_id = ?", *tenantID)
	} else {
		autoQuery = autoQuery.Where("tenant_id IS NULL OR tenant_id = 0")
	}

	if startDate != nil && endDate != nil {
		autoQuery = autoQuery.Where("created_at BETWEEN ? AND ?", *startDate, *endDate)
	} else if startDate != nil {
		autoQuery = autoQuery.Where("created_at >= ?", *startDate)
	} else if endDate != nil {
		autoQuery = autoQuery.Where("created_at <= ?", *endDate)
	}

	autoQuery = autoQuery.Where("status IN ?", []string{"recognized", "confirmed"})
	err = autoQuery.Count(&autoProcessed).Error
	if err != nil {
		return nil, err
	}
	metrics.AutoProcessed = autoProcessed

	// Calculate AI accuracy (placeholder - using 98.7% as default)
	metrics.AIAccuracy = 98.7

	// Count pending reviews
	var pendingReviews int64
	pendingQuery := r.db.Model(&model.Invoice{}).Where("user_id = ? AND status = ?", userID, "pending")

	if tenantID != nil && *tenantID != 0 {
		pendingQuery = pendingQuery.Where("tenant_id = ?", *tenantID)
	} else {
		pendingQuery = pendingQuery.Where("tenant_id IS NULL OR tenant_id = 0")
	}

	if startDate != nil && endDate != nil {
		pendingQuery = pendingQuery.Where("created_at BETWEEN ? AND ?", *startDate, *endDate)
	} else if startDate != nil {
		pendingQuery = pendingQuery.Where("created_at >= ?", *startDate)
	} else if endDate != nil {
		pendingQuery = pendingQuery.Where("created_at <= ?", *endDate)
	}

	err = pendingQuery.Count(&pendingReviews).Error
	if err != nil {
		return nil, err
	}
	metrics.PendingReviews = pendingReviews

	return metrics, nil
}

// GetExpenseCategories retrieves expense breakdown by category
func (r *AnalyticsRepository) GetExpenseCategories(tenantID *int64, userID int64, startDate, endDate *time.Time) ([]model.ExpenseCategory, error) {
	var results []struct {
		Category string
		Total    float64
	}

	query := r.db.Model(&model.Invoice{}).
		Select("COALESCE(category, '其他') as category, COALESCE(SUM(amount), 0) as total").
		Where("user_id = ? AND amount IS NOT NULL", userID)

	if tenantID != nil && *tenantID != 0 {
		query = query.Where("tenant_id = ?", *tenantID)
	} else {
		query = query.Where("tenant_id IS NULL OR tenant_id = 0")
	}

	if startDate != nil && endDate != nil {
		query = query.Where("created_at BETWEEN ? AND ?", *startDate, *endDate)
	} else if startDate != nil {
		query = query.Where("created_at >= ?", *startDate)
	} else if endDate != nil {
		query = query.Where("created_at <= ?", *endDate)
	}

	query = query.Group("COALESCE(category, '其他')").Order("total DESC")

	err := query.Scan(&results).Error
	if err != nil {
		return nil, err
	}

	// Define default colors for categories
	colors := []string{"#3B82F6", "#10B981", "#8B5CF6", "#EF4444", "#F59E0B", "#EC4899", "#06B6D4", "#84CC16"}

	categories := make([]model.ExpenseCategory, len(results))
	for i, result := range results {
		colorIndex := i % len(colors)
		categories[i] = model.ExpenseCategory{
			Name:  result.Category,
			Value: result.Total,
			Color: colors[colorIndex],
		}
	}

	return categories, nil
}

// GetWeeklyExpenses retrieves weekly expense trends
func (r *AnalyticsRepository) GetWeeklyExpenses(tenantID *int64, userID int64, startDate, endDate *time.Time) ([]model.WeeklyExpense, error) {
	var results []struct {
		Week    string
		Expense float64
	}

	query := r.db.Model(&model.Invoice{}).
		Select("DATE_FORMAT(created_at, '%Y-W%u') as week, COALESCE(SUM(amount), 0) as expense").
		Where("user_id = ? AND amount IS NOT NULL", userID).
		Group("DATE_FORMAT(created_at, '%Y-W%u')").
		Order("week")

	if tenantID != nil && *tenantID != 0 {
		query = query.Where("tenant_id = ?", *tenantID)
	} else {
		query = query.Where("tenant_id IS NULL OR tenant_id = 0")
	}

	if startDate != nil && endDate != nil {
		query = query.Where("created_at BETWEEN ? AND ?", *startDate, *endDate)
	} else if startDate != nil {
		query = query.Where("created_at >= ?", *startDate)
	} else if endDate != nil {
		query = query.Where("created_at <= ?", *endDate)
	}

	err := query.Scan(&results).Error
	if err != nil {
		return nil, err
	}

	weeklyExpenses := make([]model.WeeklyExpense, len(results))
	for i, result := range results {
		// Format the week string to be more readable
		week := fmt.Sprintf("第%s周", result.Week[len(result.Week)-2:])
		weeklyExpenses[i] = model.WeeklyExpense{
			Week:    week,
			Expense: result.Expense,
		}
	}

	return weeklyExpenses, nil
}

// GetRecentBills retrieves recent bills for the dashboard
func (r *AnalyticsRepository) GetRecentBills(tenantID *int64, userID int64, limit int, startDate, endDate *time.Time) ([]model.RecentBill, error) {
	var bills []model.RecentBill

	query := r.db.Model(&model.Invoice{}).
		Select("id, created_at as date, vendor, COALESCE(category, '其他') as category, COALESCE(amount, 0) as amount, status").
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Limit(limit)

	if tenantID != nil && *tenantID != 0 {
		query = query.Where("tenant_id = ?", *tenantID)
	} else {
		query = query.Where("tenant_id IS NULL OR tenant_id = 0")
	}

	if startDate != nil && endDate != nil {
		query = query.Where("created_at BETWEEN ? AND ?", *startDate, *endDate)
	} else if startDate != nil {
		query = query.Where("created_at >= ?", *startDate)
	} else if endDate != nil {
		query = query.Where("created_at <= ?", *endDate)
	}

	err := query.Find(&bills).Error
	if err != nil {
		return nil, err
	}

	return bills, nil
}

// GetAllInvoicesForAIInsights retrieves all invoices for AI insights analysis
func (r *AnalyticsRepository) GetAllInvoicesForAIInsights(tenantID *int64, userID int64, startDate, endDate *time.Time) ([]model.Invoice, error) {
	var invoices []model.Invoice

	query := r.db.Model(&model.Invoice{}).
		Where("user_id = ?", userID).
		Order("created_at DESC")

	if tenantID != nil && *tenantID != 0 {
		query = query.Where("tenant_id = ?", *tenantID)
	} else {
		query = query.Where("tenant_id IS NULL OR tenant_id = 0")
	}

	if startDate != nil && endDate != nil {
		query = query.Where("created_at BETWEEN ? AND ?", *startDate, *endDate)
	} else if startDate != nil {
		query = query.Where("created_at >= ?", *startDate)
	} else if endDate != nil {
		query = query.Where("created_at <= ?", *endDate)
	}

	err := query.Find(&invoices).Error
	if err != nil {
		return nil, err
	}

	return invoices, nil
}

// GetReimbursementStatistics retrieves key reimbursement statistics
func (r *AnalyticsRepository) GetReimbursementStatistics(tenantID *int64, userID int64, startDate, endDate *time.Time) (*model.ReimbursementStatistics, error) {
	stats := &model.ReimbursementStatistics{}

	// Calculate total amount
	var totalAmount float64
	totalQuery := r.db.Model(&model.Invoice{}).Select("COALESCE(SUM(amount), 0)")

	if tenantID != nil && *tenantID != 0 {
		totalQuery = totalQuery.Where("tenant_id = ?", *tenantID)
	} else {
		totalQuery = totalQuery.Where("tenant_id IS NULL OR tenant_id = 0")
	}

	if startDate != nil && endDate != nil {
		totalQuery = totalQuery.Where("created_at BETWEEN ? AND ?", *startDate, *endDate)
	} else if startDate != nil {
		totalQuery = totalQuery.Where("created_at >= ?", *startDate)
	} else if endDate != nil {
		totalQuery = totalQuery.Where("created_at <= ?", *endDate)
	}

	err := totalQuery.Where("user_id = ?", userID).Scan(&totalAmount).Error
	if err != nil {
		return nil, err
	}

	// Calculate approved amount
	var approvedAmount float64
	approvedQuery := r.db.Model(&model.Invoice{}).Select("COALESCE(SUM(amount), 0)")

	if tenantID != nil && *tenantID != 0 {
		approvedQuery = approvedQuery.Where("tenant_id = ?", *tenantID)
	} else {
		approvedQuery = approvedQuery.Where("tenant_id IS NULL OR tenant_id = 0")
	}

	if startDate != nil && endDate != nil {
		approvedQuery = approvedQuery.Where("created_at BETWEEN ? AND ?", *startDate, *endDate)
	} else if startDate != nil {
		approvedQuery = approvedQuery.Where("created_at >= ?", *startDate)
	} else if endDate != nil {
		approvedQuery = approvedQuery.Where("created_at <= ?", *endDate)
	}

	err = approvedQuery.Where("user_id = ? AND status IN ?", userID, []string{"confirmed", "approved"}).Scan(&approvedAmount).Error
	if err != nil {
		return nil, err
	}

	// Calculate pending amount
	var pendingAmount float64
	pendingQuery := r.db.Model(&model.Invoice{}).Select("COALESCE(SUM(amount), 0)")

	if tenantID != nil && *tenantID != 0 {
		pendingQuery = pendingQuery.Where("tenant_id = ?", *tenantID)
	} else {
		pendingQuery = pendingQuery.Where("tenant_id IS NULL OR tenant_id = 0")
	}

	if startDate != nil && endDate != nil {
		pendingQuery = pendingQuery.Where("created_at BETWEEN ? AND ?", *startDate, *endDate)
	} else if startDate != nil {
		pendingQuery = pendingQuery.Where("created_at >= ?", *startDate)
	} else if endDate != nil {
		pendingQuery = pendingQuery.Where("created_at <= ?", *endDate)
	}

	err = pendingQuery.Where("user_id = ? AND status = ?", userID, "pending").Scan(&pendingAmount).Error
	if err != nil {
		return nil, err
	}

	// Count total requests
	var totalRequests int64
	totalReqQuery := r.db.Model(&model.Invoice{}).Where("user_id = ?", userID)

	if tenantID != nil && *tenantID != 0 {
		totalReqQuery = totalReqQuery.Where("tenant_id = ?", *tenantID)
	} else {
		totalReqQuery = totalReqQuery.Where("tenant_id IS NULL OR tenant_id = 0")
	}

	if startDate != nil && endDate != nil {
		totalReqQuery = totalReqQuery.Where("created_at BETWEEN ? AND ?", *startDate, *endDate)
	} else if startDate != nil {
		totalReqQuery = totalReqQuery.Where("created_at >= ?", *startDate)
	} else if endDate != nil {
		totalReqQuery = totalReqQuery.Where("created_at <= ?", *endDate)
	}

	err = totalReqQuery.Count(&totalRequests).Error
	if err != nil {
		return nil, err
	}

	// Count approved requests
	var approvedCount int64
	approvedReqQuery := r.db.Model(&model.Invoice{}).Where("user_id = ? AND status IN ?", userID, []string{"confirmed", "approved"})

	if tenantID != nil && *tenantID != 0 {
		approvedReqQuery = approvedReqQuery.Where("tenant_id = ?", *tenantID)
	} else {
		approvedReqQuery = approvedReqQuery.Where("tenant_id IS NULL OR tenant_id = 0")
	}

	if startDate != nil && endDate != nil {
		approvedReqQuery = approvedReqQuery.Where("created_at BETWEEN ? AND ?", *startDate, *endDate)
	} else if startDate != nil {
		approvedReqQuery = approvedReqQuery.Where("created_at >= ?", *startDate)
	} else if endDate != nil {
		approvedReqQuery = approvedReqQuery.Where("created_at <= ?", *endDate)
	}

	err = approvedReqQuery.Count(&approvedCount).Error
	if err != nil {
		return nil, err
	}

	// Count pending requests
	var pendingCount int64
	pendingReqQuery := r.db.Model(&model.Invoice{}).Where("user_id = ? AND status = ?", userID, "pending")

	if tenantID != nil && *tenantID != 0 {
		pendingReqQuery = pendingReqQuery.Where("tenant_id = ?", *tenantID)
	} else {
		pendingReqQuery = pendingReqQuery.Where("tenant_id IS NULL OR tenant_id = 0")
	}

	if startDate != nil && endDate != nil {
		pendingReqQuery = pendingReqQuery.Where("created_at BETWEEN ? AND ?", *startDate, *endDate)
	} else if startDate != nil {
		pendingReqQuery = pendingReqQuery.Where("created_at >= ?", *startDate)
	} else if endDate != nil {
		pendingReqQuery = pendingReqQuery.Where("created_at <= ?", *endDate)
	}

	err = pendingReqQuery.Count(&pendingCount).Error
	if err != nil {
		return nil, err
	}

	// Count rejected requests
	var rejectedCount int64
	rejectedReqQuery := r.db.Model(&model.Invoice{}).Where("user_id = ? AND status IN ?", userID, []string{"rejected", "failed"})

	if tenantID != nil && *tenantID != 0 {
		rejectedReqQuery = rejectedReqQuery.Where("tenant_id = ?", *tenantID)
	} else {
		rejectedReqQuery = rejectedReqQuery.Where("tenant_id IS NULL OR tenant_id = 0")
	}

	if startDate != nil && endDate != nil {
		rejectedReqQuery = rejectedReqQuery.Where("created_at BETWEEN ? AND ?", *startDate, *endDate)
	} else if startDate != nil {
		rejectedReqQuery = rejectedReqQuery.Where("created_at >= ?", *startDate)
	} else if endDate != nil {
		rejectedReqQuery = rejectedReqQuery.Where("created_at <= ?", *endDate)
	}

	err = rejectedReqQuery.Count(&rejectedCount).Error
	if err != nil {
		return nil, err
	}

	// Calculate approval and rejection rates
	var approvalRate, rejectionRate float64
	if totalRequests > 0 {
		approvalRate = (float64(approvedCount) / float64(totalRequests)) * 100
		rejectionRate = (float64(rejectedCount) / float64(totalRequests)) * 100
	}

	stats.TotalAmount = totalAmount
	stats.ApprovedAmount = approvedAmount
	stats.PendingAmount = pendingAmount
	stats.ApprovalRate = approvalRate
	stats.RejectionRate = rejectionRate
	stats.TotalRequests = totalRequests
	stats.ApprovedCount = approvedCount
	stats.PendingCount = pendingCount
	stats.RejectedCount = rejectedCount

	return stats, nil
}

// GetReimbursementMonthlyTrends retrieves monthly reimbursement trends
func (r *AnalyticsRepository) GetReimbursementMonthlyTrends(tenantID *int64, userID int64, startDate, endDate *time.Time) ([]model.MonthlyTrend, error) {
	// We'll use the same query as weekly expenses but for months
	var results []struct {
		Month    string
		Total    float64
		Approved float64
		Pending  float64
	}

	query := r.db.Model(&model.Invoice{}).
		Select(`DATE_FORMAT(created_at, '%Y-%m') as month,
				COALESCE(SUM(amount), 0) as total,
				COALESCE(SUM(CASE WHEN status IN ('confirmed', 'approved') THEN amount ELSE 0 END), 0) as approved,
				COALESCE(SUM(CASE WHEN status = 'pending' THEN amount ELSE 0 END), 0) as pending`).
		Where("user_id = ? AND amount IS NOT NULL", userID).
		Group("DATE_FORMAT(created_at, '%Y-%m')").
		Order("month")

	if tenantID != nil && *tenantID != 0 {
		query = query.Where("tenant_id = ?", *tenantID)
	} else {
		query = query.Where("tenant_id IS NULL OR tenant_id = 0")
	}

	if startDate != nil && endDate != nil {
		query = query.Where("created_at BETWEEN ? AND ?", *startDate, *endDate)
	} else if startDate != nil {
		query = query.Where("created_at >= ?", *startDate)
	} else if endDate != nil {
		query = query.Where("created_at <= ?", *endDate)
	}

	err := query.Scan(&results).Error
	if err != nil {
		return nil, err
	}

	monthlyTrends := make([]model.MonthlyTrend, len(results))
	for i, result := range results {
		// Format the month string to be more readable (e.g., "2025-01" to "1月")
		month := result.Month
		if len(month) >= 7 {
			month = month[5:] + "月" // Get the MM part and add "月"
		} else {
			month = result.Month + "月"
		}

		monthlyTrends[i] = model.MonthlyTrend{
			Month:    month,
			Total:    result.Total,
			Approved: result.Approved,
			Pending:  result.Pending,
		}
	}

	return monthlyTrends, nil
}
