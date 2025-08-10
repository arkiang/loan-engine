package usecase

import (
	"loan-engine/internal/model"
	"loan-engine/internal/repository"
)

type paymentUsecase struct {
	repo repository.PaymentRepository
}

func NewPaymentUsecase(repo repository.PaymentRepository) PaymentUsecase {
	return &paymentUsecase{repo}
}

func (s *paymentUsecase) MakePayment(payment *model.Payment) error {
	return s.repo.Create(payment)
}

func (s *paymentUsecase) GetPaymentsByLoanID(loanID uint) ([]model.Payment, error) {
	return s.repo.FindByLoanID(loanID)
}
