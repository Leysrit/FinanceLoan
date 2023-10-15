package module

import (
	"Finance/api/repository"
)

type DataModule interface {
	GetCustomerRepository() repository.CustomerRepository
	GetLimitLoanRepository() repository.LimitLoanRepository
	GetTransactionRepository() repository.TransactionRepository
}
