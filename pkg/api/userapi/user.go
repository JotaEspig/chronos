package userapi

import (
	"chronos/config"
	"chronos/pkg/models/user"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

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
	return c.NoContent(http.StatusCreated)
}

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
