package config

import (
	"chronos/pkg/types"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

var jwtConfig = echojwt.Config{
	NewClaimsFunc: func(c echo.Context) jwt.Claims {
		return new(types.JWTClaims)
	},
	SigningKey: []byte("secret"),
}

var JWTMiddleware = echojwt.WithConfig(jwtConfig)
