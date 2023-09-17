package userapi

import (
	"chronos/config"
	"chronos/pkg/models/user"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// createUser is a user controller that receives a JSON in the body of the
// request and return a status code
func createUser(c echo.Context) error {
	u := user.User{}
	err := json.NewDecoder(c.Request().Body).Decode(&u)
	if !u.IsValid() || err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	tx, err := config.DB.Begin()
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	defer tx.Rollback()

	err = user.CreateUser(tx, &u)
	if err != nil {
		return c.NoContent(http.StatusConflict)
	}

	tx.Commit()
	return c.NoContent(http.StatusCreated)
}

// getUser is a user controller that receives a param ("id") in the url path
// and return a JSON if succeeds or a status code if something went wrong
func getUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	tx, err := config.DB.Begin()
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	defer tx.Rollback()

	u, err := user.FindUserByID(tx, uint(id))
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	uMap := u.ToMap()
	return c.JSON(http.StatusOK, uMap)
}

// updateUser is a user controller that receives a param ("id") in the url path
// and a JSON in the body of the request and return a status code.
// Attention: You must send the whole user values even if you don't want to
// update something, e.g. you want to update just the age of the user, even so
// you must include the original value in the JSON that contains the user.
// That's because of the way UpdateUser function works
func updateUser(c echo.Context) error {
	u := user.User{}
	err := json.NewDecoder(c.Request().Body).Decode(&u)
	if !u.IsValid() || err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	tx, err := config.DB.Begin()
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	defer tx.Rollback()

	u.ID = uint(id)
	err = user.UpdateUser(tx, &u)
	if err != nil {
		return c.NoContent(http.StatusConflict)
	}

	tx.Commit()
	return c.NoContent(http.StatusOK)
}

// deleteUser is a user controller that receives a param ("id") in the url path
// and a JSON in the body of the request and return a status code
func deleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	tx, err := config.DB.Begin()
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	defer tx.Rollback()

	err = user.DeleteUserByID(tx, uint(id))
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	tx.Commit()
	return c.NoContent(http.StatusOK)
}
