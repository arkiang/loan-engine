package usecase

import (
	"loan-engine/internal/dto"
	"loan-engine/internal/model"
	"loan-engine/internal/repository"
)

type customerUsecase struct {
	repo repository.CustomerRepository
}

func NewCustomerUsecase(repo repository.CustomerRepository) CustomerUsecase {
	return &customerUsecase{repo}
}

func (s *customerUsecase) CreateCustomer(customer *model.Customer) error {
	return s.repo.Create(customer)
}

func (s *customerUsecase) GetCustomerByID(id uint) (*model.Customer, error) {
	return s.repo.FindByID(id)
}

func (s *customerUsecase) ListCustomers(filter dto.CommonFilter) ([]model.Customer, error) {
	return s.repo.FindAll(filter)
}
