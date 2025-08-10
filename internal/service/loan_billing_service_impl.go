package service

import (
	"fmt"
	"loan-engine/internal/dto"
	"loan-engine/internal/mapper"
	"loan-engine/internal/model"
	"loan-engine/internal/usecase"
	"time"
)

type loanBillingService struct {
	customerUsecase          usecase.CustomerUsecase
	loanProductUsecase       usecase.LoanProductUsecase
	loanUsecase              usecase.LoanUsecase
	repaymentScheduleUsecase usecase.RepaymentScheduleUsecase
	paymentUsecase           usecase.PaymentUsecase
}

func NewLoanBillingService(
	cu usecase.CustomerUsecase,
	lpu usecase.LoanProductUsecase,
	lu usecase.LoanUsecase,
	rsu usecase.RepaymentScheduleUsecase,
	pu usecase.PaymentUsecase,
) LoanBillingService {
	return &loanBillingService{
		customerUsecase:          cu,
		loanProductUsecase:       lpu,
		loanUsecase:              lu,
		repaymentScheduleUsecase: rsu,
		paymentUsecase:           pu,
	}
}

func (s *loanBillingService) CreateLoanWithSchedule(customerID uint, productID uint, startDateStr string) (*model.Loan, error) {
	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		return nil, fmt.Errorf("invalid start date format: %w", err)
	}

	customer, err := s.customerUsecase.GetCustomerByID(customerID)
	if err != nil {
		return nil, fmt.Errorf("customer not found: %w", err)
	}

	product, err := s.loanProductUsecase.GetLoanProductByID(productID)
	if err != nil {
		return nil, fmt.Errorf("loan product not found: %w", err)
	}

	// Calculate total with flat interest
	interest := int64(float64(product.PrincipalAmount) * product.InterestRate / 100)
	total := product.PrincipalAmount + interest

	loan := &model.Loan{
		CustomerID:         customer.ID,
		LoanProductID:      product.ID,
		Principal:          product.PrincipalAmount,
		InterestRate:       product.InterestRate,
		TotalAmount:        total,
		RepaymentCount:     product.RepaymentCount,
		RepaymentFrequency: product.RepaymentFrequency,
		StartDate:          startDate,
	}

	if err := s.loanUsecase.CreateLoan(loan); err != nil {
		return nil, fmt.Errorf("create loan failed: %w", err)
	}

	if err := s.repaymentScheduleUsecase.GenerateSchedule(loan); err != nil {
		return nil, fmt.Errorf("generate schedule failed: %w", err)
	}

	return loan, nil
}

func (s *loanBillingService) MakePayment(loanID uint) error {
	// 1. Get all overdue unpaid schedules
	schedules, err := s.repaymentScheduleUsecase.GetOverdueUnpaidSchedule(loanID)
	if err != nil {
		return fmt.Errorf("failed to get overdue schedules: %w", err)
	}

	var scheduleIds []uint

	var amount int64
	if len(schedules) == 0 {
		schedule, err := s.repaymentScheduleUsecase.GetNextUnpaidSchedule(loanID)
		if err != nil {
			return fmt.Errorf("no unpaid schedule found: %w", err)
		}

		amount = schedule.Amount
		scheduleIds = append(scheduleIds, schedule.ID)
	} else {
		for _, schedule := range schedules {
			amount += schedule.Amount
			scheduleIds = append(scheduleIds, schedule.ID)
		}
	}

	// 4. Create payment record
	payment := &model.Payment{
		LoanID:        loanID,
		Amount:        amount,
		PaidAt:        time.Now(),
		PaymentMethod: "AmarthaBank",
		Status:        model.PaymentStatusSuccess,
	}

	if err := s.paymentUsecase.MakePayment(payment); err != nil {
		return fmt.Errorf("make payment failed: %w", err)
	}

	// 5. Mark all overdue schedules as paid
	for _, value := range scheduleIds {
		if err := s.repaymentScheduleUsecase.MarkAsPaid(value, payment.PaidAt); err != nil {
			return fmt.Errorf("mark schedule %d as paid failed: %w", value, err)
		}
	}

	return nil
}

func (s *loanBillingService) GetLoanOutstanding(loanID uint) (int64, error) {
	return s.loanUsecase.GetOutstandingAmount(loanID)
}

func (s *loanBillingService) CheckDelinquency(loanID uint) (bool, error) {
	return s.loanUsecase.IsDelinquent(loanID)
}

func (s *loanBillingService) GetCustomerByID(customerID uint, filter dto.CommonFilter) (*dto.CustomerDTO, error) {
	customer, err := s.customerUsecase.GetCustomerByID(customerID)
	if err != nil {
		return nil, fmt.Errorf("customer not found: %w", err)
	}

	loans, err := s.loanUsecase.ListLoansByCustomerID(customerID, filter)
	if err != nil {
		return nil, fmt.Errorf("get loans failed: %w", err)
	}

	isDelinquent, err := s.repaymentScheduleUsecase.IsDelinquent(customerID)
	if err != nil {
		return nil, fmt.Errorf("get delinquent customer failed: %w", err)
	}

	outstandingAmount, err := s.repaymentScheduleUsecase.GetTotalOutstandingAmount(customerID)
	if err != nil {
		return nil, fmt.Errorf("get outstanding amount failed: %w", err)
	}

	return mapper.ToCustomerWithLoansDTO(customer, loans, isDelinquent, outstandingAmount), nil
}
