package repository

import "Finance/entity"

type CustomerRepository interface {
	RegisterCustomer(customer *entity.Customer) (*entity.Customer, error)
}
