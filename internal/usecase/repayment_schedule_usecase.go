package usecase

import (
	"loan-engine/internal/model"
	"time"
)

type RepaymentScheduleUsecase interface {
	GenerateSchedule(loan *model.Loan) error
	GetSchedulesByLoanID(loanID uint) ([]model.RepaymentSchedule, error)
	GetNextUnpaidSchedule(loanID uint) (*model.RepaymentSchedule, error)
	GetOverdueUnpaidSchedule(loanID uint) ([]model.RepaymentSchedule, error)
	MarkAsPaid(repaymentID uint, paidAt time.Time) error
	IsDelinquent(customerID uint) (bool, error)
	GetTotalOutstandingAmount(customerID uint) (int64, error)
}
