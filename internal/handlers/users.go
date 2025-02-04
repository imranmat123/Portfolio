package handlers

import (
	"ImransProfoiloWebsite/internal/models"
	personalDB_api "ImransProfoiloWebsite/internal/personalDB-api"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"strconv"
)

func UsersHandler(db *sqlx.DB, c echo.Context) error {
	if db != nil {
		return c.JSON(500, map[string]interface{}{
			"status":  500,
			"message": "Database connection unavailable.",
		})
	}
	return c.Render(200, "user", nil)
}

func CreateUser(db *sqlx.DB, c echo.Context) error {
	a := models.User{}
	err := c.Bind(&a)

	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"status":  400,
			"message": "We've not been able to get the user",
		})
	}
	err1 := personalDB_api.CreateUser(db, a)

	if err1 != nil {
		return c.JSON(500, map[string]interface{}{
			"status":  500,
			"message": "could not create user",
		})
	}
	return c.JSON(201, map[string]interface{}{
		"status":  201,
		"message": "user has been created",
	})

}

func GetAllUsers(db *sqlx.DB, c echo.Context) error {
	a, err := personalDB_api.GetUser(db)
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"status":  400,
			"message": "we were not able to get all users",
		})
	}
	return c.JSON(200, map[string]interface{}{
		"status": 200,
		"data":   a,
	})

}

func GetUserById(db *sqlx.DB, c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"status":  400,
			"message": "the ID hasn't been converted to an int",
		})
	}
	a, err1 := personalDB_api.GetUserById(db, int64(id))

	if err1 != nil {
		return c.JSON(404, map[string]interface{}{
			"status":  404,
			"message": "no user matching this ID",
		})
	}
	return c.JSON(200, map[string]interface{}{
		"status": 200,
		"data":   a,
	})
}

func UpdateUser(db *sqlx.DB, c echo.Context) error {
	a, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"status":  400,
			"message": "ID has not been converted",
		})
	}

	_, err1 := personalDB_api.GetUserById(db, int64(a))

	if err1 != nil {
		return c.JSON(404, map[string]interface{}{
			"status":  404,
			"message": "user not found",
		})
	}

	ur := models.GetUser{}

	err2 := c.Bind(&ur)
	if err2 != nil {
		return c.JSON(400, map[string]interface{}{
			"status":  400,
			"message": "we have not been able to bind",
		})
	}

	err3 := personalDB_api.UpdateUser(db, ur)

	if err3 != nil {
		return c.JSON(500, map[string]interface{}{
			"status":  500,
			"message": "user not found",
		})
	}
	return c.JSON(200, map[string]interface{}{
		"status": 200,
		"data":   ur,
	})
}

func DeleteUser(db *sqlx.DB, c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"status":  400,
			"message": "unable to convert ID",
		})
	}

	_, err1 := personalDB_api.GetUserById(db, int64(id))

	if err1 != nil {
		return c.JSON(200, map[string]interface{}{
			"status":  200,
			"message": "User has been deleted",
		})
	}

	err2 := personalDB_api.DeleteUser(db, int64(id))

	if err2 != nil {
		return c.JSON(200, map[string]interface{}{
			"status":  200,
			"message": "user has been deleted from the database",
		})
	}
	return c.JSON(400, map[string]interface{}{
		"status":  400,
		"message": "unable to delete user",
	})

}
