package storage

import (
	"github.com/jjisolo/lastdisco-backend/types"
)

type Storage interface {
	CreateProduct (*types.Product)      error
	DeleteProduct (int)                 error
	GetProductByID(int)                 (*types.Product, error)
	GetProducts   ()                    ([]*types.Product, error)
	UpdateProduct (*types.Product, int) error

	CreateUser    (*types.User)      error
	DeleteUser    (int)              error
	GetUserByID   (int)              (*types.User, error)
	GetUserByEmail(email string)     (*types.User, error) 
	GetUsers      ()                 ([]*types.User, error)
	UpdateUser    (*types.User, int) error
}

