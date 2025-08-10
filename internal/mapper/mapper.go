package mapper

import (
	"loan-engine/internal/dto"
	"loan-engine/internal/model"
)

func ToCustomerWithLoansDTO(c *model.Customer, loans []model.Loan, IsDelinquent bool, outstandingAmount int64) *dto.CustomerDTO {
	loanDTOs := make([]dto.LoanDetailDTO, 0, len(loans))
	for _, loan := range loans {
		// Map LoanProduct
		productDTO := dto.LoanProductDTO{
			ID:           loan.Product.ID,
			Name:         loan.Product.Name,
			Description:  loan.Product.Description,
			InterestRate: loan.Product.InterestRate,
		}

		// Map RepaymentSchedule
		scheduleDTOs := make([]dto.RepaymentScheduleDTO, 0, len(loan.Schedules))
		for _, schedule := range loan.Schedules {
			scheduleDTOs = append(scheduleDTOs, dto.RepaymentScheduleDTO{
				ID:      schedule.ID,
				DueDate: schedule.DueDate,
				Amount:  schedule.Amount,
				IsPaid:  schedule.IsPaid,
			})
		}

		// Map Payments
		paymentDTOs := make([]dto.PaymentDTO, 0, len(loan.Payments))
		for _, pay := range loan.Payments {
			paymentDTOs = append(paymentDTOs, dto.PaymentDTO{
				ID:     pay.ID,
				Amount: pay.Amount,
				PaidAt: pay.PaidAt,
				Method: pay.PaymentMethod,
			})
		}

		// Add LoanDTO
		loanDTOs = append(loanDTOs, dto.LoanDetailDTO{
			ID:          loan.ID,
			Principal:   loan.Principal,
			TotalAmount: loan.TotalAmount,
			StartDate:   loan.StartDate,
			Product:     productDTO,
			Schedules:   scheduleDTOs,
			Payments:    paymentDTOs,
		})
	}

	return &dto.CustomerDTO{
		ID:               c.ID,
		Name:             c.Name,
		Email:            c.Email,
		Loans:            loanDTOs,
		IsDelinquent:     IsDelinquent,
		TotalOutstanding: outstandingAmount,
	}
}
