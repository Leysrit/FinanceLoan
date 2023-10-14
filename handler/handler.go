package handler

import (
	"Finance/api/module"
	"Finance/api/service"

	"github.com/gin-gonic/gin"
)

type handler struct {
	CustomerService service.CustomerService
}

func StartHandler(serviceModule module.ServiceModule) {
	handler := handler{
		CustomerService: serviceModule.GetCustomerService(),
	}

	router := gin.Default()

	handler.registerHandler(router)

	router.Run(":8080")
}
