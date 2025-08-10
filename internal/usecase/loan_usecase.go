package usecase

import (
	"loan-engine/internal/dto"
	"loan-engine/internal/model"
)

type LoanUsecase interface {
	CreateLoan(loan *model.Loan) error
	GetLoanByID(id uint) (*model.Loan, error)
	GetOutstandingAmount(loanID uint) (int64, error)
	IsDelinquent(loanID uint) (bool, error)
	ListLoansByCustomerID(customerID uint, filter dto.CommonFilter) ([]model.Loan, error)
}
