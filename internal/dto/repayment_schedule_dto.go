package dto

import "time"

type RepaymentScheduleDTO struct {
	ID      uint      `json:"id"`
	DueDate time.Time `json:"due_date"`
	Amount  int64     `json:"amount"`
	IsPaid  bool      `json:"is_paid"`
}
