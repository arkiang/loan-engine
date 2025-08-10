package dto

import "time"

type PaymentDTO struct {
	ID     uint      `json:"id"`
	Amount int64     `json:"amount"`
	PaidAt time.Time `json:"paid_at"`
	Method string    `json:"method"`
}
