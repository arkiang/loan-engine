package persistent

import (
	"gorm.io/gorm"
	"loan-engine/internal/dto"
	"loan-engine/internal/model"
	"loan-engine/internal/repository"
	"time"
)

type loanRepo struct {
	db *gorm.DB
}

func NewLoanRepository(db *gorm.DB) repository.LoanRepository {
	return &loanRepo{db}
}

func (r *loanRepo) Create(loan *model.Loan) error {
	return r.db.Create(loan).Error
}

func (r *loanRepo) FindByID(id uint) (*model.Loan, error) {
	var loan model.Loan
	err := r.db.Preload("Payments").Preload("Schedules").Preload("Product").First(&loan, id).Error

	return &loan, err
}

func (r *loanRepo) FindByCustomerID(customerID uint, filter dto.CommonFilter) ([]model.Loan, error) {
	var loans []model.Loan
	query := r.db.
		Preload("Payments", func(db *gorm.DB) *gorm.DB {
			return db.Order("paid_at DESC")
		}).
		Preload("Schedules", func(db *gorm.DB) *gorm.DB {
			return db.Where("is_paid = ?", false).
				Order("sequence ASC")
		}).
		Preload("Product").
		Order("created_at DESC")

	if filter.Page > 0 && filter.PageSize > 0 {
		offset := (filter.Page - 1) * filter.PageSize
		query = query.Offset(offset).Limit(filter.PageSize)
	}

	err := query.Where("customer_id = ?", customerID).Find(&loans).Error
	return loans, err
}

func (r *loanRepo) GetOutstandingAmount(loanID uint) (int64, error) {
	var outstanding int64
	err := r.db.
		Model(&model.RepaymentSchedule{}).
		Select("COALESCE(SUM(amount), 0)").
		Where("loan_id = ? AND is_paid = ?", loanID, false).
		Scan(&outstanding).Error

	if err != nil {
		return 0, err
	}
	return outstanding, nil
}

// IsDelinquent define when there are 2 unpaid loan
func (r *loanRepo) IsDelinquent(loanID uint) (bool, error) {
	var count int64
	err := r.db.Model(&model.RepaymentSchedule{}).
		Where("loan_id = ? AND due_date < ? AND is_paid = ?", loanID, time.Now(), false).
		Count(&count).Error
	return count > 1, err
}
