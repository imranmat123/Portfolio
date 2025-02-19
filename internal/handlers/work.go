package handlers

import (
	"ImransProfoiloWebsite/internal/models"
	personalDB_api "ImransProfoiloWebsite/internal/personalDB-api"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"strconv"
)

func CreateWork(db *sqlx.DB, c echo.Context) error {
	a := models.Work{}
	err1 := personalDB_api.CreateWork(db, &a)
	if err1 != nil {
		return c.JSON(400, map[string]interface{}{
			"status":  400,
			"message": "unable to create work",
		})
	}
	return c.JSON(200, map[string]interface{}{
		"status": 200,
		"data":   a,
	})
}

func GetAllWork(db *sqlx.DB, c echo.Context) error {
	a, err := personalDB_api.GetAllWork(db)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"status":  500,
			"message": "unable to get all work",
		})
	}

	if len(*a) == 0 {
		return c.JSON(404, map[string]interface{}{
			"status":  404,
			"message": "There is no work",
		})
	}

	return c.JSON(200, map[string]interface{}{
		"status": 200,
		"data":   a,
	})

}

func GetByWorkID(db *sqlx.DB, c echo.Context) error {
	a, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"status":  500,
			"message": "unable to get work id",
		})
	}
	b, err1 := personalDB_api.GetAllWorkByWorkID(db, int64(a))
	if err1 != nil {
		return c.JSON(404, map[string]interface{}{
			"status":  404,
			"message": "unable to get work by id",
		})
	}

	return c.JSON(200, map[string]interface{}{
		"status": 200,
		"data":   b,
	})
}

func GetWorkByUserId(db *sqlx.DB, c echo.Context) error {
	a, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"status":  500,
			"message": "unable to get user id",
		})
	}
	b, err1 := personalDB_api.GetAllWorkByWorkID(db, int64(a))
	if err1 != nil {
		return c.JSON(404, map[string]interface{}{
			"status":  404,
			"message": "unable to get work by user_id",
		})
	}

	return c.JSON(200, map[string]interface{}{
		"status": 200,
		"data":   b,
	})
}

func UpdateWork(db *sqlx.DB, c echo.Context) error {
	a, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"status":  500,
			"message": "unable to get user id",
		})
	}
	b, err1 := personalDB_api.GetAllWorkByWorkID(db, int64(a))
	if err1 != nil {
		return c.JSON(404, map[string]interface{}{
			"status":  404,
			"message": "unable to get work by user_id",
		})
	}
	d, err2 := personalDB_api.UpdateWork(db, *b)
	if err2 != nil {
		return c.JSON(404, map[string]interface{}{
			"status":  400,
			"message": "unable update work",
		})
	}
	return c.JSON(200, map[string]interface{}{
		"status": 200,
		"data":   d,
	})
}

func DeleteWork(db *sqlx.DB, c echo.Context) error {
	a, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"status":  500,
			"message": "unable to get work id",
		})
	}
	err1 := personalDB_api.DeleteWork(db, int64(a))
	if err1 != nil {
		return c.JSON(404, map[string]interface{}{
			"status":  404,
			"message": "unable to delete by work_id",
		})
	}

	return c.JSON(200, map[string]interface{}{
		"status":  200,
		"message": "work deleted"})
}
