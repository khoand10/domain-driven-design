package datastore

import (
	"context"
	"domain-driven-design/domain/model/entity"
	"errors"
	"github.com/jmoiron/sqlx"
)

type UserDatastore interface {
	GetByID(ctx context.Context, id int64) (*entity.User, error)
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
	Create(ctx context.Context, user *entity.User) (*entity.User, error)
}

type userDatastore struct {
	DB *sqlx.DB
}

func NewUserDatastore(db *sqlx.DB) UserDatastore {
	return &userDatastore{
		DB: db,
	}
}

func (u *userDatastore) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
	result, err := u.DB.Exec("INSERT INTO users (name, email, password, active) VALUES ($1, $2, $3, $4) RETURNING id", user.Name, user.Email, user.Password, user.Active)
	if err != nil {
		return nil, err
	}
	lastID, err := result.LastInsertId()
	if err != nil {
		return nil, errors.New("Create error")
	}
	data, err := u.GetByID(ctx, lastID)
	return data, nil
}

func (u *userDatastore) GetByID(ctx context.Context, id int64) (*entity.User, error) {
	var user entity.User
	err := u.DB.GetContext(ctx, &user, `
		SELECT id, name, email, password, active
		FROM users
		where id = $1`,
		id)
	if err != nil {
		return nil, errors.New("failed to retrieve customer by id")
	}
	return &user, nil
}

func (u *userDatastore) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User
	err := u.DB.GetContext(ctx, &user, `
		SELECT id, name, email, password, active
		FROM users
		where email = $1`,
		email)
	if err != nil {
		return nil, errors.New("failed to retrieve customer by id")
	}
	return &user, nil
}
