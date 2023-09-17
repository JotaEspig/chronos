package userapi

import (
	"chronos/config"
	"chronos/pkg/models/user"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func getUser(c echo.Context) error {
	idStr := c.Param("id")
	if idStr == "" {
		return c.NoContent(http.StatusBadRequest)
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	tx, err := config.DB.Begin()
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	u, err := user.FindUserByID(tx, uint(id))
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	uMap := u.ToMap()
	return c.JSON(http.StatusOK, uMap)
}
