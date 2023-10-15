package service

import "Finance/payload"

type CustomerService interface {
	RegisterCustomer(request payload.CustomerRequest) (*payload.CustomerResponse, error)
	GetAllCustomer(request payload.ListCustomerRequest) (*payload.ListCustomerResponse, error)
	UpdateCustomer(request payload.UpdateCustomer, customerID int) (*payload.CustomerResponse, error)
	Login(request payload.LoginRequest) (*payload.LoginResponse, error)
}
