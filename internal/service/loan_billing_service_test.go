package service

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"loan-engine/internal/dto"
	"loan-engine/internal/mapper"
	"loan-engine/internal/model"
	mocks "loan-engine/mocks/usecase"
	"testing"
)

func setupLoanBillingTest(t *testing.T) (*gomock.Controller, LoanBillingService,
	*mocks.MockCustomerUsecase,
	*mocks.MockLoanProductUsecase,
	*mocks.MockLoanUsecase,
	*mocks.MockRepaymentScheduleUsecase,
	*mocks.MockPaymentUsecase) {

	ctrl := gomock.NewController(t)

	mockCU := mocks.NewMockCustomerUsecase(ctrl)
	mockLPU := mocks.NewMockLoanProductUsecase(ctrl)
	mockLU := mocks.NewMockLoanUsecase(ctrl)
	mockRSU := mocks.NewMockRepaymentScheduleUsecase(ctrl)
	mockPU := mocks.NewMockPaymentUsecase(ctrl)

	svc := NewLoanBillingService(mockCU, mockLPU, mockLU, mockRSU, mockPU)

	return ctrl, svc, mockCU, mockLPU, mockLU, mockRSU, mockPU
}

func TestCreateLoanWithSchedule_Success(t *testing.T) {
	ctrl, svc, mockCU, mockLPU, mockLU, mockRSU, _ := setupLoanBillingTest(t)
	defer ctrl.Finish()

	startDate := "2025-08-01"

	customer := &model.Customer{ID: 1}
	product := &model.LoanProduct{
		ID:                 1,
		PrincipalAmount:    1000,
		InterestRate:       10,
		RepaymentCount:     5,
		RepaymentFrequency: "weekly",
	}

	mockCU.EXPECT().GetCustomerByID(uint(1)).Return(customer, nil)
	mockLPU.EXPECT().GetLoanProductByID(uint(1)).Return(product, nil)
	mockLU.EXPECT().CreateLoan(gomock.Any()).Return(nil)
	mockRSU.EXPECT().GenerateSchedule(gomock.Any()).Return(nil)

	loan, err := svc.CreateLoanWithSchedule(1, 1, startDate)
	assert.NoError(t, err)
	assert.Equal(t, int64(1100), loan.TotalAmount) // 1000 + 10% interest
}

func TestCreateLoanWithSchedule_InvalidDate(t *testing.T) {
	ctrl, svc, _, _, _, _, _ := setupLoanBillingTest(t)
	defer ctrl.Finish()

	loan, err := svc.CreateLoanWithSchedule(1, 1, "bad-date")
	assert.Nil(t, loan)
	assert.Error(t, err)
}

func TestCreateLoanWithSchedule_NotFoundCustomer(t *testing.T) {
	ctrl, svc, mockCU, _, _, _, _ := setupLoanBillingTest(t)
	defer ctrl.Finish()

	mockCU.EXPECT().GetCustomerByID(uint(1)).Return(nil, fmt.Errorf("customer not found"))

	startDate := "2025-08-01"
	loan, err := svc.CreateLoanWithSchedule(1, 1, startDate)
	assert.Nil(t, loan)
	assert.Error(t, err)
}

func TestCreateLoanWithSchedule_NotFoundLoanProduct(t *testing.T) {
	ctrl, svc, mockCU, mockLPU, _, _, _ := setupLoanBillingTest(t)
	defer ctrl.Finish()

	startDate := "2025-08-01"

	customer := &model.Customer{ID: 1}

	mockCU.EXPECT().GetCustomerByID(uint(1)).Return(customer, nil)
	mockLPU.EXPECT().GetLoanProductByID(uint(1)).Return(nil, fmt.Errorf("loan product not found"))

	loan, err := svc.CreateLoanWithSchedule(1, 1, startDate)
	assert.Nil(t, loan)
	assert.Error(t, err)
}

func TestCreateLoanWithSchedule_CreateLoanFailed(t *testing.T) {
	ctrl, svc, mockCU, mockLPU, mockLU, _, _ := setupLoanBillingTest(t)
	defer ctrl.Finish()

	startDate := "2025-08-01"

	customer := &model.Customer{ID: 1}
	product := &model.LoanProduct{
		ID:                 1,
		PrincipalAmount:    1000,
		InterestRate:       10,
		RepaymentCount:     5,
		RepaymentFrequency: "weekly",
	}

	mockCU.EXPECT().GetCustomerByID(uint(1)).Return(customer, nil)
	mockLPU.EXPECT().GetLoanProductByID(uint(1)).Return(product, nil)
	mockLU.EXPECT().CreateLoan(gomock.Any()).Return(fmt.Errorf("create loan failed"))

	loan, err := svc.CreateLoanWithSchedule(1, 1, startDate)
	assert.Nil(t, loan)
	assert.Error(t, err)
}

