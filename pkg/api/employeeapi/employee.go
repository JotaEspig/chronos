// package employeeapi provides api endpoints for employee operations
package employeeapi

import (
	"chronos/config"
	"chronos/pkg/models/employee"
	"chronos/pkg/models/user"
	"chronos/pkg/types"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// createEmployee is an employee controller that receives a JSON in the body of
// the request and return a status code
func createEmployee(c echo.Context) error {
	claims := c.Get("user").(*jwt.Token).Claims.(*types.JWTClaims)
	if claims.Type != user.TypeAdmin {
		return c.JSON(http.StatusForbidden, types.JsonMap{
			"message": "you cannot access this endpoint as this user",
		})
	}

	e := employee.Employee{}
	err := json.NewDecoder(c.Request().Body).Decode(&e)
	e.Sanitize(config.StrictPolicy)
	if !e.IsValid() || err != nil {
		return c.JSON(http.StatusBadRequest, types.JsonMap{
			"message": "some employee field may be missing or invalid",
		})
	}
	tx, err := config.DB.Begin()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, types.JsonMap{
			"message": "creating of database transaction failed. Try again",
		})
	}
	defer tx.Rollback()

	err = employee.CreateEmployee(tx, &e)
	if err != nil {
		return c.JSON(http.StatusConflict, types.JsonMap{
			"message": "some values aren't valid or are causing database conflict",
		})
	}

	tx.Commit()
	return c.NoContent(http.StatusCreated)
}

// getEmployee is an employee controller that receives a param ("id") in the url
// path and return a JSON if succeeds or a status code if something went wrong
func getEmployee(c echo.Context) error {
	claims := c.Get("user").(*jwt.Token).Claims.(*types.JWTClaims)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, types.JsonMap{
			"message": "id param is invalid",
		})
	}
	if claims.Type != user.TypeAdmin && claims.UserID != uint(id) {
		return c.JSON(http.StatusForbidden, types.JsonMap{
			"message": "you cannot access this endpoint as this user",
		})
	}

	tx, err := config.DB.Begin()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, types.JsonMap{
			"message": "creating of database transaction failed. Try again",
		})
	}
	defer tx.Rollback()

	e, err := employee.FindEmployeeByID(tx, uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, types.JsonMap{
			"message": "employee not found",
		})
	}

	eMap := e.ToMap()
	return c.JSON(http.StatusOK, eMap)
}

// updateEmployee is an employee controller that receives a param ("id") in the
// url path and a JSON in the body of the request and return a status code.
// Attention: You must send the whole employee values even if you don't want to
// update something, e.g. you want to update just the type of the employee,
// even so you must include the originals values in the JSON that contains the
// employee.
// That's because of the way UpdateEmployee function works
func updateEmployee(c echo.Context) error {
	claims := c.Get("user").(*jwt.Token).Claims.(*types.JWTClaims)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, types.JsonMap{
			"message": "id param is invalid",
		})
	}
	if claims.Type != user.TypeAdmin && claims.UserID != uint(id) {
		return c.JSON(http.StatusForbidden, types.JsonMap{
			"message": "you cannot access this endpoint as this user",
		})
	}

	e := employee.Employee{}
	err = json.NewDecoder(c.Request().Body).Decode(&e)
	e.Sanitize(config.StrictPolicy)
	if !e.IsValid() || err != nil {
		return c.JSON(http.StatusBadRequest, types.JsonMap{
			"message": "some employee field may be missing or invalid",
		})
	}
	tx, err := config.DB.Begin()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, types.JsonMap{
			"message": "creating of database transaction failed. Try again",
		})
	}
	defer tx.Rollback()

	e.ID = uint(id)
	err = employee.UpdateEmployee(tx, &e)
	if err != nil {
		return c.JSON(http.StatusConflict, types.JsonMap{
			"message": "some values aren't valid or are causing database conflict",
		})
	}

	tx.Commit()
	return c.NoContent(http.StatusOK)
}

// deleteEmployee is an employee controller that receives a param ("id") in the
// url path and a JSON in the body of the request and return a status code
func deleteEmployee(c echo.Context) error {
	claims := c.Get("user").(*jwt.Token).Claims.(*types.JWTClaims)
	if claims.Type != user.TypeAdmin {
		return c.JSON(http.StatusForbidden, types.JsonMap{
			"message": "you cannot access this endpoint as this user",
		})
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, types.JsonMap{
			"message": "id param is invalid",
		})
	}
	tx, err := config.DB.Begin()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, types.JsonMap{
			"message": "creating of database transaction failed. Try again",
		})
	}
	defer tx.Rollback()

	err = employee.DeleteEmployeeByID(tx, uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, types.JsonMap{
			"message": "unknown error when executing sql query",
		})
	}

	tx.Commit()
	return c.NoContent(http.StatusOK)
}
