package types

import (
	"time"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserRequest struct {
	FirstName   string    `json:firstName`
	Email       string    `json:lastName`
	Password    string    `json:password`
}

type UpdateUserRequest struct {
	FirstName   string    `json:firstName`
	Email       string    `json:lastName`
	Password    string    `json:password`
	UpdatedAt   time.Time `json:updatedAt`
}

type User struct {
	ID               int       `json:id`
	Email            string    `json:emailAddress`
	FirstName        string    `json:firstName`
	EncryptedPasword string    `json:-`
	CreatedAt        time.Time `json:createdAt`
	UpdatedAt        time.Time `json:updatedAt`
}

func (u *User) HasValidPassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.EncryptedPasword), []byte(password)) == nil
}

func NewUser(firstName string, email string, password string) (*User, error) {
	encpass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &User{
		Email           : email,
		FirstName       : firstName,
		EncryptedPasword: string(encpass),
		CreatedAt       : time.Now().UTC(),
		UpdatedAt       : time.Now().UTC(),
	}, nil
}