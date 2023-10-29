package api

import (
	"net/http"
	"encoding/json"
	"strconv"
	"log"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jjisolo/lastdisco-backend/storage"
)

// APIError represents the error that can
// be raised by the API server.
type APIError struct {
	Error string
}

// APIServer reprenents the abstract server
// datastructure with the storage and address
// inside.
type APIServer struct {
	listenAddress string
	storage       storage.Storage
}

// WriteJSON function is responsible for baking the
// JSON-formatted response from such primitives as 
// result message, HTTP status code etc.
func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

// makeHttpHandleFunc is reponsible for decorating
// the default API endpoint callback functions.
func makeHttpHandleFunc(f func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
		}
	}
}

// NewAPIServer function is responsible for creating 
// the new APIServer object instance.
func NewAPIServer(listenAddress string, storage storage.Storage) *APIServer {
	return &APIServer{
		listenAddress: listenAddress,
		storage      : storage,
	}
}

// APIServer.Run function is reponsible for initializing the
// APIserver instance.
//
// Setup the callbacks and start serving with the binded APIServer's
// port.
func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/product",      makeHttpHandleFunc(s.handleProduct))
	router.HandleFunc("/product/{id}", makeHttpHandleFunc(s.handleProductByID))

	router.HandleFunc("/user",         makeHttpHandleFunc(s.handleUser))
	router.HandleFunc("/user/{id}",    makeHttpHandleFunc(s.handleUserByID))

	log.Println("JSON API server is now running on port: ", s.listenAddress)

	http.ListenAndServe(s.listenAddress, router)
}

// The extractID function is responsible for extracting the ID from the
// request body.
func extractID(r *http.Request) (int, error) {
	idString := mux.Vars(r)["id"]

	id, err := strconv.Atoi(idString)
	if err != nil {
		return id, fmt.Errorf("The provided ID(%s) is invalid", idString)
	}

	return id, nil
}