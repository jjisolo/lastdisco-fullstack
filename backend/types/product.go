package types

import "time"

type CreateProductRequest struct {
	ShortName   string  `json:shortName`
	Description string  `json:description`
	Count       int     `json:count`
	Price       string  `json:price`
}

type UpdateProductRequest struct {
	ShortName   string    `json:shortName`
	Description string    `json:description`
	Count       int       `json:count`
	Price       string    `json:price`
	UpdatedAt   time.Time `json:updatedAt`
}

type Product struct {
	ID 		 	int       `json:id`
	ShortName 	string    `json:shortName`
	Description string    `json:description`
	Count       int       `json:count`
	Price       string    `json:price`
	CreatedAt   time.Time `json:createdAt`
	UpdatedAt   time.Time `json:updatedAt`
}

func NewProduct(shortName string, description string, count int, price string) *Product {
	return &Product{
		ShortName:   shortName,
		Description: description,
		Count: 		 count,
		Price:		 price,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}
}