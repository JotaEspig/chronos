package schedulingapi

import (
	"chronos/config"
	"chronos/pkg/models/scheduling"
	"chronos/pkg/models/user"
	"chronos/pkg/types"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// createScheduling is a scheduling controller that receives a JSON in the body
// of the request and return a status code
func createScheduling(c echo.Context) error {
	claims := c.Get("user").(*jwt.Token).Claims.(*types.JWTClaims)
	s := scheduling.Scheduling{}
	err := json.NewDecoder(c.Request().Body).Decode(&s)
	s.Sanitize(config.StrictPolicy)
	s.UserID = claims.UserID
	if !s.IsValid() || err != nil {
		return c.JSON(http.StatusBadRequest, types.JsonMap{
			"message": "some scheduling field may be missing or invalid",
		})
	}
	tx, err := config.DB.Begin()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, types.JsonMap{
			"message": "creating of database transaction failed. Try again",
		})
	}
	defer tx.Rollback()

	// check if there is already a scheduling for this Time
	var id uint
	err = tx.QueryRow(`
        SELECT id FROM scheduling
        WHERE ("time_id" = ?) AND
        ("start" BETWEEN ? AND ?) AND ("end" BETWEEN ? AND ?);`,
		s.TimeID, s.Start, s.End, s.Start, s.End,
	).Scan(&id)
	if id != 0 || err == nil {
		return c.JSON(http.StatusConflict, types.JsonMap{
			"message": "There's already a scheduling for this time. Select other start or end time",
		})
	}

	err = scheduling.CreateScheduling(tx, &s)
	if err != nil {
		return c.JSON(http.StatusConflict, types.JsonMap{
			"message": "some values aren't valid or are causing database conflict",
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

	t, err := scheduling.FindSchedulingByID(tx, uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, types.JsonMap{
			"message": "scheduling not found",
		})
	}

	tMap := t.ToMap()
	return c.JSON(http.StatusOK, tMap)
}

// getSchedulingsByDate is a scheduling controller that receives a JSON in the body that
// contains the minimal start date and the page you want to retrieve
// JSON should look like this:
// {"date": "2020-01-01 12:00:00", "page": 0}
func getSchedulingsByDate(c echo.Context) error {
	date := c.QueryParam("date")
	pageStr := c.QueryParam("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil || date == "" {
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

	schedulings, err := scheduling.GetSchedulingsByDate(tx, date, uint(page))
	if err != nil {
		return c.JSON(http.StatusNotFound, types.JsonMap{
			"message": "no schedulings found",
		})
	}

	jsonToSend := make([]types.JsonMap, len(schedulings))
	for idx, elem := range schedulings {
		jsonToSend[idx] = elem.ToMap()
	}

	return c.JSON(http.StatusOK, jsonToSend)
}

// updateScheduling is a scheduling controller that receives a param ("id") in the
// url path and a JSON in the body of the request and return a status code.
// Attention: You must send the whole scheduling values even if you don't want to
// update something, e.g. you want to update just the start of the scheduling,
// even so you must include the originals values in the JSON that contains the
// scheduling.
// That's because of the way UpdateScheduling function works
func updateScheduling(c echo.Context) error {
	claims := c.Get("user").(*jwt.Token).Claims.(*types.JWTClaims)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, types.JsonMap{
			"message": "id param is invalid",
		})
	}

	var userID uint
	err = config.DB.QueryRow(`
        SELECT "user"."id" FROM "user"
        INNER JOIN "scheduling" ON "scheduling"."user_id" = "user"."id"
        WHERE "scheduling"."id" = ?;
        `, id).Scan(&userID)
	if err != nil {
		panic(err)
	}

	if claims.Type != user.TypeAdmin && claims.UserID != userID {
		return c.JSON(http.StatusForbidden, types.JsonMap{
			"message": "you cannot access this endpoint as this user",
		})
	}

	e := scheduling.Scheduling{}
	err = json.NewDecoder(c.Request().Body).Decode(&e)
	e.Sanitize(config.StrictPolicy)
	if !e.IsValid() || err != nil {
		return c.JSON(http.StatusBadRequest, types.JsonMap{
			"message": "some scheduling field may be missing or invalid",
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
	err = scheduling.UpdateScheduling(tx, &e)
	if err != nil {
		return c.JSON(http.StatusConflict, types.JsonMap{
			"message": "some values aren't valid or are causing database conflict",
		})
	}

	tx.Commit()
	return c.NoContent(http.StatusOK)
}

// deleteScheduling is a scheduling controller that receives a param ("id") in the
// url path and a JSON in the body of the request and return a status code
func deleteScheduling(c echo.Context) error {
	claims := c.Get("user").(*jwt.Token).Claims.(*types.JWTClaims)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, types.JsonMap{
			"message": "id param is invalid",
		})
	}

	var userID uint
	err = config.DB.QueryRow(`
        SELECT "user"."id" FROM "user"
        INNER JOIN "scheduling" ON "scheduling"."user_id" = "user"."id"
        WHERE "scheduling"."id" = ?;
        `, id).Scan(&userID)
	if err != nil {
		panic(err)
	}

	if claims.Type != user.TypeAdmin && claims.UserID != userID {
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

	err = scheduling.DeleteSchedulingByID(tx, uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, types.JsonMap{
			"message": "unknown error when executing sql query",
		})
	}

	tx.Commit()
	return c.NoContent(http.StatusOK)
}
