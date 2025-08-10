package main

import (
	"loan-engine/config"
	"loan-engine/infrastructure/persistent"
	"loan-engine/internal/handler"
	"loan-engine/internal/router"
	"loan-engine/internal/service"
	"loan-engine/internal/usecase"
	"loan-engine/migration"
	"log"
)

func main() {
	// Load environment variables
	err := config.LoadConfig()
	if err != nil {
		log.Fatal("Could not load environment variables ", err)
	}

	// Connect to Postgres
	db, err := config.ConnectDB(&config.AppConfig)
	if err != nil {
		log.Fatal("Could not connect to database", err)
	}

	// Run DB migration
	migration.AutoMigrate(db)

	// Init dependencies
	customerRepo := persistent.NewCustomerRepository(db)
	loanProductRepo := persistent.NewLoanProductRepository(db)
	paymentRepo := persistent.NewPaymentRepository(db)
	repaymentScheduleRepo := persistent.NewRepaymentScheduleRepository(db)
	loanRepo := persistent.NewLoanRepository(db)

	customerUsecase := usecase.NewCustomerUsecase(customerRepo)
	loanProductUsecase := usecase.NewLoanProductUsecase(loanProductRepo)
	paymentUsecase := usecase.NewPaymentUsecase(paymentRepo)
	repaymentScheduleUsecase := usecase.NewRepaymentScheduleUsecase(repaymentScheduleRepo)
	loanUsecase := usecase.NewLoanUsecase(loanRepo)

	loanBillingService := service.NewLoanBillingService(customerUsecase, loanProductUsecase, loanUsecase, repaymentScheduleUsecase, paymentUsecase)

	customerHandler := handler.NewCustomerHandler(customerUsecase, loanBillingService)
	loanProductHandler := handler.NewLoanProductHandler(loanProductUsecase)
	loanBillingHandler := handler.NewLoanHandler(loanBillingService)

	// Register routes
	routes := router.RegisterRoutes(
		loanBillingHandler,
		customerHandler,
		loanProductHandler,
	)

	// Run server
	log.Printf("Server is running on port %s...", config.AppConfig.ServerPort)
	if err := routes.Run(":" + config.AppConfig.ServerPort); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
