package main

import "math/rand"

type Product struct {
	ID 		 	int     `json:id`
	ShortName 	string  `json:shortName`
	Description string  `json:description`
	Count       int64   `json:count`
	Price       float64 `json:price`
}

func NewProduct(shortName string, description string, count int64, price float64) *Product {
	return &Product{
		ID: 		 rand.Intn(10000),
		ShortName:   shortName,
		Description: description,
		Count: 		 count,
		Price:		 price,
	}
}

