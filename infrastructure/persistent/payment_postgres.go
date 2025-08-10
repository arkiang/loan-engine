package persistent

import (
	"gorm.io/gorm"
	"loan-engine/internal/model"
	"loan-engine/internal/repository"
)

type paymentRepo struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) repository.PaymentRepository {
	return &paymentRepo{db}
}

func (r *paymentRepo) Create(payment *model.Payment) error {
	return r.db.Create(payment).Error
}

func (r *paymentRepo) FindByLoanID(loanID uint) ([]model.Payment, error) {
	var payments []model.Payment
	err := r.db.Where("loan_id = ?", loanID).Order("paid_at DESC").Find(&payments).Error
	return payments, err
}
