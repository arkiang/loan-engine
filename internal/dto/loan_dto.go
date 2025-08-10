package dto

import "time"

type CreateLoanRequest struct {
	CustomerID uint   `json:"customer_id" binding:"required"`
	ProductID  uint   `json:"product_id" binding:"required"`
	StartDate  string `json:"start_date" binding:"required"` // e.g. "2025-08-10"
}

type LoanResponse struct {
	LoanID         uint   `json:"loan_id"`
	CustomerID     uint   `json:"customer_id"`
	ProductID      uint   `json:"product_id"`
	Principal      int64  `json:"principal"`
	TotalAmount    int64  `json:"total_amount"`
	RepaymentCount int    `json:"repayment_count"`
	Frequency      string `json:"repayment_frequency"`
	StartDate      string `json:"start_date"`
}

type OutstandingResponse struct {
	LoanID      uint  `json:"loan_id"`
	Outstanding int64 `json:"outstanding"`
	IsFullyPaid bool  `json:"is_fully_paid"`
}

type DelinquencyResponse struct {
	LoanID       uint `json:"loan_id"`
	IsDelinquent bool `json:"is_delinquent"`
}

type LoanDetailDTO struct {
	ID                 uint                   `json:"id"`
	Principal          int64                  `json:"principal"`
	InterestRate       float64                `json:"interest_rate"`
	TotalAmount        int64                  `json:"total_amount"`
	RepaymentFrequency string                 `json:"repayment_frequency"`
	StartDate          time.Time              `json:"start_date"`
	Product            LoanProductDTO         `json:"product"`
	Schedules          []RepaymentScheduleDTO `json:"schedules"`
	Payments           []PaymentDTO           `json:"payments"`
}
