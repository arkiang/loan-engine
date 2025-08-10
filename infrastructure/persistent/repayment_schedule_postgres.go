package persistent

import (
	"gorm.io/gorm"
	"loan-engine/internal/model"
	"loan-engine/internal/repository"
	"time"
)

type scheduleRepo struct {
	db *gorm.DB
}

func NewRepaymentScheduleRepository(db *gorm.DB) repository.RepaymentScheduleRepository {
	return &scheduleRepo{db}
}

func (r *scheduleRepo) BulkCreate(schedules []model.RepaymentSchedule) error {
	return r.db.Create(&schedules).Error
}

func (r *scheduleRepo) FindByLoanID(loanID uint) ([]model.RepaymentSchedule, error) {
	var schedules []model.RepaymentSchedule
	err := r.db.Where("loan_id = ?", loanID).Order("sequence ASC").Find(&schedules).Error
	return schedules, err
}

func (r *scheduleRepo) MarkAsPaid(scheduleID uint, paidAt time.Time) error {
	return r.db.Model(&model.RepaymentSchedule{}).
		Where("id = ?", scheduleID).
		Updates(map[string]interface{}{"is_paid": true, "paid_at": paidAt}).Error
}

func (r *scheduleRepo) GetUnpaidSchedules(loanID uint) ([]model.RepaymentSchedule, error) {
	var schedules []model.RepaymentSchedule
	err := r.db.
		Where("loan_id = ? AND is_paid = ?", loanID, false).
		Order("sequence ASC").
		Find(&schedules).Error
	return schedules, err
}

func (u *scheduleRepo) GetOverdueUnpaidSchedules(loanID uint) ([]model.RepaymentSchedule, error) {
	var schedules []model.RepaymentSchedule
	err := u.db.
		Where("loan_id = ? AND is_paid = ? AND due_date < ?", loanID, false, time.Now()).
		Order("sequence ASC").
		Find(&schedules).Error
	return schedules, err
}

func (r *scheduleRepo) GetCustomerOutstandingAmount(customerID uint) (int64, error) {
	var outstanding int64
	err := r.db.
		Model(&model.RepaymentSchedule{}).
		Select("COALESCE(SUM(repayment_schedules.amount), 0)").
		Joins("JOIN loans ON loans.id = repayment_schedules.loan_id").
		Where("loans.customer_id = ? AND repayment_schedules.is_paid = ?", customerID, false).
		Scan(&outstanding).Error

	if err != nil {
		return 0, err
	}
	return outstanding, nil
}

// IsDelinquent define when there are 2 unpaid loan
func (r *scheduleRepo) IsCustomerDelinquent(customerID uint) (bool, error) {
	var count int64
	err := r.db.Model(&model.RepaymentSchedule{}).
		Joins("JOIN loans ON loans.id = repayment_schedules.loan_id").
		Where("loans.customer_id = ? AND repayment_schedules.is_paid = false AND repayment_schedules.due_date < ?", customerID, time.Now()).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 1, nil
}
