package model

import "time"

type RepaymentSchedule struct {
	ID        uint      `gorm:"primaryKey"`
	LoanID    uint      `gorm:"not null;uniqueIndex:idx_loan_sequence"`
	Sequence  int       `gorm:"not null;uniqueIndex:idx_loan_sequence"` // 1..N
	DueDate   time.Time `gorm:"not null"`
	Amount    int64     `gorm:"not null"`
	IsPaid    bool      `gorm:"default:false"`
	PaidAt    *time.Time
	CreatedAt time.Time `gorm:"not null;autoCreateTime"`
}
