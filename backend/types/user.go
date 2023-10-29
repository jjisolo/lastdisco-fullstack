package types

import (
	"time"
)

type CreateUserRequest struct {
	FirstName   string    `json:firstName`
	Email       string    `json:lastName`
}

type UpdateUserRequest struct {
	FirstName   string    `json:firstName`
	LastName    string    `json:lastName`
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

func NewUser(firstName string, email string) *User {
	return &User{
		Email    : email,
		FirstName: firstName,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
}