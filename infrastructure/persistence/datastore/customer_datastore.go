package datastore

import (
	"context"
	"domain-driven-design/domain/model/entity"
	"errors"
	"github.com/jmoiron/sqlx"
)

type CustomerDatastore interface {
	GetByID(ctx context.Context, id int64) (*entity.Customer, error)
	GetAll(ctx context.Context) ([]*entity.Customer, error)
	GetByEmail(ctx context.Context, email string) (*entity.Customer, error)
	Create(ctx context.Context, customer *entity.Customer) (*entity.Customer, error)
}

type customerDatastore struct {
	DB *sqlx.DB
}

func NewCustomerDatastore(db *sqlx.DB) CustomerDatastore {
	return &customerDatastore{
		DB: db,
	}
}

func (c *customerDatastore) GetAll(ctx context.Context) ([]*entity.Customer, error) {
	var customers []*entity.Customer

	err := c.DB.Select(&customers, `
		SELECT id, name, email, address
		FROM customers`)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, errors.New("failed to retrieve customer by id")
	}
	return customers, nil
}

func (c *customerDatastore) GetByID(ctx context.Context, id int64) (*entity.Customer, error) {
	var customer entity.Customer
	err := c.DB.GetContext(ctx, &customer, `
		SELECT id, name, email, address
		FROM customers
		where id = $1`,
		id)
	if err != nil {
		return nil, errors.New("failed to retrieve customer by id")
	}
	return &customer, nil
}

func (c *customerDatastore) GetByEmail(ctx context.Context, email string) (*entity.Customer, error) {
	var customer entity.Customer
	err := c.DB.GetContext(ctx, &customer, `
		SELECT id, name, email, address
		FROM customers
		where email = $1`,
		email)
	if err != nil {
		return nil, errors.New("failed to retrieve customer by email")
	}
	return &customer, nil
}

func (c *customerDatastore) Create(ctx context.Context, customer *entity.Customer) (*entity.Customer, error) {
	result, err := c.DB.Exec("INSERT INTO customers (name, email, address) VALUES ($1, $2, $3) RETURNING id", customer.Name, customer.Email, customer.Address)
	if err != nil {
		return nil, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return nil, errors.New("Create error")
	}

	data, err := c.GetByID(ctx, lastID)

	return data, nil
}
