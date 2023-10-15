package module

import (
	"Finance/api/repository"
	repositoryImpl "Finance/impl/repository"
	"database/sql"
)

type dataModuleImpl struct {
	customerRepository  repository.CustomerRepository
	limitLoanRepository repository.LimitLoanRepository
}

func NewDataModuleImpl(db *sql.DB) *dataModuleImpl {
	return &dataModuleImpl{
		customerRepository:  repositoryImpl.NewCustomerRepositoryImpl(db),
		limitLoanRepository: repositoryImpl.NewLimitLoanRepositoryImpl(db),
	}
}

func (d *dataModuleImpl) GetCustomerRepository() repository.CustomerRepository {
	return d.customerRepository
}

func (d *dataModuleImpl) GetLimitLoanRepository() repository.LimitLoanRepository {
	return d.limitLoanRepository
}
