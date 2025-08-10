package repository

import "loan-engine/internal/model"

type PaymentRepository interface {
	Create(payment *model.Payment) error
	FindByLoanID(loanID uint) ([]model.Payment, error)
}
