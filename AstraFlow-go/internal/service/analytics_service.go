package service

import (
	"AstraFlow-go/internal/model"
	"AstraFlow-go/internal/repository"
	"errors"
	"time"
)

// AnalyticsService handles analytics business logic
type AnalyticsService struct {
	analyticsRepo *repository.AnalyticsRepository
}

// NewAnalyticsService creates a new analytics service instance
func NewAnalyticsService() *AnalyticsService {
	return &AnalyticsService{
		analyticsRepo: repository.NewAnalyticsRepository(),
	}
}

// GetDashboardData retrieves all dashboard analytics data
func (s *AnalyticsService) GetDashboardData(tenantID *int64, userID int64, startDate, endDate *time.Time) (*model.DashboardData, error) {
	if userID <= 0 {
		return nil, errors.New("user ID must be greater than 0")
	}

	// Get metrics
	metrics, err := s.analyticsRepo.GetDashboardMetrics(tenantID, userID, startDate, endDate)
	if err != nil {
		return nil, err
	}

	// Get expense categories
	categories, err := s.analyticsRepo.GetExpenseCategories(tenantID, userID, startDate, endDate)
	if err != nil {
		return nil, err
	}

	// Get monthly expenses
	monthlyExpenses, err := s.analyticsRepo.GetMonthlyExpenses(tenantID, userID, startDate, endDate)
	if err != nil {
		return nil, err
	}

	// Get recent bills
	recentBills, err := s.analyticsRepo.GetRecentBills(tenantID, userID, 5, startDate, endDate) // Limit to 5 recent bills
	if err != nil {
		return nil, err
	}

	// For now, return empty AI insights array as requested by the user
	// AI insights will be implemented later in the Flask backend
	aiInsights := []model.AIInsight{}

	// Construct dashboard data response
	dashboardData := &model.DashboardData{
		Metrics:           *metrics,
		ExpenseCategories: categories,
		MonthlyExpenses:   monthlyExpenses,
		RecentBills:       recentBills,
		AIInsights:        aiInsights,
	}

	return dashboardData, nil
}

// GetMetrics retrieves key dashboard metrics only
func (s *AnalyticsService) GetMetrics(tenantID *int64, userID int64, startDate, endDate *time.Time) (*model.DashboardMetrics, error) {
	if userID <= 0 {
		return nil, errors.New("user ID must be greater than 0")
	}

	return s.analyticsRepo.GetDashboardMetrics(tenantID, userID, startDate, endDate)
}

// GetExpenseCategories retrieves expense breakdown by category only
func (s *AnalyticsService) GetExpenseCategories(tenantID *int64, userID int64, startDate, endDate *time.Time) ([]model.ExpenseCategory, error) {
	if userID <= 0 {
		return nil, errors.New("user ID must be greater than 0")
	}

	return s.analyticsRepo.GetExpenseCategories(tenantID, userID, startDate, endDate)
}

// GetWeeklyExpenses retrieves weekly expense trends only
func (s *AnalyticsService) GetWeeklyExpenses(tenantID *int64, userID int64, startDate, endDate *time.Time) ([]model.MonthlyExpense, error) {
	if userID <= 0 {
		return nil, errors.New("user ID must be greater than 0")
	}

	return s.analyticsRepo.GetMonthlyExpenses(tenantID, userID, startDate, endDate)
}

// GetMonthlyExpenses retrieves monthly expense trends only
func (s *AnalyticsService) GetMonthlyExpenses(tenantID *int64, userID int64, startDate, endDate *time.Time) ([]model.MonthlyExpense, error) {
	if userID <= 0 {
		return nil, errors.New("user ID must be greater than 0")
	}

	return s.analyticsRepo.GetMonthlyExpenses(tenantID, userID, startDate, endDate)
}

// GetRecentBills retrieves recent bills only
func (s *AnalyticsService) GetRecentBills(tenantID *int64, userID int64, limit int, startDate, endDate *time.Time) ([]model.RecentBill, error) {
	if userID <= 0 {
		return nil, errors.New("user ID must be greater than 0")
	}

	if limit <= 0 {
		limit = 10 // Default limit
	}

	return s.analyticsRepo.GetRecentBills(tenantID, userID, limit, startDate, endDate)
}

