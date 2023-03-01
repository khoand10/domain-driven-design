package db

import (
	"context"
	"domain-driven-design/config"
	"domain-driven-design/domain/model/entity"
	"domain-driven-design/domain/repository"
	"domain-driven-design/infrastructure/persistence/datastore"
	"domain-driven-design/pkg/utils"
	"github.com/jmoiron/sqlx"
	"io/ioutil"
	"path/filepath"
)

func Connect(cfg *config.Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("sqlite3", cfg.SqlitePath)
	if err != nil {
		return nil, err
	}

	// migrate tables
	err = migrateTable(db)
	if err != nil {
		return nil, err
	}

	// migrate data
	err = migrateData(db, cfg)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func migrateTable(db *sqlx.DB) error {
	path, err := filepath.Abs("../db/migrations/migrate_tables.sql")
	if err != nil {
		return err
	}

	c, ioErr := ioutil.ReadFile(path)
	if ioErr != nil {
		return ioErr
	}
	sql := string(c)
	db.MustExec(sql)
	return nil
}

func migrateData(db *sqlx.DB, cfg *config.Config) error {
	err := migrateCustomer(db)
	err = migrateUser(db)
	return err
}

func migrateCustomer(db *sqlx.DB) error {
	path, err := filepath.Abs("../db/migrations/migrate_customer.sql")
	if err != nil {
		return err
	}

	c, ioErr := ioutil.ReadFile(path)
	if ioErr != nil {
		return ioErr
	}
	sql := string(c)
	db.MustExec(sql)
	return nil
}

func migrateUser(db *sqlx.DB) error {
	defaultPassword := "123456"
	hashedPassword := utils.HashPassword(defaultPassword)
	rootUser := &entity.User{
		Name:     "admin",
		Email:    "admin@gmail.com",
		Active:   true,
		Password: hashedPassword,
	}

	userDatastore := datastore.NewUserDatastore(db)
	userRepo := repository.NewUserRepository(userDatastore)
	ctx := context.Background()
	_, err := userRepo.Create(ctx, rootUser)
	if err != nil {
		return err
	}
	return nil
}
