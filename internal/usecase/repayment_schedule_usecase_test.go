package usecase

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"loan-engine/internal/model"
	mocks "loan-engine/mocks/repository"
	"testing"
	"time"
)

func TestGenerateSchedule(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepaymentScheduleRepository(ctrl)
	uc := NewRepaymentScheduleUsecase(mockRepo)

	loan := &model.Loan{
		ID:                 1,
		TotalAmount:        1000,
		RepaymentCount:     4,
		RepaymentFrequency: "weekly",
		StartDate:          time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	// Expect BulkCreate to be called with the correct schedule count
	mockRepo.EXPECT().BulkCreate(gomock.Len(4)).Return(nil)

	err := uc.GenerateSchedule(loan)
	assert.NoError(t, err)
}

func TestGetSchedulesByLoanID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepaymentScheduleRepository(ctrl)
	uc := NewRepaymentScheduleUsecase(mockRepo)

	expected := []model.RepaymentSchedule{
		{LoanID: 1, Sequence: 1},
	}

	mockRepo.EXPECT().FindByLoanID(uint(1)).Return(expected, nil)

	schedules, err := uc.GetSchedulesByLoanID(1)
	assert.NoError(t, err)
	assert.Equal(t, expected, schedules)
}

func TestGetNextUnpaidSchedule(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepaymentScheduleRepository(ctrl)
	uc := NewRepaymentScheduleUsecase(mockRepo)

	t.Run("success", func(t *testing.T) {
		expected := []model.RepaymentSchedule{{LoanID: 1, Sequence: 1}}
		mockRepo.EXPECT().GetUnpaidSchedules(uint(1)).Return(expected, nil)

		sch, err := uc.GetNextUnpaidSchedule(1)
		assert.NoError(t, err)
		assert.Equal(t, &expected[0], sch)
	})

	t.Run("repo error", func(t *testing.T) {
		mockRepo.EXPECT().GetUnpaidSchedules(uint(1)).Return(nil, errors.New("db error"))

		sch, err := uc.GetNextUnpaidSchedule(1)
		assert.Nil(t, sch)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "get next unpaid schedule failed")
	})

	t.Run("empty schedule", func(t *testing.T) {
		mockRepo.EXPECT().GetUnpaidSchedules(uint(1)).Return([]model.RepaymentSchedule{}, nil)

		sch, err := uc.GetNextUnpaidSchedule(1)
		assert.Nil(t, sch)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "active loan not found")
	})
}

func TestMarkAsPaid(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepaymentScheduleRepository(ctrl)
	uc := NewRepaymentScheduleUsecase(mockRepo)

	mockRepo.EXPECT().MarkAsPaid(uint(1), gomock.Any()).Return(nil)

	err := uc.MarkAsPaid(1, time.Now())
	assert.NoError(t, err)
}
