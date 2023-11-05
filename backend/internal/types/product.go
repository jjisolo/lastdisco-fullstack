package types

import "time"

type CreateProductRequest struct {
	ShortName   string `json:"shortName"`
	Description string `json:"description"`
	CategoryID  int    `json:"categoryID"`
	Count       int    `json:"count"`
	Price       string `json:"price"`
}

type UpdateProductRequest struct {
	ShortName   string    `json:"shortName"`
	Description string    `json:"description"`
	Count       int       `json:"count"`
	Price       string    `json:"price"`
	CategoryID  int       `json:"categoryID"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type Product struct {
	ID          int       `json:"id"`
	ShortName   string    `json:"shortName"`
	Description string    `json:"description"`
	Count       int       `json:"count"`
	Price       string    `json:"price"`
	CategoryID  int       `json:"categoryID"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// NewProduct function is creating the new product
func NewProduct(shortName string, description string, count int, price string, categoryId int) *Product {
	return &Product{
		ShortName:   shortName,
		Description: description,
		Count:       count,
		Price:       price,
		CategoryID:  categoryId,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}
}
