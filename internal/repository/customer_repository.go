package repository

import (
	"loan-engine/internal/dto"
	"loan-engine/internal/model"
)

type CustomerRepository interface {
	Create(customer *model.Customer) error
	FindByID(id uint) (*model.Customer, error)
	FindAll(filter dto.CommonFilter) ([]model.Customer, error)
}
