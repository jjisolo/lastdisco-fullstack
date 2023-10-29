package types

type SigninRequest struct {
	Email    string `json:email`
	Password string `json:password`
}

type SigninResponse struct {
	Email string `json:email`
	Token string `json:token`
}