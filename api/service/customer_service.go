package service

import "Finance/payload"

type CustomerService interface {
	RegisterCustomer(request payload.CustomerRequest) (*payload.CustomerResponse, error)
}
