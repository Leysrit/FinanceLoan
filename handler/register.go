package handler

import (
	"Finance/middleware"
	"Finance/payload"
	"html"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handler) registerHandler(r *gin.Engine) {
	baseEndpoints := r.Group("/api")

	baseEndpoints.POST("/login", h.handleLogin)

	baseEndpoints.POST("/customer", h.handleRegisterCustomer)
	baseEndpoints.GET("/customer", h.handleGetAllCustomer)
	baseEndpoints.PUT("/customer/:id", h.handleUpdateCustomer)

	baseEndpoints.GET("/limit-loan/:customer_id", h.handleGetLimitLoan)
	baseEndpoints.PUT("/limit-loan/:customer_id", h.handleUpdateLimitLoan)

	baseEndpoints.POST("/transaction/:customer_id", middleware.AuthorizationMiddleware(), h.handleAddTransaction)
}

func (h *handler) handleLogin(c *gin.Context) {
	request := payload.LoginRequest{}
	if err := c.Bind(&request); err != nil {
		log.Printf("Error binding request: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"error":   "Bad Request",
		})
		return
	}

	response, err := h.CustomerService.Login(request)
	if err != nil {
		log.Printf("Error login: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"error":   "Bad Request",
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *handler) handleAddTransaction(c *gin.Context) {
	request := payload.AddTransactionRequest{}
	if err := c.Bind(&request); err != nil {
		log.Printf("Error binding request: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"error":   "Bad Request",
		})
		return
	}

	customerID := c.Param("customer_id")
	id, err := strconv.Atoi(customerID)
	if err != nil {
		log.Printf("Error converting customer ID to integer: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"error":   "Bad Request",
		})
		return
	}

	request.CustomerID = id
	log.Println("Customer ID: ", request.CustomerID)

	// Menghindari XSS dengan membersihkan data input pengguna
	cleanedAssetName := html.EscapeString(request.AssetName)
	request.AssetName = cleanedAssetName

	response, err := h.TransactionService.AddTransaction(&request)
	if err != nil {
		log.Printf("Error adding transaction: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"error":   "Bad Request",
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *handler) handleUpdateLimitLoan(c *gin.Context) {
	request := payload.UpdateLimitLoanRequest{}
	if err := c.Bind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"error":   "Bad Request",
		})
		return
	}

	customerID := c.Param("customer_id")
	id, err := strconv.Atoi(customerID)
	if err != nil {
		log.Printf("Error converting customer ID to integer: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"error":   "Bad Request",
		})
		return
	}

	request.CustomerID = id
	response, err := h.LimitLoanService.UpdateLimitLoan(&request)
	if err != nil {
		log.Printf("Error updating limit loan: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"error":   "Bad Request",
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *handler) handleGetLimitLoan(c *gin.Context) {
	customerID := c.Param("customer_id")
	id, err := strconv.Atoi(customerID)
	if err != nil {
		log.Printf("Error converting customer ID to integer: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"error":   "Bad Request",
		})
		return
	}

	if id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid customer ID",
			"error":   "Bad Request",
		})
		return
	}

	response, err := h.LimitLoanService.GetLimitLoan(id)
	if err != nil {
		log.Printf("Error retrieving limit loan: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"error":   "Bad Request",
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *handler) handleUpdateCustomer(c *gin.Context) {
	request := payload.UpdateCustomer{}
	if err := c.Bind(&request); err != nil {
		c.JSON(http.StatusBadRequest, struct {
			Message string `json:"message"`
			Error   string `json:"error"`
		}{Message: err.Error(), Error: "Bad Request"})
		return
	}

	requestID := c.Param("id")
	id, err := strconv.Atoi(requestID)
	if err != nil {
		c.JSON(http.StatusBadRequest, struct {
			Message string `json:"message"`
			Error   string `json:"error"`
		}{Message: err.Error(), Error: "Bad Request"})
		return
	}

	response, err := h.CustomerService.UpdateCustomer(request, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, struct {
			Message string `json:"message"`
		}{Message: err.Error()})
	}

	c.JSON(http.StatusOK, response)

}

func (h *handler) handleRegisterCustomer(c *gin.Context) {
	request := payload.CustomerRequest{}
	if err := c.Bind(&request); err != nil {
		c.JSON(http.StatusBadRequest, struct {
			Message string `json:"message"`
			Error   string `json:"error"`
		}{Message: err.Error(), Error: "Bad Request"})
		return
	}

	response, err := h.CustomerService.RegisterCustomer(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, struct {
			Message string `json:"message"`
		}{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
	return
}

func (h *handler) handleGetAllCustomer(c *gin.Context) {
	request := payload.ListCustomerRequest{}

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		c.JSON(http.StatusBadRequest, struct {
			Message string `json:"message"`
			Error   string `json:"error"`
		}{Message: err.Error(), Error: "Bad Request"})
		return
	}

	request.Page = page

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		c.JSON(http.StatusBadRequest, struct {
			Message string `json:"message"`
			Error   string `json:"error"`
		}{Message: err.Error(), Error: "Bad Request"})
		return
	}

	request.Limit = limit

	request.Nama = c.Query("nama")

	response, err := h.CustomerService.GetAllCustomer(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, struct {
			Message string `json:"message"`
			Error   string `json:"error"`
		}{Message: err.Error(), Error: "Bad Request"})
		return
	}

	c.JSON(http.StatusOK, response)
	return
}
