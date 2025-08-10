package usecase

import (
	"loan-engine/internal/dto"
	"loan-engine/internal/model"
)

type CustomerUsecase interface {
	CreateCustomer(customer *model.Customer) error
	GetCustomerByID(id uint) (*model.Customer, error)
	ListCustomers(filter dto.CommonFilter) ([]model.Customer, error)
}
