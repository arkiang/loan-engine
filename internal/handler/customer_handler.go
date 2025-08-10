package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"loan-engine/internal/dto"
	"loan-engine/internal/model"
	"loan-engine/internal/service"
	"loan-engine/internal/usecase"
	"net/http"
	"strconv"
)

type CustomerHandler struct {
	customerUsecase usecase.CustomerUsecase
	service         service.LoanBillingService
}

func NewCustomerHandler(customerService usecase.CustomerUsecase, billingService service.LoanBillingService) *CustomerHandler {
	return &CustomerHandler{customerUsecase: customerService, service: billingService}
}

func (h *CustomerHandler) CreateCustomer(c *gin.Context) {
	var req dto.CreateCustomerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customer := &model.Customer{
		Name:  req.Name,
		Email: req.Email,
		Phone: req.Phone,
	}
	err := h.customerUsecase.CreateCustomer(customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create customer"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"result": "Successfully created customer"})
}

func (h *CustomerHandler) GetCustomerByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
		return
	}

	var filter dto.CommonFilter
	if err := c.ShouldBindQuery(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameters"})
		return
	}

	customer, err := h.service.GetCustomerByID(uint(id), filter)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, "Customer not found")
		}

		c.JSON(http.StatusNotFound, gin.H{"error": "Failed to find customer"})
		return
	}

	c.JSON(http.StatusOK, customer)
}

func (h *CustomerHandler) ListCustomers(c *gin.Context) {
	var filter dto.CommonFilter

	if err := c.ShouldBindQuery(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameters"})
		return
	}

	customers, err := h.customerUsecase.ListCustomers(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list customers"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":      customers,
		"page":      filter.Page,
		"page_size": filter.PageSize,
	})
}
