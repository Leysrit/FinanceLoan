package module

import (
	"Finance/api/module"
	"Finance/api/service"
	serviceImpl "Finance/impl/service"
)

type serviceModuleImpl struct {
	customerService    service.CustomerService
	limitLoanService   service.LimitLoanService
	transactionService service.TransactionService
}

func NewServiceModuleImpl(dataModule module.DataModule) *serviceModuleImpl {
	return &serviceModuleImpl{
		customerService: serviceImpl.NewCustomerServiceImpl(
			dataModule.GetCustomerRepository(),
		),
		limitLoanService: serviceImpl.NewLimitLoanServiceImpl(
			dataModule.GetLimitLoanRepository(),
		),
		transactionService: serviceImpl.NewTransactionServiceImpl(
			dataModule.GetTransactionRepository(),
		),
	}
}

func (s *serviceModuleImpl) GetCustomerService() service.CustomerService {
	return s.customerService
}

func (s *serviceModuleImpl) GetLimitLoanService() service.LimitLoanService {
	return s.limitLoanService
}

func (s *serviceModuleImpl) GetTransactionService() service.TransactionService {
	return s.transactionService
}
