package common

import "github.com/golang-jwt/jwt/v5"

type CustomJwt struct {
	jwt.RegisteredClaims
	UserId      string
	Name        string
	AccountNo   string
	PhoneNumber string
	Email       string
}