func TestCreateLoanWithSchedule_GenerateScheduleFailed(t *testing.T) {
	ctrl, svc, mockCU, mockLPU, mockLU, mockRSU, _ := setupLoanBillingTest(t)
	defer ctrl.Finish()

	startDate := "2025-08-01"

	customer := &model.Customer{ID: 1}
	product := &model.LoanProduct{
		ID:                 1,
		PrincipalAmount:    1000,
		InterestRate:       10,
		RepaymentCount:     5,
		RepaymentFrequency: "weekly",
	}

	mockCU.EXPECT().GetCustomerByID(uint(1)).Return(customer, nil)
	mockLPU.EXPECT().GetLoanProductByID(uint(1)).Return(product, nil)
	mockLU.EXPECT().CreateLoan(gomock.Any()).Return(nil)
	mockRSU.EXPECT().GenerateSchedule(gomock.Any()).Return(fmt.Errorf("generate schedule failed"))

	loan, err := svc.CreateLoanWithSchedule(1, 1, startDate)
	assert.Nil(t, loan)
	assert.Error(t, err)
}

func TestMakeScheduledPayment_Success(t *testing.T) {
	ctrl, svc, _, _, _, mockRSU, mockPU := setupLoanBillingTest(t)
	defer ctrl.Finish()

	schedule := &model.RepaymentSchedule{ID: 1, LoanID: 2, Amount: 500}

	mockRSU.EXPECT().GetNextUnpaidSchedule(uint(2)).Return(schedule, nil)
	mockPU.EXPECT().MakePayment(gomock.Any()).Return(nil)
	mockRSU.EXPECT().MarkAsPaid(uint(1), gomock.Any()).Return(nil)

	err := svc.MakeScheduledPayment(2)
	assert.NoError(t, err)
}

func TestGetLoanOutstanding(t *testing.T) {
	ctrl, svc, _, _, mockLU, _, _ := setupLoanBillingTest(t)
	defer ctrl.Finish()

	mockLU.EXPECT().GetOutstandingAmount(uint(1)).Return(int64(200), nil)

	amount, err := svc.GetLoanOutstanding(1)
	assert.NoError(t, err)
	assert.Equal(t, int64(200), amount)
}

func TestCheckDelinquency(t *testing.T) {
	ctrl, svc, _, _, mockLU, _, _ := setupLoanBillingTest(t)
	defer ctrl.Finish()

	mockLU.EXPECT().IsDelinquent(uint(1)).Return(true, nil)

	isDelinquent, err := svc.CheckDelinquency(1)
	assert.NoError(t, err)
	assert.True(t, isDelinquent)
}

func TestGetCustomerByID_Success(t *testing.T) {
	ctrl, svc, mockCU, _, mockLU, mockRPU, _ := setupLoanBillingTest(t)
	defer ctrl.Finish()

	customer := &model.Customer{ID: 1}
	loans := []model.Loan{{ID: 1}}

	mockCU.EXPECT().GetCustomerByID(uint(1)).Return(customer, nil)
	mockLU.EXPECT().ListLoansByCustomerID(uint(1), dto.CommonFilter{}).Return(loans, nil)
	mockRPU.EXPECT().IsDelinquent(uint(1)).Return(false, nil)
	mockRPU.EXPECT().GetTotalOutstandingAmount(uint(1)).Return(int64(10), nil)

	result, err := svc.GetCustomerByID(1, dto.CommonFilter{})
	assert.NoError(t, err)

	expected := mapper.ToCustomerWithLoansDTO(customer, loans, false, 10)
	assert.Equal(t, expected, result)
}

func TestGetCustomerByID_FailCustomerNotFound(t *testing.T) {
	ctrl, svc, mockCU, _, _, _, _ := setupLoanBillingTest(t)
	defer ctrl.Finish()

	mockCU.EXPECT().GetCustomerByID(uint(1)).Return(nil, errors.New("not found"))

	res, err := svc.GetCustomerByID(1, dto.CommonFilter{})
	assert.Nil(t, res)
	assert.Error(t, err)
}
