package usecase

import (
	"loan-engine/internal/dto"
	"loan-engine/internal/model"
	"loan-engine/internal/repository"
)

type loanProductUsecase struct {
	repo repository.LoanProductRepository
}

func NewLoanProductUsecase(repo repository.LoanProductRepository) LoanProductUsecase {
	return &loanProductUsecase{repo}
}

func (s *loanProductUsecase) CreateLoanProduct(product *model.LoanProduct) error {
	return s.repo.Create(product)
}

func (s *loanProductUsecase) GetLoanProductByID(id uint) (*model.LoanProduct, error) {
	return s.repo.FindByID(id)
}

func (s *loanProductUsecase) ListActiveLoanProducts(filter dto.CommonFilter) ([]model.LoanProduct, error) {
	return s.repo.FindActive(filter)
}
