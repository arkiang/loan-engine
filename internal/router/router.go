package router

import (
	"github.com/gin-gonic/gin"
	"loan-engine/internal/handler"
	"loan-engine/internal/middleware"
)

func RegisterRoutes(loanHandler *handler.LoanHandler,
	customerHandler *handler.CustomerHandler,
	productHandler *handler.LoanProductHandler) *gin.Engine {

	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.RecoveryMiddleware())

	r.POST("/loans", loanHandler.CreateLoan)
	r.POST("/loans/:id/pay", loanHandler.MakePayment)
	r.POST("/loans/:id/pay/test", loanHandler.MakeScheduledPayment)
	r.GET("/loans/:id/outstanding", loanHandler.GetOutstanding)
	r.GET("/loans/:id/delinquent", loanHandler.CheckDelinquency)

	r.POST("/customers", customerHandler.CreateCustomer)
	r.GET("/customers", customerHandler.ListCustomers)
	r.GET("/customers/:id", customerHandler.GetCustomerByID)

	r.POST("/loan-products", productHandler.CreateLoanProduct)
	r.GET("/loan-products", productHandler.ListLoanProducts)
	r.GET("/loan-products/:id", productHandler.GetLoanProductByID)

	return r
}
