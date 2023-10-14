package handler

import (
	"Finance/payload"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handler) registerHandler(r *gin.Engine) {
	baseEndpoints := r.Group("/api")

	baseEndpoints.POST("/customer", h.handleRegisterCustomer)
	baseEndpoints.GET("/customer", h.handleGetAllCustomer)
	baseEndpoints.PUT("/customer", h.handleUpdateCustomer)
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

	response, err := h.CustomerService.UpdateCustomer(request, request.CustomerID)
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
}
