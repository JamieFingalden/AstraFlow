package model

import (
	"time"
)

// DashboardMetrics represents key metrics for the dashboard
type DashboardMetrics struct {
	MonthlyExpense float64 `json:"monthly_expense"`
	AutoProcessed  int64   `json:"auto_processed"`
	AIAccuracy     float64 `json:"ai_accuracy"`
	PendingReviews int64   `json:"pending_reviews"`
}

// ExpenseCategory represents expense distribution by category for pie chart
type ExpenseCategory struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
	Color string  `json:"color"`
}

// WeeklyExpense represents weekly expense trends for line chart
type WeeklyExpense struct {
	Week    string  `json:"week"`
	Expense float64 `json:"expense"`
}

// RecentBill represents recent bill data for table
type RecentBill struct {
	ID       int64     `json:"id"`
	Date     time.Time `json:"date"`
	Vendor   string    `json:"vendor"`
	Category string    `json:"category"`
	Amount   float64   `json:"amount"`
	Status   string    `json:"status"`
}

// AIInsight represents AI-generated insights and warnings
type AIInsight struct {
	ID      int64  `json:"id"`
	Message string `json:"message"`
	Type    string `json:"type"` // warning, info, alert
}

// DashboardData represents the complete dashboard response
type DashboardData struct {
	Metrics           DashboardMetrics  `json:"metrics"`
	ExpenseCategories []ExpenseCategory `json:"expense_categories"`
	WeeklyExpenses    []WeeklyExpense   `json:"weekly_expenses"`
	RecentBills       []RecentBill      `json:"recent_bills"`
	AIInsights        []AIInsight       `json:"ai_insights"`
}

// AnalyticsRequest represents the request parameters for analytics
type AnalyticsRequest struct {
	StartDate *time.Time `json:"start_date,omitempty" form:"start_date"`
	EndDate   *time.Time `json:"end_date,omitempty" form:"end_date"`
}

// ReimbursementStatistics represents key reimbursement statistics
type ReimbursementStatistics struct {
	TotalAmount     float64 `json:"total_amount"`
	ApprovedAmount  float64 `json:"approved_amount"`
	PendingAmount   float64 `json:"pending_amount"`
	RejectionRate   float64 `json:"rejection_rate"`
	ApprovalRate    float64 `json:"approval_rate"`
	TotalRequests   int64   `json:"total_requests"`
	ApprovedCount   int64   `json:"approved_count"`
	PendingCount    int64   `json:"pending_count"`
	RejectedCount   int64   `json:"rejected_count"`
}

// MonthlyTrend represents monthly reimbursement trends for bar chart
type MonthlyTrend struct {
	Month string  `json:"month"`
	Total float64 `json:"total"`
	Approved float64 `json:"approved"`
	Pending float64 `json:"pending"`
}

// ReimbursementStatisticsData represents the complete reimbursement statistics response
type ReimbursementStatisticsData struct {
	Statistics    ReimbursementStatistics `json:"statistics"`
	MonthlyTrends []MonthlyTrend          `json:"monthly_trends"`
	Categories    []ExpenseCategory       `json:"categories"`
}