package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"loan-engine/internal/dto"
	"loan-engine/internal/service"
	"net/http"
	"strconv"
)

type LoanHandler struct {
	Service service.LoanBillingService
}

func NewLoanHandler(s service.LoanBillingService) *LoanHandler {
	return &LoanHandler{Service: s}
}

func (h *LoanHandler) CreateLoan(c *gin.Context) {
	var req dto.CreateLoanRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	loan, err := h.Service.CreateLoanWithSchedule(req.CustomerID, req.ProductID, req.StartDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := dto.LoanResponse{
		LoanID:         loan.ID,
		CustomerID:     loan.CustomerID,
		ProductID:      loan.LoanProductID,
		Principal:      loan.Principal,
		TotalAmount:    loan.TotalAmount,
		RepaymentCount: loan.RepaymentCount,
		Frequency:      loan.RepaymentFrequency,
		StartDate:      loan.StartDate.Format("01/02/2006"),
	}
	c.JSON(http.StatusCreated, resp)
}

// MakeScheduledPayment an api to make a payment without concern about due date (for testing)
func (h *LoanHandler) MakeScheduledPayment(c *gin.Context) {
	loanID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid loan id"})
		return
	}

	err = h.Service.MakeScheduledPayment(uint(loanID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": "payment successful"})
}

// MakePayment an api to make a payment with concern about due date
func (h *LoanHandler) MakePayment(c *gin.Context) {
	loanID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid loan id"})
		return
	}

	err = h.Service.MakePayment(uint(loanID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": "payment successful"})
}

func (h *LoanHandler) GetOutstanding(c *gin.Context) {
	loanID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid loan id"})
		return
	}

	amount, err := h.Service.GetLoanOutstanding(uint(loanID))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, "loan not found")
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	resp := dto.OutstandingResponse{
		LoanID:      uint(loanID),
		Outstanding: amount,
	}
	c.JSON(http.StatusOK, resp)
}

func (h *LoanHandler) CheckDelinquency(c *gin.Context) {
	loanID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid loan id"})
		return
	}

	isDelinquent, err := h.Service.CheckDelinquency(uint(loanID))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, "loan not found")
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	resp := dto.DelinquencyResponse{
		LoanID:       uint(loanID),
		IsDelinquent: isDelinquent,
	}
	c.JSON(http.StatusOK, resp)
}
