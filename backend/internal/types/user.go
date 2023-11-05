package types

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"regexp"
	"time"
)

type CreateUserRequest struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	PhoneNumber string `json:"phoneNumber"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}

type UpdateUserRequest struct {
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phoneNumber"`
	Password    string    `json:"password"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type User struct {
	ID                int       `json:"id" db:"id"`
	Role              string    `json:"role" db:"role"`
	Email             string    `json:"emailAddress" db:"email"`
	FirstName         string    `json:"firstName" db:"first_name"`
	LastName          string    `json:"lastName" db:"last_name"`
	PhoneNumber       string    `json:"phoneNumber" db:"phone_number"`
	EncryptedPassword string    `json:"-" db:"password_encrypted"`
	CreatedAt         time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt         time.Time `json:"updatedAt" db:"updated_at"`
}

// NewUser is a constructor for the User structure
func NewUser(firstName string, lastName string, phoneNumber string, email string, password string) (*User, error) {
	encpass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &User{
		Email:             email,
		Role:              "ROLE_USER",
		FirstName:         firstName,
		LastName:          lastName,
		PhoneNumber:       phoneNumber,
		EncryptedPassword: string(encpass),
		CreatedAt:         time.Now().UTC(),
		UpdatedAt:         time.Now().UTC(),
	}

	err = validateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// HasValidPassword parses the password and validates if the
// provided password is the password that bound to the user.
func (u *User) HasValidPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(password))
	if err != nil {
		log.Fatal(err)
	}

	return err == nil
}

// ValidateUser is used to validate each field of the provided
// types.User structure.
func validateUser(u *User) error {
	err := validateUserEmail(u.Email)
	if err != nil {
		return err
	}

	err = validateUserPhone(u.PhoneNumber)
	if err != nil {
		return err
	}

	return nil
}

// validateUserPhone is used to validate the provided user phone number
func validateUserPhone(phoneNumber string) error {
	pattern := "^[\\+]?[(]?[0-9]{3}[)]?[-\\s\\.]?[0-9]{3}[-\\s\\.]?[0-9]{4,6}$"

	regex := regexp.MustCompile(pattern)

	if !regex.MatchString(phoneNumber) {
		return fmt.Errorf("user phone number is incorrect")
	}

	return nil
}

// validateUserEmail is used to validate provided user email
func validateUserEmail(email string) error {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	regex := regexp.MustCompile(pattern)

	if !regex.MatchString(email) {
		return fmt.Errorf("user email is incorrect")
	}

	return nil
}
