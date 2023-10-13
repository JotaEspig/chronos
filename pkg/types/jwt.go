package types

import "github.com/golang-jwt/jwt/v5"

type JWTClaims struct {
	Username string `json:"username"`
	Type     uint8  `json:"type"`
	jwt.RegisteredClaims
}
