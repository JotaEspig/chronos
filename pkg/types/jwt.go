package types

import "github.com/golang-jwt/jwt/v5"

type JWTClaims struct {
	UserID uint  `json:"user_id"`
	Type   uint8 `json:"type"`
	jwt.RegisteredClaims
}
