package api

import(
	"encoding/json"
	"net/http"

	"github.com/jjisolo/lastdisco-backend/types"
)

func (s *APIServer) handleSignin(w http.ResponseWriter, r *http.Request) error {
	var req types.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, req)
}

func (s *APIServer) handleSignup(w http.ResponseWriter, r *http.Request) error {
	return nil
}