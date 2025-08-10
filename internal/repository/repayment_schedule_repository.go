package repository

import (
	"loan-engine/internal/model"
	"time"
)

type RepaymentScheduleRepository interface {
	BulkCreate(schedules []model.RepaymentSchedule) error
	FindByLoanID(loanID uint) ([]model.RepaymentSchedule, error)
	MarkAsPaid(scheduleID uint, paidAt time.Time) error
	GetUnpaidSchedules(loanID uint) ([]model.RepaymentSchedule, error)
	GetOverdueUnpaidSchedules(loanID uint) ([]model.RepaymentSchedule, error)
	GetCustomerOutstandingAmount(customerID uint) (int64, error)
	IsCustomerDelinquent(customerID uint) (bool, error)
}
