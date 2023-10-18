package other

import (
	"chronos/pkg/types"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello There")
}

func needsJWT(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*types.JWTClaims)
	return c.JSON(200, types.JsonMap{
		"user_id": claims.UserID,
		"type":    claims.Type,
	})
}
