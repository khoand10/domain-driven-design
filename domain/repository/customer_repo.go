package repository

import (
	"context"
	"domain-driven-design/domain/model/entity"
	"domain-driven-design/infrastructure/persistence/datastore"
)

type CustomerRepository interface {
	Create(ctx context.Context, customer *entity.Customer) (*entity.Customer, error)
	FindById(ctx context.Context, id int64) (*entity.Customer, error)
	FindAll(ctx context.Context) ([]*entity.Customer, error)
	Update(ctx context.Context, customer *entity.Customer) error
	Delete(ctx context.Context, id int64) error
}

//var CustomerData = []*entity.Customer{
//	{
//		ID:    "1",
//		Email: "demo@gmail.com",
//		Name:  "demo user",
//	},
//	{
//		ID:    "2",
//		Email: "admin@gmail.com",
//		Name:  "admin user",
//	},
//}

type customerRepository struct {
	datastore datastore.CustomerDatastore
}

func NewCustomerRepository(datastore datastore.CustomerDatastore) CustomerRepository {
	return &customerRepository{
		datastore: datastore,
	}
}

func (cr *customerRepository) FindById(ctx context.Context, id int64) (*entity.Customer, error) {
	return cr.datastore.GetByID(ctx, id)
}

func (cr *customerRepository) FindAll(ctx context.Context) ([]*entity.Customer, error) {
	return cr.datastore.GetAll(ctx)
}

func (cr *customerRepository) Update(ctx context.Context, customer *entity.Customer) error {
	return nil
}

func (cr *customerRepository) Delete(ctx context.Context, id int64) error {
	return nil
}

func (cr *customerRepository) Create(ctx context.Context, customer *entity.Customer) (*entity.Customer, error) {
	return cr.datastore.Create(ctx, customer)
}
