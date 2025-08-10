package usecase

import "loan-engine/internal/model"

type PaymentUsecase interface {
	MakePayment(payment *model.Payment) error
	GetPaymentsByLoanID(loanID uint) ([]model.Payment, error)
}
