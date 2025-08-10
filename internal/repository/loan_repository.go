package repository

import (
	"loan-engine/internal/dto"
	"loan-engine/internal/model"
)

type LoanRepository interface {
	Create(loan *model.Loan) error
	FindByID(id uint) (*model.Loan, error)
	FindByCustomerID(customerID uint, filter dto.CommonFilter) ([]model.Loan, error)
	GetOutstandingAmount(loanID uint) (int64, error)
	IsDelinquent(loanID uint) (bool, error)
}
