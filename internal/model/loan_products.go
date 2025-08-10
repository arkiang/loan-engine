package model

import "time"

type LoanProduct struct {
	ID                 uint      `gorm:"primaryKey"`
	Name               string    `gorm:"type:varchar(100);not null"`
	Description        string    `gorm:"type:text"`
	PrincipalAmount    int64     `gorm:"not null"`
	InterestRate       float64   `gorm:"not null"`
	RepaymentCount     int       `gorm:"not null"`
	RepaymentFrequency string    `gorm:"type:varchar(20);not null;check:repayment_frequency IN ('daily','weekly','monthly')"`
	IsActive           bool      `gorm:"default:true"`
	CreatedAt          time.Time `gorm:"not null;autoCreateTime"`
}
