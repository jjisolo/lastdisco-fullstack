package pgstorage

import (
	"database/sql"
	"fmt"
	"github.com/jjisolo/lastdisco-backend/config"

	_ "github.com/lib/pq"
)

// PostgresStorage is a structure that holds all necessary
// datastructures that are needed for correct use of the
// PostgresSQL database.
type PostgresStorage struct {
	db *sql.DB
}

// NewPostgresStorage creates the storage with PostgresSQL database as a backend driver.
func NewPostgresStorage() (*PostgresStorage, error) {
	connStr := fmt.Sprintf(
		"user=%s dbname=%s password=%s sslmode=disable",
		config.PG_USER,
		config.PG_NAME,
		config.PG_PASS,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStorage{
		db: db,
	}, nil
}

// PostgresStorage is responsible for initializing the database container.
func (s *PostgresStorage) Initialize() error {
	err := s.createCategoryTable()
	if err != nil {
		return err
	}

	err = s.createProductTable()
	if err != nil {
		return err
	}

	err = s.createUserTable()
	if err != nil {
		return err
	}

	return nil
}
