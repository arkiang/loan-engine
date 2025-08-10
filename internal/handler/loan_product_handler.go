package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"loan-engine/internal/dto"
	"loan-engine/internal/model"
	"loan-engine/internal/usecase"
	"net/http"
	"strconv"
)

type LoanProductHandler struct {
	Usecase usecase.LoanProductUsecase
}

func NewLoanProductHandler(usecase usecase.LoanProductUsecase) *LoanProductHandler {
	return &LoanProductHandler{Usecase: usecase}
}

func (h *LoanProductHandler) CreateLoanProduct(c *gin.Context) {
	var input dto.CreateLoanProductRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	req := &model.LoanProduct{
		Name:               input.Name,
		Description:        input.Description,
		PrincipalAmount:    input.PrincipalAmount,
		InterestRate:       input.InterestRate,
		RepaymentCount:     input.RepaymentCount,
		RepaymentFrequency: input.RepaymentFrequency,
		IsActive:           input.IsActive != nil && *input.IsActive,
	}

	err := h.Usecase.CreateLoanProduct(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create loan product"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"result": "Successfully created loan product"})
}

func (h *LoanProductHandler) ListLoanProducts(c *gin.Context) {
	var filter dto.CommonFilter
	if err := c.ShouldBindQuery(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameters"})
		return
	}

	products, err := h.Usecase.ListActiveLoanProducts(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list loan products"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":      products,
		"page":      filter.Page,
		"page_size": filter.PageSize,
	})
}

func (h *LoanProductHandler) GetLoanProductByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid loan product ID"})
		return
	}

	product, err := h.Usecase.GetLoanProductByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, "Loan product not found")
		}

		c.JSON(http.StatusNotFound, gin.H{"error": "Loan product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}
