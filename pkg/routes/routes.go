package routes

import (
	"ImransProfoiloWebsite/internal/handlers"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(db *sqlx.DB, e *echo.Echo) {
	//home
	e.GET("/", handlers.HomeHandler)

	//user
	e.GET("/users",
		func(c echo.Context) error {
			err := handlers.UsersHandler(db, c)

			if err != nil {
				return err
			}
			return c.Render(200, "user", nil)
		})

	e.GET("/usersAll", func(e echo.Context) error {
		return handlers.GetAllUsers(db, e)
	})

	e.GET("/users/:id", func(e echo.Context) error {
		return handlers.GetUserById(db, e)
	})

	e.GET("/users/profilePicture/:id", func(e echo.Context) error {
		return handlers.GetProfilePicture(db, e)
	})

	e.GET("/users/aboutMe/:id", func(e echo.Context) error { return handlers.GetAboutMe(db, e) })

	e.POST("/usersCreate", func(e echo.Context) error {
		return handlers.CreateUser(db, e)
	})

	e.PUT("/users/:id", func(e echo.Context) error {
		return handlers.UpdateUser(db, e)
	})

	e.DELETE("/users/:id", func(e echo.Context) error {
		return handlers.DeleteUser(db, e)
	})
}
