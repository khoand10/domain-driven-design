package repository

import (
	"context"
	"domain-driven-design/domain/model/entity"
	"domain-driven-design/infrastructure/persistence/datastore"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) (*entity.User, error)
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
	GetByID(ctx context.Context, id int64) (*entity.User, error)
}

type userRepository struct {
	datastore datastore.UserDatastore
}

func (u userRepository) GetByID(ctx context.Context, id int64) (*entity.User, error) {
	return u.datastore.GetByID(ctx, id)
}

func (u userRepository) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
	return u.datastore.Create(ctx, user)
}

func (u userRepository) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	return u.datastore.GetByEmail(ctx, email)
}

func NewUserRepository(datastore datastore.UserDatastore) UserRepository {
	return &userRepository{
		datastore: datastore,
	}
}
