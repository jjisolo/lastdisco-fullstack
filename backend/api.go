package main

import (
	"net/http"
	"encoding/json"
	"log"
	"github.com/gorilla/mux"
)

type APIError struct {
	Error string
}

type APIServer struct {
	listenAddress string
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Conent-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

func makeHttpHandleFunc(f func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
		}
	}
}

func NewAPIServer(listenAddress string) *APIServer {
	return &APIServer{
		listenAddress: listenAddress,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/product", makeHttpHandleFunc(s.handleProduct))

	log.Println("JSON API server is now running on port: ", s.listenAddress)
	http.ListenAndServe(s.listenAddress, router)
}

func (s *APIServer) handleProduct(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		return s.handleGetProduct(w, r)
	case "POST":
		return s.handleCreateProduct(w, r)
	case "PUT":
		return s.handleUpdateProduct(w, r)
	case "DELETE":
		return s.handleDeleteProduct(w, r)
	}

	return nil
}

func (s *APIServer) handleGetProduct(w http.ResponseWriter, r *http.Request) error {
	product := NewProduct("Short Name", "Description", 10, 50.0)
	return WriteJSON(w, http.StatusOK, product)
}

func (s *APIServer) handleCreateProduct(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleDeleteProduct(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleUpdateProduct(w http.ResponseWriter, r *http.Request) error {
	return nil
}