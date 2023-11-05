package rest

import (
	"encoding/json"
	"fmt"
	"github.com/jjisolo/lastdisco-backend/internal/types"
	"net/http"
)

// handleGetProduct serves the GET request for the '/product' API endpoint. The result of this
// function is all products that are registered in the internal database.
func (s *APIServer) handleGetProduct(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "GET" {
		return fmt.Errorf("the requested method is not allowed at current endpoint")
	}

	products, err := s.storage.GetProducts()
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, products)
}

// handleGetProduct serves the GET request for the '/product/{id}' API endpoint. The result of
// this function is the product with requested id, that is registered in the internal database.
func (s *APIServer) handleGetProductByID(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "GET" {
		return fmt.Errorf("the requested method is not allowed at current endpoint")
	}

	id, err := extractID(r)
	if err != nil {
		return err
	}

	product, err := s.storage.GetProductByID(id)
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, product)
}

// handleCreateProduct serves the POST request for the '/product' API endpoint.
// The result of this function is the brand-new product.
func (s *APIServer) handleCreateProduct(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "POST" {
		return fmt.Errorf("the requested method is not allowed at current endpoint")
	}

	productRequest := new(types.CreateProductRequest)
	if err := json.NewDecoder(r.Body).Decode(productRequest); err != nil {
		return err
	}

	product := types.NewProduct(
		productRequest.ShortName,
		productRequest.Description,
		productRequest.Count,
		productRequest.Price,
		productRequest.CategoryID,
	)

	if err := s.storage.CreateProduct(product); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, product)
}

// handleDeleteProductByID serves the DELETE request for the '/product/{id}' API endpoint.
// The result of this function is the freshly baked product stored in the internal database.
func (s *APIServer) handleDeleteProductByID(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "DELETE" {
		return fmt.Errorf("the requested method is not allowed at current endpoint")
	}

	id, err := extractID(r)
	if err != nil {
		return err
	}

	if err := s.storage.DeleteProduct(id); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, map[string]int{"deleted": id})
}

// handleCreateProduct serves the PUT/PATCH request for the '/product/{id}' API endpoint.
// The result of this function is the freshly baked product stored in the internal database.
func (s *APIServer) handleUpdateProductByID(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "PUT" && r.Method != "PATCH" {
		return fmt.Errorf("the requested method is not allowed at current endpoint")
	}

	id, err := extractID(r)
	if err != nil {
		return err
	}

	productRequest := new(types.UpdateProductRequest)
	if err := json.NewDecoder(r.Body).Decode(productRequest); err != nil {
		return err
	}

	product := types.NewProduct(
		productRequest.ShortName,
		productRequest.Description,
		productRequest.Count,
		productRequest.Price,
		productRequest.CategoryID,
	)
	if err := s.storage.UpdateProduct(product, id); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, product)
}
