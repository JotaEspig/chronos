package employeeapi

import (
	"chronos/config"
	"chronos/pkg/models/employee"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

// createEmployee is an employee controller that receives a JSON in the body of
// the request and return a status code
func createEmployee(c echo.Context) error {
	e := employee.Employee{}
	err := json.NewDecoder(c.Request().Body).Decode(&e)
	if !e.IsValid() || err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "some employee field may be missing or user_id equals 0",
		})
	}
	tx, err := config.DB.Begin()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "creating of database transaction failed. Try again",
		})
	}
	defer tx.Rollback()

	err = employee.CreateEmployee(tx, &e)
	if err != nil {
		return c.JSON(http.StatusConflict, map[string]string{
			"error": "some values causes are causing conflict",
		})
	}

	tx.Commit()
	return c.NoContent(http.StatusCreated)
}
