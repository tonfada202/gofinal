package http

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tonfada202/gofinal/models"
	repository "github.com/tonfada202/gofinal/repository"
)

type Handler struct {
	DB *sql.DB
}

func (h *Handler) GetAllCusHandler(c *gin.Context) {
	hRepo := repository.Handler{
		DB: h.DB,
	}
	customers, err := hRepo.QueryGetAllCustomer()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, customers)
}

func (h *Handler) GetCusByIdHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hRepo := repository.Handler{
		DB: h.DB,
	}
	customer, err := hRepo.QueryGetCustomerById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, customer)
}

func (h *Handler) SaveCusHandler(c *gin.Context) {
	cusRq := models.CustomerModel{}
	hRepo := repository.Handler{
		DB: h.DB,
	}
	if err := c.ShouldBindJSON(&cusRq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	customer, err := hRepo.QuerySaveCustomer(cusRq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, customer)

}

func (h *Handler) UpdateCusByIdHandler(c *gin.Context) {
	cusRq := models.CustomerModel{}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hRepo := repository.Handler{
		DB: h.DB,
	}
	if err := c.ShouldBindJSON(&cusRq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	customer, err := hRepo.QueryUpdateCustomerById(id, cusRq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, customer)
}

func (h *Handler) DelCusByIdHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hRepo := repository.Handler{
		DB: h.DB,
	}
	err = hRepo.QueryDelCustomerById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "customer deleted"})
}
