// package timeapi provides api endpoints for time operations
package timeapi

import (
	"chronos/config"
	"chronos/pkg/models/time"
	"chronos/pkg/models/user"
	"chronos/pkg/types"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// createTime is a time controller that receives a JSON in the body of the
// request and return a status code
func createTime(c echo.Context) error {
	claims := c.Get("user").(*jwt.Token).Claims.(*types.JWTClaims)
	if claims.Type != user.TypeEmployee && claims.Type != user.TypeAdmin {
		return c.JSON(http.StatusForbidden, types.JsonMap{
			"message": "you cannot access this endpoint as this user",
		})
	}

	var employeeID uint
	err := config.DB.QueryRow(`
        SELECT "id" FROM "employee"
        WHERE "user_id" = ?;
        `, claims.UserID).Scan(&employeeID)
	if err != nil {
		panic(err)
	}

	t := time.Time{}
	err = json.NewDecoder(c.Request().Body).Decode(&t)
	t.Sanitize(config.StrictPolicy)
	t.EmployeeID = employeeID
	if !t.IsValid() || err != nil {
		return c.JSON(http.StatusBadRequest, types.JsonMap{
			"message": "some time field may be missing or invalid",
		})
	}
	tx, err := config.DB.Begin()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, types.JsonMap{
			"message": "creating of database transaction failed. Try again",
		})
	}
	defer tx.Rollback()

	err = time.CreateTime(tx, &t)
	if err != nil {
		return c.JSON(http.StatusConflict, types.JsonMap{
			"message": "some values aren't valid or are causing database conflict",
		})
	}

	tx.Commit()
	return c.NoContent(http.StatusCreated)
}

// getTime is a time controller that receives a param ("id") in the url path
// and return a JSON if succeeds or a status code if something went wrong
func getTime(c echo.Context) error {
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

	t, err := time.FindTimeByID(tx, uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, types.JsonMap{
			"message": "time not found",
		})
	}

	tMap := t.ToMap()
	return c.JSON(http.StatusOK, tMap)
}

// getTimesByDate is a time controller that receives a JSON in the body that
// contains the minimal start date and the page you want to retrieve
// JSON should look like this:
// {"date": "2020-01-01 12:00:00", "page": 0}
func getTimesByDate(c echo.Context) error {
	date := c.QueryParam("date")
	pageStr := c.QueryParam("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil || date == "" {
		return c.JSON(http.StatusBadRequest, types.JsonMap{
			"message": "some time field may be missing or invalid",
		})
	}
	tx, err := config.DB.Begin()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, types.JsonMap{
			"message": "creating of database transaction failed. Try again",
		})
	}
	defer tx.Rollback()

	times, err := time.GetTimesByDate(tx, date, uint(page))
	if err != nil {
		return c.JSON(http.StatusNotFound, types.JsonMap{
			"message": "no times found",
		})
	}

	jsonToSend := make([]types.JsonMap, len(times))
	for idx, elem := range times {
		jsonToSend[idx] = elem.ToMap()
	}

	return c.JSON(http.StatusOK, jsonToSend)
}

// updateTime is a time controller that receives a param ("id") in the
// url path and a JSON in the body of the request and return a status code.
// Attention: You must send the whole time values even if you don't want to
// update something, e.g. you want to update just the start of the time,
// even so you must include the originals values in the JSON that contains the
// time.
// That's because of the way UpdateTime function works
func updateTime(c echo.Context) error {
	claims := c.Get("user").(*jwt.Token).Claims.(*types.JWTClaims)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, types.JsonMap{
			"message": "id param is invalid",
		})
	}

	var userID uint
	err = config.DB.QueryRow(`
        SELECT "employee"."user_id" FROM "employee"
        INNER JOIN "time" ON "time"."employee_id" = "employee"."id"
        WHERE "time"."id" = ?;
        `, id).Scan(&userID)
	if err != nil {
		panic(err)
	}

	canAccess := (userID == claims.UserID && claims.Type == user.TypeEmployee) ||
		claims.Type == user.TypeAdmin
	if !canAccess {
		return c.JSON(http.StatusForbidden, types.JsonMap{
			"message": "you cannot access this endpoint as this user",
		})
	}

	e := time.Time{}
	err = json.NewDecoder(c.Request().Body).Decode(&e)
	e.Sanitize(config.StrictPolicy)
	if !e.IsValid() || err != nil {
		return c.JSON(http.StatusBadRequest, types.JsonMap{
			"message": "some JSON field may be missing or invalid",
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
	err = time.UpdateTime(tx, &e)
	if err != nil {
		return c.JSON(http.StatusConflict, types.JsonMap{
			"message": "some values aren't valid or are causing database conflict",
		})
	}

	tx.Commit()
	return c.NoContent(http.StatusOK)
}

// deleteTime is a time controller that receives a param ("id") in the
// url path and a JSON in the body of the request and return a status code
func deleteTime(c echo.Context) error {
	claims := c.Get("user").(*jwt.Token).Claims.(*types.JWTClaims)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, types.JsonMap{
			"message": "id param is invalid",
		})
	}

	var userID uint
	err = config.DB.QueryRow(`
        SELECT "employee"."user_id" FROM "employee"
        INNER JOIN "time" ON "time"."employee_id" = "employee"."id"
        WHERE "time"."id" = ?;
        `, id).Scan(&userID)
	if err != nil {
		panic(err)
	}

	canAccess := (userID == claims.UserID && claims.Type == user.TypeEmployee) ||
		claims.Type == user.TypeAdmin
	if !canAccess {
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

	err = time.DeleteTimeByID(tx, uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, types.JsonMap{
			"message": "unknown error when executing sql query",
		})
	}

	tx.Commit()
	return c.NoContent(http.StatusOK)
}
