// package login provides endpoints for login and signup
package login

import (
	"chronos/config"
	"chronos/pkg/models/user"
	"chronos/pkg/types"
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func login(c echo.Context) error {

	u := user.User{}
	err := json.NewDecoder(c.Request().Body).Decode(&u)
	if !u.IsValid() || err != nil {
		return c.JSON(http.StatusBadRequest, types.JsonMap{
			"error": "some user field may be missing or invalid",
		})
	}
	tx, err := config.DB.Begin()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, types.JsonMap{
			"error": "creating of database transaction failed. Try again",
		})
	}
	defer tx.Rollback()
	savedUser, err := user.FindUserByUsername(tx, u.Username)
	if err != nil {
		return c.JSON(http.StatusNotFound, types.JsonMap{
			"error": "user not found",
		})
	}
	if !savedUser.Validate(u.Username, u.Password) {
		return c.JSON(http.StatusUnauthorized, types.JsonMap{
			"error": "unauthorized",
		})
	}
	claims := &types.JWTClaims{
		u.Username,
		"",
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, types.JsonMap{"token": t})
}