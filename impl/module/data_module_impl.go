package module

import (
	"Finance/api/repository"
	repositoryImpl "Finance/impl/repository"
	"database/sql"
)

type dataModuleImpl struct {
	customerRepository repository.CustomerRepository
}

func NewDataModuleImpl(db *sql.DB) *dataModuleImpl {
	return &dataModuleImpl{
		customerRepository: repositoryImpl.NewCustomerRepositoryImpl(db),
	}
}

func (d *dataModuleImpl) GetCustomerRepository() repository.CustomerRepository {
	return d.customerRepository
}