// CalculateMonthOverMonthChange calculates the percentage change between two months
func (s *AnalyticsService) CalculateMonthOverMonthChange(currentMonth, previousMonth float64) float64 {
	if previousMonth == 0 {
		if currentMonth == 0 {
			return 0
		}
		return 100 // If previous month was 0 and current is not, infinite growth
	}
	return ((currentMonth - previousMonth) / previousMonth) * 100
}

// FormatDateRange creates a standardized date range for analytics
func (s *AnalyticsService) FormatDateRange(startDate, endDate *time.Time) (time.Time, time.Time) {
	var start, end time.Time

	if startDate == nil {
		// Default to beginning of current month if no start date provided
		now := time.Now()
		start = time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	} else {
		start = *startDate
	}

	if endDate == nil {
		// Default to end of current month if no end date provided
		now := time.Now()
		end = time.Date(now.Year(), now.Month()+1, 0, 23, 59, 59, 999999999, now.Location())
	} else {
		end = *endDate
	}

	return start, end
}

// GetReimbursementStatisticsData retrieves all reimbursement statistics data
func (s *AnalyticsService) GetReimbursementStatisticsData(tenantID *int64, userID int64, startDate, endDate *time.Time) (*model.ReimbursementStatisticsData, error) {
	if userID <= 0 {
		return nil, errors.New("user ID must be greater than 0")
	}

	// For reimbursement statistics, if no dates are provided, show all data
	// We only format dates if they are provided by the client
	var start, end *time.Time
	if startDate != nil && endDate != nil {
		sDate, eDate := s.FormatDateRange(startDate, endDate)
		start = &sDate
		end = &eDate
	}
	// If dates are nil, keep them as nil to show all data

	// Get reimbursement statistics
	statistics, err := s.analyticsRepo.GetReimbursementStatistics(tenantID, userID, start, end)
	if err != nil {
		return nil, err
	}

	// Get monthly trends
	monthlyTrends, err := s.analyticsRepo.GetReimbursementMonthlyTrends(tenantID, userID, start, end)
	if err != nil {
		return nil, err
	}

	// Get expense categories
	categories, err := s.analyticsRepo.GetExpenseCategories(tenantID, userID, start, end)
	if err != nil {
		return nil, err
	}

	// Construct reimbursement statistics data response
	reimbursementData := &model.ReimbursementStatisticsData{
		Statistics:    *statistics,
		MonthlyTrends: monthlyTrends,
		Categories:    categories,
	}

	return reimbursementData, nil
}

// GetReimbursementStatistics retrieves key reimbursement statistics only
func (s *AnalyticsService) GetReimbursementStatistics(tenantID *int64, userID int64, startDate, endDate *time.Time) (*model.ReimbursementStatistics, error) {
	if userID <= 0 {
		return nil, errors.New("user ID must be greater than 0")
	}

	// For reimbursement statistics, if no dates are provided, show all data
	// We only format dates if they are provided by the client
	var start, end *time.Time
	if startDate != nil && endDate != nil {
		sDate, eDate := s.FormatDateRange(startDate, endDate)
		start = &sDate
		end = &eDate
	}
	// If dates are nil, keep them as nil to show all data

	return s.analyticsRepo.GetReimbursementStatistics(tenantID, userID, start, end)
}

// GetReimbursementMonthlyTrends retrieves monthly reimbursement trends only
func (s *AnalyticsService) GetReimbursementMonthlyTrends(tenantID *int64, userID int64, startDate, endDate *time.Time) ([]model.MonthlyTrend, error) {
	if userID <= 0 {
		return nil, errors.New("user ID must be greater than 0")
	}

	// For reimbursement statistics, if no dates are provided, show all data
	// We only format dates if they are provided by the client
	var start, end *time.Time
	if startDate != nil && endDate != nil {
		sDate, eDate := s.FormatDateRange(startDate, endDate)
		start = &sDate
		end = &eDate
	}
	// If dates are nil, keep them as nil to show all data

	return s.analyticsRepo.GetReimbursementMonthlyTrends(tenantID, userID, start, end)
}
