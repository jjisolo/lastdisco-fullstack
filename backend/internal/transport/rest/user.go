package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jjisolo/lastdisco-backend/internal/types"
)

// handleUser serves the '/user' rest enpoint, with allowed
// GET and POST methods.
func (s *APIServer) handleGetUser(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetUsers(w, r)
	}

	return fmt.Errorf("the requested method is not allowed at current endpoint")
}

// handleUser serves the '/user' rest enpoint, with allowed
// GET and POST methods.
func (s *APIServer) handlePostUser(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "POST" {
		return s.handleCreateUser(w, r)
	}

	return fmt.Errorf("the requested method is not allowed at current endpoint")
}

// handleUser serves the '/user/{id}' API endpoint, with allowed
// GET and DELETE methods.
func (s *APIServer) handleUserByID(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		return s.handleGetUserByID(w, r)
	case "DELETE":
		return s.handleDeleteUser(w, r)
	case "PUT":
		return s.handleUpdateUser(w, r)
	}

	return fmt.Errorf("the requested method is not allowed at current endpoint")
}

// handleGetUser serves the GET request for the '/user' API endpoint.
// The result of this function is all users that are registered in the
// internal database.
func (s *APIServer) handleGetUsers(w http.ResponseWriter, r *http.Request) error {
	users, err := s.storage.GetUsers()
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, users)
}

// handleGetProduct serves the GET request for the '/users/{id}' API endpoint.
// The result of this function is the user with requested id, that is
// registered in the internal database.
func (s *APIServer) handleGetUserByID(w http.ResponseWriter, r *http.Request) error {
	id, err := extractID(r)
	if err != nil {
		return err
	}

	jwtUserID := r.Context().Value(ContextUserIDKey).(int)
	if id != jwtUserID {
		return fmt.Errorf("permission denied")
	}

	user, err := s.storage.GetUserByID(id)
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, user)
}

// handleCreateUser serves the POST request for the '/user' API endpoint.
// The result of this function is the brand-new user.
func (s *APIServer) handleCreateUser(w http.ResponseWriter, r *http.Request) error {
	userRequest := new(types.CreateUserRequest)
	if err := json.NewDecoder(r.Body).Decode(userRequest); err != nil {
		return err
	}

	user, err := types.NewUser(
		userRequest.FirstName,
		userRequest.LastName,
		userRequest.Email,
		userRequest.PhoneNumber,
		userRequest.Password,
	)
	if err != nil {
		return err
	}

	if err := s.storage.CreateUser(user); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, user)
}

// handleCreateUser serves the DELETE request for the '/user/{id}' API endpoint.
// The result of this function is the freshly baked user stored in the internal
// database.
func (s *APIServer) handleDeleteUser(w http.ResponseWriter, r *http.Request) error {
	id, err := extractID(r)
	if err != nil {
		return err
	}

	if err := s.storage.DeleteUser(id); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, map[string]int{"deleted": id})
}

// handleCreateUser serves the PUT request for the '/user/{id}' API endpoint.
// The result of this function is the freshly baked user stored in the internal
// database.
func (s *APIServer) handleUpdateUser(w http.ResponseWriter, r *http.Request) error {
	id, err := extractID(r)
	if err != nil {
		return err
	}

	userRequest := new(types.UpdateUserRequest)
	if err := json.NewDecoder(r.Body).Decode(userRequest); err != nil {
		return err
	}

	user, err := types.NewUser(
		userRequest.FirstName,
		userRequest.LastName,
		userRequest.Email,
		userRequest.PhoneNumber,
		userRequest.Password,
	)
	if err != nil {
		return err
	}

	if err = s.storage.UpdateUser(user, id); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, user)
}

const ContextUserIDKey = "userid"
