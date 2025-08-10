package dto

type CreateLoanProductRequest struct {
	Name               string  `json:"name" binding:"required"`
	Description        string  `json:"description"`
	PrincipalAmount    int64   `json:"principal_amount" binding:"required"`
	InterestRate       float64 `json:"interest_rate" binding:"required"`
	RepaymentCount     int     `json:"repayment_count" binding:"required"`
	RepaymentFrequency string  `json:"repayment_frequency" binding:"required,oneof=daily weekly monthly"`
	IsActive           *bool   `json:"is_active"` // Optional; default true in DB
}

type LoanProductDTO struct {
	ID           uint    `json:"id"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	InterestRate float64 `json:"interest_rate"`
}
