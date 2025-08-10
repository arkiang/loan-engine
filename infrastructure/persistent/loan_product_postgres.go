package persistent

import (
	"gorm.io/gorm"
	"loan-engine/internal/dto"
	"loan-engine/internal/model"
	"loan-engine/internal/repository"
)

type loanProductRepo struct {
	db *gorm.DB
}

func NewLoanProductRepository(db *gorm.DB) repository.LoanProductRepository {
	return &loanProductRepo{db}
}

func (r *loanProductRepo) Create(product *model.LoanProduct) error {
	return r.db.Create(product).Error
}

func (r *loanProductRepo) FindByID(id uint) (*model.LoanProduct, error) {
	var product model.LoanProduct
	err := r.db.First(&product, id).Error
	return &product, err
}

func (r *loanProductRepo) FindActive(filter dto.CommonFilter) ([]model.LoanProduct, error) {
	var products []model.LoanProduct
	query := r.db.Order("created_at DESC")

	if filter.Page > 0 && filter.PageSize > 0 {
		offset := (filter.Page - 1) * filter.PageSize
		query = query.Offset(offset).Limit(filter.PageSize)
	}

	err := query.Where("is_active = ?", true).Find(&products).Error
	return products, err
}
