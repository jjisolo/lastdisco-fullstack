package storage

import (
	"database/sql"
	"fmt"
	"log"
	"github.com/jjisolo/lastdisco-backend/deffs"
	"github.com/jjisolo/lastdisco-backend/types"

	_ "github.com/lib/pq"
)

type PostgresStorage struct {
	db *sql.DB
}

// NewPostgresStorage function is reponsible for creating the storage
// with PostgreSQL database as a backend driver.
func NewPostgresStorage() (*PostgresStorage, error) {
	connStr := fmt.Sprintf(
		"user=%s dbname=%s password=%s sslmode=disable",
		deffs.PG_USER,
		deffs.PG_NAME,
		deffs.PG_PASS,
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

// PostgresStorage is responsible for initializing
// the database container.
func (s *PostgresStorage) Initialize() error {
	s.createProductTable()
	return nil
}

// createProductTable function creates the product table 
// in the database.
func (s *PostgresStorage) createProductTable() error {
	if deffs.TESTING {
		log.Printf("DROP products table")

		_, err := s.db.Exec("DROP TABLE product")
		if err != nil {
			return err
		}
	}

	log.Printf("CREATE products table")

	query := `CREATE TABLE IF NOT EXISTS product(
    id          SERIAL PRIMARY KEY,
    short_name  VARCHAR(75),
    description VARCHAR(1300),
    count       INT,
    price       VARCHAR(12),
    created_at  TIMESTAMP,
    updated_at  TIMESTAMP
	)`

	_, err := s.db.Exec(string(query))
	return err
}

// createUserTable function creates the user table in
// the database.
func (s *PostgresStorage) createUserTable() error {
	if deffs.TESTING {
		log.Printf("DROP users table")

		_, err := s.db.Exec("DROP TABLE user")
		if err != nil {
			return err
		}
	}

	query := `CREATE TABLE IF NOT EXISTS user(
	id          SERIAL PRIMARY KEY
	first_name  VARCHAR(75)
	last_name   VARCHAR(75)
	created_at  TIMESTAMP,
	updated_at  TIMESTAMP,
	)`

	_, err := s.db.Exec(string(query))
	return err
}

// CreateProduct function is responsible for creating the single
// product entry in the database.
func (s *PostgresStorage) CreateProduct(product *types.Product) error {
	log.Printf("CREATE product with id of %s", product.ID)

	query := `INSERT INTO product
	(short_name, description, count, price, created_at, updated_at)
	VALUES
	($1, $2, $3, $4, $5, $6)`

	_, err := s.db.Query(
		string(query), 
		product.ShortName,
		product.Description,
		product.Count,
		product.Price,
		product.CreatedAt,
		product.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

// CreateUser function is responsible for creating the single
// user entity in the database.
func (s *PostgresStorage) CreateUser(user *types.User) error {
	log.Printf("CREATE user")

	query := `INSERT INTO user
	(first_name, last_name, created_at, updated_at)
	VALUES
	($1, $2, $3, $4)`

	_, err := s.db.Query(
		string(query),
		user.FirstName,
		user.LastName,
		user.CreatedAt,
		user.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

// UpdateProduct function is reponsible for updating the existing(in the database)
// product.
func (s *PostgresStorage) UpdateProduct(product *types.Product, id int) error {
	log.Printf("UPDATE product with id of %d", product.ID)

	query := `UPDATE product
	SET
		short_name = $1,
		description = $2,
		count = $3,
		price = $4,
		updated_at = NOW() 
	WHERE id = $5;
	`

	_, err := s.db.Query(
		string(query),
		product.ShortName,
		product.Description,
		product.Count,
		product.Price,
		id,
	)
	return err
}

// UpdateUser function is responsible for updating the existing
// user entity with the provided data.
func (s *PostgresStorage) UpdateUser(user *types.User, id int) error {
	log.Printf("UPDATE user with id of %d", user.ID)

	query := `UPDATE user 
	SET
		first_name = $1,
		last_name = $2,
		updated_at = NOW() 
	WHERE id = $3;
	`

	_, err := s.db.Query(
		string(query),
		user.FirstName,
		user.LastName,
		id,
	)
	return err
}

// DeleteProduct function is reponsible for existing(in the database) product
// deletion.
func (s *PostgresStorage) DeleteProduct(id int) error {
	log.Printf("Request to DELETE product with id of %d", id)
	query := "DELETE FROM product WHERE id = $1"

	_, err := s.db.Query(string(query), id)
	return err
}

// DeleteUser function is reponsible for existing(in the database) product
// deletion.
func (s *PostgresStorage) DeleteUser(id int) error {
	log.Printf("Request to DELETE user with id of %d", id)

	query := "DELETE FROM user WHERE id = $1"

	_, err := s.db.Query(string(query), id)
	return err
}

// GetProductByID function is reponsible for getting individual product from
// the database using its unique identifier.
func (s *PostgresStorage) GetProductByID(id int) (*types.Product, error) {
	query := "SELECT * FROM product WHERE id = $1"

	rows, err := s.db.Query(string(query), id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return readProductFromQuery(rows)
	}

	return nil, fmt.Errorf("product %d does not exist", id)
}

// GetUserByID function is reponsible for getting individual product from
// the database using its unique identifier.
func (s *PostgresStorage) GetUserByID(id int) (*types.User, error) {
	query := "SELECT * FROM user WHERE id = $1"

	rows, err := s.db.Query(string(query), id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return readUserFromQuery(rows)
	}

	return nil, fmt.Errorf("user %d does not exist", id)
}

// GetProducts function is responsible for retrieving all products,
// that are located in the internal database.
func (s *PostgresStorage) GetProducts() ([]*types.Product, error) {
	query := "SELECT * FROM product"

	rows, err := s.db.Query(string(query))
	if err != nil {
		return nil, err
	}

	products := []*types.Product{}
	for rows.Next() {
		product, err := readProductFromQuery(rows)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

// GetUsers function is responsible for retrieving all users,
// that are located in the internal database.
func (s *PostgresStorage) GetUsers() ([]*types.User, error) {
	query := "SELECT * FROM user"

	rows, err := s.db.Query(string(query))
	if err != nil {
		return nil, err
	}

	users := []*types.User{}
	for rows.Next() {
		user, err := readUserFromQuery(rows)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// readProductFromQuery function is reponsible for scanning the database
// query response to the product datatype.
func readProductFromQuery(rows *sql.Rows) (*types.Product, error) {
	product := new(types.Product)
	err := rows.Scan(
		&product.ID,
		&product.ShortName,
		&product.Description,
		&product.Count,
		&product.Price,
		&product.CreatedAt, 
		&product.UpdatedAt)

	return product, err
}

// readUserFromQuery function is reponsible for scanning the database
// query response to the user datatype.
func readUserFromQuery(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)
	err  := rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.CreatedAt, 
		&user.UpdatedAt)

	return user, err
}