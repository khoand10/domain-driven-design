package db

import (
	"github.com/jmoiron/sqlx"
	"io/ioutil"
	"path/filepath"
)

func Connect(path string) (*sqlx.DB, error) {
	db, err := sqlx.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	err = migrateTable(db)
	if err != nil {
		return nil, err
	}

	err = migrateData(db)
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

func migrateData(db *sqlx.DB) error {
	return migrateCustomer(db)
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
