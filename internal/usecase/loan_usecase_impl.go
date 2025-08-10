package usecase

import (
	"loan-engine/internal/dto"
	"loan-engine/internal/model"
	"loan-engine/internal/repository"
)

type loanUsecase struct {
	repo repository.LoanRepository
}

func NewLoanUsecase(
	repo repository.LoanRepository,
) LoanUsecase {
	return &loanUsecase{repo}
}

func (s *loanUsecase) CreateLoan(loan *model.Loan) error {
	return s.repo.Create(loan)
}

func (s *loanUsecase) GetLoanByID(id uint) (*model.Loan, error) {
	return s.repo.FindByID(id)
}

func (s *loanUsecase) ListLoansByCustomerID(customerID uint, filter dto.CommonFilter) ([]model.Loan, error) {
	return s.repo.FindByCustomerID(customerID, filter)
}

func (s *loanUsecase) GetOutstandingAmount(loanID uint) (int64, error) {
	return s.repo.GetOutstandingAmount(loanID)
}

func (s *loanUsecase) IsDelinquent(customerID uint) (bool, error) {
	return s.repo.IsDelinquent(customerID)
}
