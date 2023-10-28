package types

import "time"

type User struct {
	ID          int       `json:id`
	FirstName   string    `json:firstName`
	LastName    string    `json:lastName`
	CreatedAt   time.Time `json:createdAt`
	UpdatedAt   time.Time `json:updatedAt`
}

func NewUser(firstName string, lastName string) *User {
	return &User{
		FirstName: firstName,
		LastName : lastName,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
}