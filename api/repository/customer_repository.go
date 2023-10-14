package repository

import "Finance/entity"

type CustomerRepository interface {
	RegisterCustomer(customer *entity.Customer) (*entity.Customer, error)
	GetAllCustomer(page int, limit int, nama string) ([]*entity.Customer, error)
	GetTotalCustomer(nama string) (int, error)
}
