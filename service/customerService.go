package service

import "GO_Practise/banking/domain"

type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, error)
	GetCustomer(string) (*domain.Customer, error)
}

type DefaultCustomerService struct {
	repo domain.CustomerReposity
}

func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer, error) {
	return s.repo.ById(id)
}

func NewCustomerService(repository domain.CustomerReposity) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
