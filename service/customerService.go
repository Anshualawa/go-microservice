package service

import "github.com/Anshualawa/go-microservice/domain"

// TODO - Create CustomerService Interface
// defined GetAllCustomers - which returns list of customers.

type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
}

// TODO - Create defaultCustomerService Struct has dependency of domain.custormerRepository interface

type DefaultCustomerService struct {
	repository domain.CustomerRepository
}

// implement GetAllCustomers

func (s DefaultCustomerService) GetAllCustomers() ([]domain.Customer, error) {
	return s.repository.FindAll()
}

// TODO create Helper function NewCustomerService
// To create instance of DefaultCustomer service and inject dependency
func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
