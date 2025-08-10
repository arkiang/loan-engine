package model

import "time"

type Customer struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"type:varchar(100);not null"`
	Email     string    `gorm:"type:varchar(100);unique"`
	Phone     string    `gorm:"type:varchar(20);unique"`
	CreatedAt time.Time `gorm:"not null;autoCreateTime"`
}
