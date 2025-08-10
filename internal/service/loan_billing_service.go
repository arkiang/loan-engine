package service

import (
	"loan-engine/internal/dto"
	"loan-engine/internal/model"
)

type LoanBillingService interface {
	CreateLoanWithSchedule(customerID uint, productID uint, startDate string) (*model.Loan, error)
	MakeScheduledPayment(loanID uint) error
	MakePayment(loanID uint) error
	GetLoanOutstanding(loanID uint) (int64, error)
	CheckDelinquency(loanID uint) (bool, error)
	GetCustomerByID(customerID uint, filter dto.CommonFilter) (*dto.CustomerDTO, error)
}
