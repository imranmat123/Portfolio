package handlers

import (
	"ImransProfoiloWebsite/internal/models"
	personalDB_api "ImransProfoiloWebsite/internal/personalDB-api"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"strconv"
)

func CreateEducation(db *sqlx.DB, c echo.Context) error {
	var a models.Education
	err123 := c.Bind(&a)
	fmt.Print(a)
	if err123 != nil {
		return c.JSON(500, map[string]interface{}{
			"status":  500,
			"message": "unable bind",
		})
	}
	err := personalDB_api.CreateEducation(db, &a)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"status":  500,
			"message": "unable to create eduction",
		})
	}
	return c.JSON(200, map[string]interface{}{
		"status": 200,
		"data":   a,
	})
}

func GetAllEducation(db *sqlx.DB, c echo.Context) error {
	a, err := personalDB_api.GetAllEducation(db)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"status":  500,
			"message": "unable to get all education",
		})
	}

	if len(a) == 0 {
		return c.JSON(404, map[string]interface{}{
			"status":  404,
			"message": "unable to find work %v",
		})
	}

	return c.JSON(200, map[string]interface{}{
		"status": 200,
		"data":   a,
	})
}

func GetEducationByEducationID(db *sqlx.DB, c echo.Context) error {
	a, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"status":  500,
			"message": "unable to convert id to int",
		})
	}
	b, err1 := personalDB_api.GetEducationByEducationID(db, int64(a))
	if err1 != nil {
		return c.JSON(404, map[string]interface{}{
			"status":  404,
			"message": "unable to find eduction by id",
		})
	}
	return c.JSON(200, map[string]interface{}{
		"status": 200,
		"data":   b,
	})
}
func GetEducationByUserID(db *sqlx.DB, c echo.Context) error {
	a, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"status":  500,
			"message": "unable to convert id to int",
		})
	}
	b, err1 := personalDB_api.GetUsersEducationByUserID(db, int64(a))
	if err1 != nil {
		return c.JSON(404, map[string]interface{}{
			"status":  404,
			"message": "unable to find eduction by id",
		})
	}
	return c.JSON(200, map[string]interface{}{
		"status": 200,
		"data":   b,
	})
}
func UpdateEducation(db *sqlx.DB, c echo.Context) error {
	a, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"status":  500,
			"message": "unable to convert id to int",
		})
	}
	b, err1 := personalDB_api.GetEducationByEducationID(db, int64(a))
	if err1 != nil {
		return c.JSON(404, map[string]interface{}{
			"status":  404,
			"message": "unable to find eduction by id",
		})
	}
	d, err2 := personalDB_api.UpdateEducation(db, b)
	if err2 != nil {
		return c.JSON(404, map[string]interface{}{
			"status":  404,
			"message": "unable to update eduction",
		})
	}
	return c.JSON(200, map[string]interface{}{
		"status": 200,
		"data":   d,
	})
}
func DeleteEducation(db *sqlx.DB, c echo.Context) error {
	a, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"status":  500,
			"message": "unable to convert id to int",
		})
	}
	b, err1 := personalDB_api.GetEducationByEducationID(db, int64(a))
	if err1 != nil {
		return c.JSON(404, map[string]interface{}{
			"status":  404,
			"message": "unable to find eduction by id",
		})
	}
	err2 := personalDB_api.DeleteEducation(db, b)
	if err2 != nil {
		return c.JSON(400, map[string]interface{}{
			"status":  400,
			"message": "could not delete education"})

	}
	return c.JSON(200, map[string]interface{}{
		"status":  200,
		"message": "successfully deleted edacation",
	})
}
