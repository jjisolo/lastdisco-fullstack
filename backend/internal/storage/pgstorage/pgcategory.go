package pgstorage

import (
	"database/sql"
	"fmt"
	"github.com/jjisolo/lastdisco-backend/config"
	"github.com/jjisolo/lastdisco-backend/internal/types"
	"log"
)

// createCategoryTable function creates the category table in the database.
func (s *PostgresStorage) createCategoryTable() error {
	if config.TESTING {
		log.Printf("DROP product_category table")

		_, err := s.db.Exec("DROP TABLE IF EXISTS product_category CASCADE")
		if err != nil {
			return err
		}
	}
	log.Printf("CREATE products_category table")

	query := `CREATE TABLE IF NOT EXISTS product_category(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(75),
    description TEXT,
    created_at  TIMESTAMP,
    updated_at  TIMESTAMP
	)`

	_, err := s.db.Exec(string(query))
	return err
}

// CreateCategory function is responsible for creating the single category entry in the database.
func (s *PostgresStorage) CreateCategory(category *types.ProductCategory) error {
	log.Printf("CREATE product_category with id of %s", category.ID)

	query := `INSERT INTO product_category
	(name, description)
	VALUES
	($1, $2)`

	_, err := s.db.Query(
		string(query),
		category.Name,
		category.Description,
	)
	if err != nil {
		return err
	}

	return nil
}

// UpdateCategory function is responsible for updating the existing product entity with the provided data.
func (s *PostgresStorage) UpdateCategory(category *types.ProductCategory, id int) error {
	log.Printf("UPDATE category with id of %d", category.ID)

	query := `UPDATE product_category
	SET
		name = $1,
		description = $2,
		updated_at = NOW() 
	WHERE id = $3;
	`

	_, err := s.db.Query(
		string(query),
		category.Name,
		category.Description,
		id,
	)
	return err
}

// DeleteCategory function is responsible for existing(in the database) category deletion.
func (s *PostgresStorage) DeleteCategory(id int) error {
	log.Printf("Request to DELETE category with id of %d", id)
	query := "DELETE FROM product_category WHERE id = $1"

	_, err := s.db.Query(string(query), id)
	return err
}

// GetCategoryByID function is responsible for getting individual category from the database using its unique identifier.
func (s *PostgresStorage) GetCategoryByID(id int) (*types.ProductCategory, error) {
	query := "SELECT * FROM product_category WHERE id = $1"

	rows, err := s.db.Query(string(query), id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return readCategoryFromQuery(rows)
	}

	return nil, fmt.Errorf("product category %d does not exist", id)
}

// GetCategories function is responsible for retrieving all categories, that are located in the internal database.
func (s *PostgresStorage) GetCategories() ([]*types.ProductCategory, error) {
	query := "SELECT * FROM product_category"

	rows, err := s.db.Query(string(query))
	if err != nil {
		return nil, err
	}

	categories := []*types.ProductCategory{}
	for rows.Next() {
		category, err := readCategoryFromQuery(rows)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

// readCategoryFromQuery function is responsible for scanning the database query response to the product_category datatype.
func readCategoryFromQuery(rows *sql.Rows) (*types.ProductCategory, error) {
	category := new(types.ProductCategory)
	err := rows.Scan(
		&category.ID,
		&category.Description,
	)

	return category, err
}
