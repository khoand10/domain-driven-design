package service

import (
	"context"
	"domain-driven-design/domain/model/entity"
	"domain-driven-design/domain/repository"
)

type CustomerService interface {
	GetCustomer(id int64) (*entity.Customer, error)
	GetCustomers() ([]*entity.Customer, error)
	CreateCustomer(customer *entity.Customer) (*entity.Customer, error)
	UpdateCustomer(customer *entity.Customer) error
	DeleteCustomer(id int64) error
}

type customerService struct {
	customerRepository repository.CustomerRepository
}

func (c customerService) GetCustomers() ([]*entity.Customer, error) {
	ctx := context.Background()
	return c.customerRepository.FindAll(ctx)
}

func (c customerService) GetCustomer(id int64) (*entity.Customer, error) {
	ctx := context.Background()
	return c.customerRepository.FindById(ctx, id)
}

func (c customerService) CreateCustomer(customer *entity.Customer) (*entity.Customer, error) {
	ctx := context.Background()
	data, err := c.customerRepository.Create(ctx, customer)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c customerService) UpdateCustomer(customer *entity.Customer) error {
	ctx := context.Background()
	return c.customerRepository.Update(ctx, customer)
}

func (c customerService) DeleteCustomer(id int64) error {
	ctx := context.Background()
	return c.customerRepository.Delete(ctx, id)
}

func NewCustomerService(repo repository.CustomerRepository) CustomerService {
	return &customerService{
		customerRepository: repo,
	}
}
