package types

import jwt "github.com/golang-jwt/jwt/v4"

type SigninRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SigninResponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

type JWTClaims struct {
	UserID int `json:"userID"`
	jwt.RegisteredClaims
}
