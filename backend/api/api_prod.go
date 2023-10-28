package api

import (
	"encoding/json"
	"net/http"
	"time"
	"fmt"

	"github.com/jjisolo/lastdisco-backend/types"
)

// handleProduct serves the '/product' api enpoint, with allowed
// GET and POST methods.
func (s *APIServer) handleProduct(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		return s.handleGetProduct(w, r)
	case "POST":
		return s.handleCreateProduct(w, r)
	}

	return fmt.Errorf("The requested method is not allowed at current endpoint")
}

// handleProduct serves the '/product/{id}' API endpoint, with allowed
// GET and DELETE methods.
func (s *APIServer) handleProductByID(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		return s.handleGetProductByID(w, r)
	case "DELETE":
		return s.handleDeleteProduct(w, r)
	case "PUT":
		return s.handleUpdateProduct(w, r)
	}

	return fmt.Errorf("The requested method is not allowed at current endpoint")
}

// handleGetProduct serves the GET request for the '/product' API endpoint.
// The result of this function is all products that are registered in the
// internal database.
func (s *APIServer) handleGetProduct(w http.ResponseWriter, r *http.Request) error {
	products, err := s.storage.GetProducts()
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, products)
}

// handleGetProduct serves the GET request for the '/product/{id}' API endpoint.
// The result of this function is the product with requested id, that is
// registered in the internal database.
func (s *APIServer) handleGetProductByID(w http.ResponseWriter, r *http.Request) error {
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
// The result of this function is the brand new product.
func (s *APIServer) handleCreateProduct(w http.ResponseWriter, r *http.Request) error {
	productRequest := new(types.CreateProductRequest)
	if err := json.NewDecoder(r.Body).Decode(productRequest); err != nil {
		return err
	}

	product := &types.Product{
		ShortName   : productRequest.ShortName,
		Description : productRequest.Description,
		Count       : productRequest.Count,
		Price       : productRequest.Price,
		UpdatedAt   : time.Now().UTC(),
	}

	if err := s.storage.CreateProduct(product); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, product) 
}

// handleCreateProduct serves the DELETE request for the '/product/{id}' API endpoint.
// The result of this function is the freshly baked product stored in the internal 
// database.
func (s *APIServer) handleDeleteProduct(w http.ResponseWriter, r *http.Request) error {
	id, err := extractID(r)
	if err != nil {
		return err
	}

	if err := s.storage.DeleteProduct(id); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, map[string]int{"deleted": id})
}

// handleCreateProduct serves the PUT request for the '/product/{id}' API endpoint.
// The result of this function is the freshly baked product stored in the internal 
// database.
var productUpdateRequiredParams = []string{
	"shortName", "description", "count", "price",
}

func (s *APIServer) handleUpdateProduct(w http.ResponseWriter, r *http.Request) error {
	id, err := extractID(r)
	if err != nil {
		return err
	}

	productRequest := new(types.UpdateProductRequest)
	if err := json.NewDecoder(r.Body).Decode(productRequest); err != nil {
		return err
	}

	product := types.NewProduct(productRequest.ShortName, productRequest.Description, productRequest.Count, productRequest.Price)
	if err := s.storage.UpdateProduct(product, id); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, product) 
}
