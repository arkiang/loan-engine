package migration

import (
	"gorm.io/gorm"
	"loan-engine/internal/model"
	"log"
)

func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&model.Customer{},
		&model.LoanProduct{},
		&model.Loan{},
		&model.RepaymentSchedule{},
		&model.Payment{},
	)

	if err != nil {
		log.Fatalf("Failed to run AutoMigrate: %v", err)
	}

	log.Println("Database migration completed.")
}
