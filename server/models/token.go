package models

import "github.com/dgrijalva/jwt-go"

// Token model
type Token struct {
	UserID string
	Email  string
	*jwt.StandardClaims
}
