package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/jjisolo/lastdisco-backend/internal/types"
)

// handleCategory serves the '/category' rest endpoint, with allowed GET and POST methods.
func (s *APIServer) handleCategory(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		return s.handleGetCategory(w, r)
	case "POST":
		return s.handleCreateCategory(w, r)
	}

	return fmt.Errorf("the requested method is not allowed at current endpoint")
}

// handleCategoryByID serves the '/category/{id}' API endpoint, with allowed GET and DELETE methods.
func (s *APIServer) handleCategoryByID(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		return s.handleGetCategoryByID(w, r)
	case "DELETE":
		return s.handleDeleteCategory(w, r)
	case "PUT":
		return s.handleUpdateCategory(w, r)
	}

	return fmt.Errorf("the requested method is not allowed at current endpoint")
}

// handleGetCategory serves the GET request for the '/category' API endpoint. The result of this function is all
// products that are registered in the internal database.
func (s *APIServer) handleGetCategory(w http.ResponseWriter, r *http.Request) error {
	products, err := s.storage.GetCategories()
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, products)
}

// handleGetCategory serves the GET request for the '/category/{id}' API endpoint. The result of this function is the
// product with requested id, that is registered in the internal database.
func (s *APIServer) handleGetCategoryByID(w http.ResponseWriter, r *http.Request) error {
	id, err := extractID(r)
	if err != nil {
		return err
	}

	product, err := s.storage.GetCategoryByID(id)
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, product)
}

// handleCreateCategory serves the POST request for the '/category' API endpoint. The result of this function is the
// brand-new category.
func (s *APIServer) handleCreateCategory(w http.ResponseWriter, r *http.Request) error {
	categoryRequest := new(types.CreateProductCategoryRequest)
	if err := json.NewDecoder(r.Body).Decode(categoryRequest); err != nil {
		return err
	}

	category := &types.ProductCategory{
		Name:        categoryRequest.Name,
		Description: categoryRequest.Description,
		UpdatedAt:   time.Now().UTC(),
	}

	if err := s.storage.CreateCategory(category); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, category)
}

// handleCreateCategory serves the DELETE request for the '/category/{id}' API endpoint. The result of this function is
// the freshly baked product stored in the internal database.
func (s *APIServer) handleDeleteCategory(w http.ResponseWriter, r *http.Request) error {
	id, err := extractID(r)
	if err != nil {
		return err
	}

	if err := s.storage.DeleteCategory(id); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, map[string]int{"deleted": id})
}

// handleUpdateCategory serves the PUT request for the '/category/{id}' API endpoint. The result of this function is the
// freshly baked product stored in the internal database.
func (s *APIServer) handleUpdateCategory(w http.ResponseWriter, r *http.Request) error {
	id, err := extractID(r)
	if err != nil {
		return err
	}

	categoryRequest := new(types.UpdateProductCategoryRequest)
	if err := json.NewDecoder(r.Body).Decode(categoryRequest); err != nil {
		return err
	}

	category, err := types.NewProductCategory(categoryRequest.Name, categoryRequest.Description)
	if err != nil {
		return err

	}

	if err := s.storage.UpdateCategory(category, id); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, category)
}
