package handler

import (
	"Finance/payload"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) registerHandler(r *gin.Engine) {
	baseEndpoints := r.Group("/api")

	baseEndpoints.POST("/customer", h.handleRegister)
}

func (h *handler) handleRegister(c *gin.Context) {
	request := payload.CustomerRequest{}

	if err := c.Bind(&request); err != nil {
		c.JSON(http.StatusBadRequest, struct {
			Message string `json:"message"`
		}{Message: err.Error()})
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
