package dto

import "time"

type CreateCustomerRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"omitempty,email"`
	Phone string `json:"phone" binding:"omitempty"`
}

type CustomerDTO struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`

	Loans            []LoanDetailDTO `json:"loans"`
	IsDelinquent     bool            `json:"is_delinquent"`
	TotalOutstanding int64           `json:"total_outstanding"`
}
