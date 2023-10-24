package main

type Storage interface {
	CreateProduct (*Product) error
	DeleteProduct (int)      error
	UpdateProduct (*Product) error
	GetProductByID(int)      (*Product, error)
}