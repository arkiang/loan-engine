package repository

import (
	"loan-engine/internal/dto"
	"loan-engine/internal/model"
)

type LoanProductRepository interface {
	Create(product *model.LoanProduct) error
	FindByID(id uint) (*model.LoanProduct, error)
	FindActive(filter dto.CommonFilter) ([]model.LoanProduct, error)
}
