package usecase

import (
	"loan-engine/internal/dto"
	"loan-engine/internal/model"
)

type LoanProductUsecase interface {
	CreateLoanProduct(product *model.LoanProduct) error
	GetLoanProductByID(id uint) (*model.LoanProduct, error)
	ListActiveLoanProducts(filter dto.CommonFilter) ([]model.LoanProduct, error)
}
