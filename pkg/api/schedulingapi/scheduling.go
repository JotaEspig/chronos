package schedulingapi

import (
	"chronos/config"
	"chronos/pkg/models/scheduling"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// createScheduling is a scheduling controller that receives a JSON in the body
// of the request and return a status code
func createScheduling(c echo.Context) error {
	s := scheduling.Scheduling{}
	err := json.NewDecoder(c.Request().Body).Decode(&s)
	s.Sanitize(config.StrictPolicy)
	if !s.IsValid() || err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "some scheduling field may be missing or invalid",
		})
	}
	tx, err := config.DB.Begin()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "creating of database transaction failed. Try again",
		})
	}
	defer tx.Rollback()

	err = scheduling.CreateScheduling(tx, &s)
	if err != nil {
		return c.JSON(http.StatusConflict, map[string]string{
			"error": "some values aren't valid or are causing database conflict",
		})
	}

	tx.Commit()
	return c.NoContent(http.StatusCreated)
}

// getScheduling is a scheduling controller that receives a param ("id") in the url path
// and return a JSON if succeeds or a status code if something went wrong
func getScheduling(c echo.Context) error {
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

	t, err := scheduling.FindSchedulingByID(tx, uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "scheduling not found",
		})
	}

	tMap := t.ToMap()
	return c.JSON(http.StatusOK, tMap)
}

// updateScheduling is a scheduling controller that receives a param ("id") in the
// url path and a JSON in the body of the request and return a status code.
// Attention: You must send the whole scheduling values even if you don't want to
// update something, e.g. you want to update just the start of the scheduling,
// even so you must include the originals values in the JSON that contains the
// scheduling.
// That's because of the way UpdateScheduling function works
func updateScheduling(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "id param is invalid",
		})
	}
	e := scheduling.Scheduling{}
	err = json.NewDecoder(c.Request().Body).Decode(&e)
	e.Sanitize(config.StrictPolicy)
	if !e.IsValid() || err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "some scheduling field may be missing or invalid",
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
	err = scheduling.UpdateScheduling(tx, &e)
	if err != nil {
		return c.JSON(http.StatusConflict, map[string]string{
			"error": "some values aren't valid or are causing database conflict",
		})
	}

	tx.Commit()
	return c.NoContent(http.StatusOK)
}

// deleteScheduling is a scheduling controller that receives a param ("id") in the
// url path and a JSON in the body of the request and return a status code
func deleteScheduling(c echo.Context) error {
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

	err = scheduling.DeleteSchedulingByID(tx, uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "unknown error when executing sql query",
		})
	}

	tx.Commit()
	return c.NoContent(http.StatusOK)
}
