// package timeapi provides api endpoints for time operations
package timeapi

import (
	"chronos/config"
	"chronos/pkg/models/time"
	"encoding/json"
	"net/http"
	"strconv"
	"chronos/pkg/types"

	"github.com/labstack/echo/v4"
)

// createTime is a time controller that receives a JSON in the body of the
// request and return a status code
func createTime(c echo.Context) error {
	t := time.Time{}
	err := json.NewDecoder(c.Request().Body).Decode(&t)
	t.Sanitize(config.StrictPolicy)
	if !t.IsValid() || err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "some time field may be missing or invalid",
		})
	}
	tx, err := config.DB.Begin()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "creating of database transaction failed. Try again",
		})
	}
	defer tx.Rollback()

	err = time.CreateTime(tx, &t)
	if err != nil {
		return c.JSON(http.StatusConflict, map[string]string{
			"error": "some values aren't valid or are causing database conflict",
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
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "id param is invalid",
		})
	}
	tx, err := config.DB.Begin()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "creating of database transaction failed. Try again",
		})
	}
	defer tx.Rollback()

	t, err := time.FindTimeByID(tx, uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "time not found",
		})
	}

	tMap := t.ToMap()
	return c.JSON(http.StatusOK, tMap)
}

func getTimes(c echo.Context) error {

	tx, err := config.DB.Begin()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "creating of database transaction failed. Try again",
		})
	}
	defer tx.Rollback()

	t, err := time.GetTimes(tx)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "time not found",
		})
	}
  timesJson := make(types.JsonMap)
  timesJson["times"] = t

	//tMap := t.ToMap()
	return c.JSON(http.StatusOK, timesJson)
}
// updateTime is a time controller that receives a param ("id") in the
// url path and a JSON in the body of the request and return a status code.
// Attention: You must send the whole time values even if you don't want to
// update something, e.g. you want to update just the start of the time,
// even so you must include the originals values in the JSON that contains the
// time.
// That's because of the way UpdateTime function works
func updateTime(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "id param is invalid",
		})
	}
	e := time.Time{}
	err = json.NewDecoder(c.Request().Body).Decode(&e)
	e.Sanitize(config.StrictPolicy)
	if !e.IsValid() || err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "some time field may be missing or invalid",
		})
	}
	tx, err := config.DB.Begin()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "creating of database transaction failed. Try again",
		})
	}
	defer tx.Rollback()

	e.ID = uint(id)
	err = time.UpdateTime(tx, &e)
	if err != nil {
		return c.JSON(http.StatusConflict, map[string]string{
			"error": "some values aren't valid or are causing database conflict",
		})
	}

	tx.Commit()
	return c.NoContent(http.StatusOK)
}

// deleteTime is a time controller that receives a param ("id") in the
// url path and a JSON in the body of the request and return a status code
func deleteTime(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "id param is invalid",
		})
	}
	tx, err := config.DB.Begin()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "creating of database transaction failed. Try again",
		})
	}
	defer tx.Rollback()

	err = time.DeleteTimeByID(tx, uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "unknown error when executing sql query",
		})
	}

	tx.Commit()
	return c.NoContent(http.StatusOK)
}
