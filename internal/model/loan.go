package model

import "time"

type Loan struct {
	ID                 uint      `gorm:"primaryKey"`
	CustomerID         uint      `gorm:"not null"`
	LoanProductID      uint      `gorm:"not null"`
	Principal          int64     `gorm:"not null"`
	InterestRate       float64   `gorm:"not null"`
	TotalAmount        int64     `gorm:"not null"`
	RepaymentCount     int       `gorm:"not null"`
	RepaymentFrequency string    `gorm:"type:varchar(20);not null;check:repayment_frequency IN ('daily','weekly','monthly')"`
	StartDate          time.Time `gorm:"type:date;not null;index"`
	CreatedAt          time.Time `gorm:"not null;autoCreateTime"`

	Product   LoanProduct         `gorm:"foreignKey:LoanProductID"`
	Schedules []RepaymentSchedule `gorm:"foreignKey:LoanID"`
	Payments  []Payment           `gorm:"foreignKey:LoanID"`
}
