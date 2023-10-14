package module

import (
	"Finance/api/module"
	"Finance/api/service"
	serviceImpl "Finance/impl/service"
)

type serviceModuleImpl struct {
	customerService service.CustomerService
}

func NewServiceModuleImpl(dataModule module.DataModule) *serviceModuleImpl {
	return &serviceModuleImpl{
		customerService: serviceImpl.NewCustomerServiceImpl(
			dataModule.GetCustomerRepository(),
		),
	}
}

func (s *serviceModuleImpl) GetCustomerService() service.CustomerService {
	return s.customerService
}
