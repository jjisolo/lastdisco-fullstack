package pgstorage

import (
	"database/sql"
	"fmt"
	"github.com/jjisolo/lastdisco-backend/config"
	"github.com/jjisolo/lastdisco-backend/internal/types"
	"log"
)

// createUserTable function creates the user table in the database.
func (s *PostgresStorage) createUserTable() error {
	if config.TESTING {
		log.Printf("DROP users table")

		_, err := s.db.Exec("DROP TABLE IF EXISTS users")
		if err != nil {
			return err
		}
	}

	query := `CREATE TABLE IF NOT EXISTS users(
	id                 SERIAL PRIMARY KEY,
	role               VARCHAR(30) DEFAULT 'ROLE_USER',
	first_name         VARCHAR(75),
	last_name          VARCHAR(75),
	password_encrypted VARCHAR(100),
	email              VARCHAR(100),
    phone_number       VARCHAR(20),
	created_at         TIMESTAMP,
	updated_at         TIMESTAMP
	)`

	_, err := s.db.Exec(string(query))
	return err
}

// CreateUser function is responsible for creating the single
// user entity in the database.
func (s *PostgresStorage) CreateUser(user *types.User) error {
	log.Printf("CREATE user")

	query := `INSERT INTO users
	(role, first_name, last_name, email, phone_number, password_encrypted, created_at, updated_at)
	VALUES
	($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err := s.db.Query(
		string(query),
		user.Role,
		user.FirstName,
		user.LastName,
		user.Email,
		user.PhoneNumber,
		user.EncryptedPassword,
		user.CreatedAt,
		user.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

// UpdateUser function is responsible for updating the existing
// user entity with the provided data.
func (s *PostgresStorage) UpdateUser(user *types.User, id int) error {
	log.Printf("UPDATE user with id of %d", user.ID)

	query := `UPDATE users
	SET
	    role = $1
		first_name = $2,
		last_name = $3
		email = $4,
	    phone_number = $5
		password_encrypted = $6
		updated_at = NOW() 
	WHERE id = $7;
	`

	_, err := s.db.Query(
		string(query),
		user.Role,
		user.FirstName,
		user.LastName,
		user.Email,
		user.PhoneNumber,
		user.EncryptedPassword,
		id,
	)
	return err
}

// DeleteUser function is responsible for existing(in the database) product deletion.
func (s *PostgresStorage) DeleteUser(id int) error {
	log.Printf("Request to DELETE user with id of %d", id)

	query := "DELETE FROM users WHERE id = $1"

	_, err := s.db.Query(string(query), id)
	return err
}

// GetUserByID function is responsible for getting individual user from the database using its unique identifier.
func (s *PostgresStorage) GetUserByID(id int) (*types.User, error) {
	query := "SELECT * FROM users WHERE id = $1"

	rows, err := s.db.Query(string(query), id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return readUserFromQuery(rows)
	}

	return nil, fmt.Errorf("user with id of %d does not exist", id)
}

// GetUserByEmail function is responsible for getting individual user from  the database using his email.
func (s *PostgresStorage) GetUserByEmail(email string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE email = $1", email)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return readUserFromQuery(rows)
	}

	return nil, fmt.Errorf("user %d was not found", email)
}

// GetUsers function is responsible for retrieving all users, that are located in the internal database.
func (s *PostgresStorage) GetUsers() ([]*types.User, error) {
	query := "SELECT * FROM users"

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

// readUserFromQuery function is responsible for scanning the database query response to the user datatype.
func readUserFromQuery(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)
	err := rows.Scan(
		&user.ID,
		&user.Role,
		&user.FirstName,
		&user.LastName,
		&user.EncryptedPassword,
		&user.Email,
		&user.PhoneNumber,
		&user.CreatedAt,
		&user.UpdatedAt)

	return user, err
}
