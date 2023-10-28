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
}

