package storage

import (
	"github.com/jjisolo/lastdisco-backend/internal/types"
)

type Storage interface {
	CreateProduct(*types.Product) error
	DeleteProduct(int) error
	GetProductByID(int) (*types.Product, error)
	GetProducts() ([]*types.Product, error)
	UpdateProduct(*types.Product, int) error

	CreateCategory(*types.ProductCategory) error
	DeleteCategory(int) error
	GetCategoryByID(int) (*types.ProductCategory, error)
	GetCategories() ([]*types.ProductCategory, error)
	UpdateCategory(*types.ProductCategory, int) error

	CreateUser(*types.User) error
	DeleteUser(int) error
	GetUserByID(int) (*types.User, error)
	GetUserByEmail(email string) (*types.User, error)
	GetUsers() ([]*types.User, error)
	UpdateUser(*types.User, int) error
}
