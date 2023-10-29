package types

import (
	"math/rand"
	"time"
)

type CreateUserRequest struct {
	FirstName   string    `json:firstName`
	LastName    string    `json:lastName`
}

type UpdateUserRequest struct {
	FirstName   string    `json:firstName`
	LastName    string    `json:lastName`
	UpdatedAt   time.Time `json:updatedAt`
}

type User struct {
	ID          int       `json:id`
	Number      int       `json:number`
	FirstName   string    `json:firstName`
	LastName    string    `json:lastName`
	CreatedAt   time.Time `json:createdAt`
	UpdatedAt   time.Time `json:updatedAt`
}

func NewUser(firstName string, lastName string) *User {
	return &User{
		Number   : int(rand.Intn(10000000)),
		FirstName: firstName,
		LastName : lastName,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
}