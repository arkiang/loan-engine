package persistent

import (
	"gorm.io/gorm"
	"loan-engine/internal/dto"
	"loan-engine/internal/model"
	"loan-engine/internal/repository"
)

type customerRepo struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) repository.CustomerRepository {
	return &customerRepo{db}
}

func (r *customerRepo) Create(customer *model.Customer) error {
	return r.db.Create(customer).Error
}

func (r *customerRepo) FindByID(id uint) (*model.Customer, error) {
	var customer model.Customer
	err := r.db.First(&customer, id).Error
	return &customer, err
}

func (r *customerRepo) FindAll(filter dto.CommonFilter) ([]model.Customer, error) {
	var customers []model.Customer
	query := r.db.Order("created_at DESC")

	if filter.Page > 0 && filter.PageSize > 0 {
		offset := (filter.Page - 1) * filter.PageSize
		query = query.Offset(offset).Limit(filter.PageSize)
	}

	err := query.Find(&customers).Error
	return customers, err
}
