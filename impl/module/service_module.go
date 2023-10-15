package module

import (
	"Finance/api/module"
	"Finance/api/service"
	serviceImpl "Finance/impl/service"
)

type serviceModuleImpl struct {
	customerService  service.CustomerService
	limitLoanService service.LimitLoanService
}

func NewServiceModuleImpl(dataModule module.DataModule) *serviceModuleImpl {
	return &serviceModuleImpl{
		customerService: serviceImpl.NewCustomerServiceImpl(
			dataModule.GetCustomerRepository(),
		),
		limitLoanService: serviceImpl.NewLimitLoanServiceImpl(
			dataModule.GetLimitLoanRepository(),
		),
	}
}

func (s *serviceModuleImpl) GetCustomerService() service.CustomerService {
	return s.customerService
}

func (s *serviceModuleImpl) GetLimitLoanService() service.LimitLoanService {
	return s.limitLoanService
}
