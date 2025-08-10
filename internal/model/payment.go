package model

import "time"

type Payment struct {
	ID            uint          `gorm:"primaryKey"`
	LoanID        uint          `gorm:"not null"`
	Amount        int64         `gorm:"not null"`
	PaidAt        time.Time     `gorm:"not null"`
	PaymentMethod string        `gorm:"type:varchar(50);default:'manual'"`
	Status        PaymentStatus `gorm:"type:varchar(20);not null;check:status IN ('success')"`
	CreatedAt     time.Time     `gorm:"not null;autoCreateTime"`
}

type PaymentStatus string

const (
	PaymentStatusSuccess PaymentStatus = "success"
)
