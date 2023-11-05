package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jjisolo/lastdisco-backend/internal/types"
)

func (s *APIServer) handleSignin(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "POST" {
		return fmt.Errorf("method %s is not allowed at this current endpoint", r.Method)
	}

	var req types.SigninRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return err
	}

	user, err := s.storage.GetUserByEmail(req.Email)
	if err != nil {
		return err
	}

	if !user.HasValidPassword(req.Password) {
		return fmt.Errorf("invalid username or password")
	}

	token, err := createJWT(user)
	if err != nil {
		return err
	}

	res := types.SigninResponse{
		Token: token,
		Email: user.Email,
	}

	return WriteJSON(w, http.StatusOK, res)
}
