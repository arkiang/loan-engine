package usecase

import (
	"fmt"
	"loan-engine/internal/model"
	"loan-engine/internal/repository"
	"time"
)

type repaymentScheduleUsecase struct {
	repo repository.RepaymentScheduleRepository
}

func NewRepaymentScheduleUsecase(repo repository.RepaymentScheduleRepository) RepaymentScheduleUsecase {
	return &repaymentScheduleUsecase{repo}
}

func (s *repaymentScheduleUsecase) GenerateSchedule(loan *model.Loan) error {
	var schedules []model.RepaymentSchedule
	amount := loan.TotalAmount / int64(loan.RepaymentCount)
	start := loan.StartDate

	for i := 1; i <= loan.RepaymentCount; i++ {
		var dueDate time.Time
		switch loan.RepaymentFrequency {
		case "weekly":
			dueDate = start.AddDate(0, 0, 7*(i-1))
		case "daily":
			dueDate = start.AddDate(0, 0, i-1)
		case "monthly":
			dueDate = start.AddDate(0, i-1, 0)
		}

		schedules = append(schedules, model.RepaymentSchedule{
			LoanID:   loan.ID,
			Sequence: i,
			DueDate:  dueDate,
			Amount:   amount,
		})
	}

	return s.repo.BulkCreate(schedules)
}

func (s *repaymentScheduleUsecase) GetSchedulesByLoanID(loanID uint) ([]model.RepaymentSchedule, error) {
	return s.repo.FindByLoanID(loanID)
}

func (s *repaymentScheduleUsecase) GetNextUnpaidSchedule(loanID uint) (*model.RepaymentSchedule, error) {
	schedule, err := s.repo.GetUnpaidSchedules(loanID)
	if err != nil {
		return nil, fmt.Errorf("get next unpaid schedule failed: %w", err)
	}

	if len(schedule) == 0 {
		return nil, fmt.Errorf("active loan not found : %v", loanID)
	}

	return &schedule[0], nil
}

func (s *repaymentScheduleUsecase) GetOverdueUnpaidSchedule(loanID uint) ([]model.RepaymentSchedule, error) {
	return s.repo.GetOverdueUnpaidSchedules(loanID)
}

func (s *repaymentScheduleUsecase) MarkAsPaid(repaymentID uint, paidAt time.Time) error {
	return s.repo.MarkAsPaid(repaymentID, paidAt)
}

func (s *repaymentScheduleUsecase) IsDelinquent(customerID uint) (bool, error) {
	return s.repo.IsCustomerDelinquent(customerID)
}

func (s *repaymentScheduleUsecase) GetTotalOutstandingAmount(customerID uint) (int64, error) {
	return s.repo.GetCustomerOutstandingAmount(customerID)
}
