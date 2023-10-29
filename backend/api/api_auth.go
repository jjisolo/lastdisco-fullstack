package api

import(
	"encoding/json"
	"net/http"
	"fmt"

	"github.com/jjisolo/lastdisco-backend/types"
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
		return fmt.Errorf("not authenticated")
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

func (s *APIServer) handleSignup(w http.ResponseWriter, r *http.Request) error {
	return nil
}