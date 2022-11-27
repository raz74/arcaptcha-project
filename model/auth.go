package model

import (
	"github.com/dgrijalva/jwt-go"
)

type JwtCustomClimes struct {
	ID    int  `json:"id"`
	jwt.StandardClaims
}
