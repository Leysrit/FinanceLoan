package handler

import (
	"Finance/api/module"
	"Finance/api/service"

	"github.com/gin-gonic/gin"
)

type handler struct {
	CustomerService    service.CustomerService
	LimitLoanService   service.LimitLoanService
	TransactionService service.TransactionService
}

func StartHandler(serviceModule module.ServiceModule) {
	handler := handler{
		CustomerService:    serviceModule.GetCustomerService(),
		LimitLoanService:   serviceModule.GetLimitLoanService(),
		TransactionService: serviceModule.GetTransactionService(),
	}

	router := gin.Default()

	handler.registerHandler(router)

	router.Run(":8080")
}
