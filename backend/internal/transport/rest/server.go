package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jjisolo/lastdisco-backend/config"
	"github.com/jjisolo/lastdisco-backend/internal/setup"
	"github.com/jjisolo/lastdisco-backend/internal/storage"
	"github.com/jjisolo/lastdisco-backend/internal/types"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"

	"github.com/golang-jwt/jwt/v4"
)

// APIError represents the error that can be raised by the API server.
type APIError struct {
	Error string
}

// APIServer represents the abstract server data structure with the storage and address inside.
type APIServer struct {
	listenAddress string
	storage       storage.Storage
}

// WriteJSON function is responsible for baking the JSON-formatted response from such primitives as a result message,
// HTTP status code, etc.
func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

// raisePermissionDenied raises the `permission denied` error
func raisePermissionDenied(w http.ResponseWriter) {
	err := WriteJSON(w, http.StatusForbidden, APIError{Error: "permission denied"})
	if err != nil {
		return
	}
}

// withJWTAuth provides the ability to secure the requests with JWT authentication as middleware.
func withJWTAuth(handlerFunc http.HandlerFunc, s storage.Storage, allowedGroups []string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rToken := r.Header.Get("Authorization")
		sToken := strings.Split(rToken, "Bearer ")

		var fToken string
		if len(sToken) == 0 {
			raisePermissionDenied(w)
		} else if len(sToken) == 1 {
			fToken = sToken[0]
		} else {
			fToken = sToken[1]
		}

		token, err := validateJWT(fToken)
		if err != nil {
			if config.TESTING {
				err := WriteJSON(w, http.StatusForbidden, APIError{Error: err.Error()})
				if err != nil {
					return
				}
			} else {
				raisePermissionDenied(w)
			}

			return
		}

		var userID int
		if claims, ok := token.Claims.(*types.JWTClaims); ok && token.Valid {
			userID = claims.UserID
		} else {
			if config.TESTING {
				err := WriteJSON(w, http.StatusForbidden, APIError{Error: "claims validation failed"})
				if err != nil {
					return
				}
			} else {
				raisePermissionDenied(w)
			}

			return
		}

		user, err := s.GetUserByID(userID)
		if err != nil {
			if config.TESTING {
				err := WriteJSON(w, http.StatusForbidden, APIError{Error: err.Error()})
				if err != nil {
					return
				}
			} else {
				raisePermissionDenied(w)
			}

			return
		}

		found := false
		for _, str := range allowedGroups {
			if str == user.Role {
				found = true
			}
		}

		if !found {
			if config.TESTING {
				err := WriteJSON(w, http.StatusForbidden, APIError{Error: "specified role does tolerate with the policy"})
				if err != nil {
					return
				}
			} else {
				raisePermissionDenied(w)
			}
		}

		ctx := context.WithValue(r.Context(), ContextUserIDKey, userID)
		handlerFunc(w, r.WithContext(ctx))
	}
}

// createJWT function is responsible for creating the JWT token.
func createJWT(user *types.User) (string, error) {
	claims := &types.JWTClaims{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "lastdisco",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.JWT_SECRET))
}

// validateJWT function responsible for the validating provided JWT token.
func validateJWT(tokenString string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenString, &types.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signin method: %v", token.Header["alg"])
		}

		return []byte(config.JWT_SECRET), nil
	})
}

// makeHttpHandleFunc is responsible for decorating the default API endpoint callback functions.
func makeHttpHandleFunc(f func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			err := WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
			if err != nil {
				return
			}
		}
	}
}

// NewAPIServer function is responsible for creating the new APIServer object instance.
func NewAPIServer(listenAddress string, storage storage.Storage) *APIServer {
	return &APIServer{
		listenAddress: listenAddress,
		storage:       storage,
	}
}

// Run function is responsible for initializing the APIServer instance. Set up the callbacks and start serving with
// the bound APIServer's port.
func (s *APIServer) Run() {
	setup.SeedDatabase(s.storage)

	router := mux.NewRouter()

	router.HandleFunc("/v1/product/get", makeHttpHandleFunc(s.handleGetProduct))
	router.HandleFunc("/v1/product/create", withJWTAuth(makeHttpHandleFunc(s.handleCreateProduct), s.storage, config.PROD_CRUD_ROLES))

	router.HandleFunc("/v1/product/get/{id}", makeHttpHandleFunc(s.handleGetProductByID))
	router.HandleFunc("/v1/product/delete/{id}", withJWTAuth(makeHttpHandleFunc(s.handleDeleteProductByID), s.storage, config.PROD_CRUD_ROLES))
	router.HandleFunc("/v1/product/update/{id}", withJWTAuth(makeHttpHandleFunc(s.handleUpdateProductByID), s.storage, config.PROD_CRUD_ROLES))

	router.HandleFunc("/v1/signin", makeHttpHandleFunc(s.handleSignin))
	router.HandleFunc("/v1/signup", makeHttpHandleFunc(s.handlePostUser))

	router.HandleFunc("/v1/user", withJWTAuth(makeHttpHandleFunc(s.handleGetUser), s.storage, config.USER_CRUD_ROLES))
	router.HandleFunc("/v1/user/{id}", withJWTAuth(makeHttpHandleFunc(s.handleUserByID), s.storage, config.USER_CRUD_ROLES))

	log.Println("JSON API server is now running on port: ", s.listenAddress)

	err := http.ListenAndServe(s.listenAddress, router)
	if err != nil {
		log.Fatal("http.ListenAndServe returned error code!")
		return
	}
}

// The extractID function is responsible for extracting the ID from the request body.
func extractID(r *http.Request) (int, error) {
	idString := mux.Vars(r)["id"]

	id, err := strconv.Atoi(idString)
	if err != nil {
		return id, fmt.Errorf("the provided ID(%s) is invalid", idString)
	}

	return id, nil
}
