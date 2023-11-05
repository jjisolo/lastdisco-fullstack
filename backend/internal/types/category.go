package types

import (
	"fmt"
	"regexp"
	"time"
)

type CreateProductCategoryRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateProductCategoryRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ProductCategory struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func NewProductCategory(name string, description string) (*ProductCategory, error) {
	category := &ProductCategory{
		Name:        name,
		Description: description,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}

	err := validateCategoryName(name)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func validateCategoryName(name string) error {
	pattern := "^[a-z]*$"

	regex := regexp.MustCompile(pattern)

	if !regex.MatchString(name) {
		return fmt.Errorf("category name is incorrect")
	}

	return nil
}
