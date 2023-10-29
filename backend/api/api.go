package api

import (
	"net/http"
	"encoding/json"
	"strconv"
	"log"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jjisolo/lastdisco-backend/types"
	"github.com/jjisolo/lastdisco-backend/deffs"
	"github.com/jjisolo/lastdisco-backend/storage"

	jwt "github.com/golang-jwt/jwt/v4"
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

func raisePermissionDenied(w http.ResponseWriter) {
	WriteJSON(w, http.StatusForbidden, APIError{Error: "permission denied"})
}

// withJWTAuth provides the ability to secure the requests
// with JWT authentication as a middleware.
func withJWTAuth(handlerFunc http.HandlerFunc, s storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("x-jwt-token")

		token, err := validateJWT(tokenString)
		if err != nil {
			raisePermissionDenied(w)
			return
		}
		
		if !token.Valid {
			raisePermissionDenied(w)
			return
		}

		userID, err := extractID(r)
		if err != nil {
			raisePermissionDenied(w)
			return
		}

		user, err := s.GetUserByID(userID)
		if err != nil {
			raisePermissionDenied(w)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		if user.Number != int(claims["accountNumber"].(float64)) {
			raisePermissionDenied(w)
			return
		}

		if err != nil {
			raisePermissionDenied(w)
			return
		}

		handlerFunc(w, r)
	}
}

// createJWT function is responsible for creating the 
// JWT token.
func createJWT(user *types.User) (string, error) {
	claims := &jwt.MapClaims{
		"expiresAt"    : 15000,
		"accountNumber": user.Number,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(deffs.JWT_SECRET))
}

// validateJWT function responsible for the validating 
// provided JWT token.
func validateJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signin method: %v", token.Header["alg"])
		}

		return []byte(deffs.JWT_SECRET), nil
	})
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

	router.HandleFunc("/product",      withJWTAuth(makeHttpHandleFunc(s.handleProduct),     s.storage))
	router.HandleFunc("/product/{id}", withJWTAuth(makeHttpHandleFunc(s.handleProductByID), s.storage))

	router.HandleFunc("/user/create",  makeHttpHandleFunc(s.handlePostUser))
	router.HandleFunc("/user",         withJWTAuth(makeHttpHandleFunc(s.handlePostUser), s.storage))
	router.HandleFunc("/user/{id}",    withJWTAuth(makeHttpHandleFunc(s.handleUserByID), s.storage))

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