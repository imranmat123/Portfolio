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

	e.GET("/usersAll", func(e echo.Context) error {
		return handlers.GetAllUsers(db, e)
	})
	e.GET("/users/:id", func(e echo.Context) error {
		return handlers.GetUserById(db, e)
	})
	e.GET("/users/profilePicture/:id", func(e echo.Context) error { return handlers.GetProfilePicture(db, e) })
	e.GET("/users/aboutMe/:id", func(e echo.Context) error { return handlers.GetAboutMe(db, e) })
	e.POST("/usersCreate", func(e echo.Context) error { return handlers.CreateUser(db, e) })
	e.PUT("/users/:id", func(e echo.Context) error {
		return handlers.UpdateUser(db, e)
	})
	e.DELETE("/users/:id", func(e echo.Context) error {
		return handlers.DeleteUser(db, e)
	})

	//projects
	e.POST("/projects/create", func(e echo.Context) error { return handlers.CreateProject(db, e) })
	e.GET("/projects/GetAllProjects", func(e echo.Context) error { return handlers.GetAllProjects(db, e) })
	e.GET("/projects/GetProject/:id", func(e echo.Context) error { return handlers.GetProjectById(db, e) })
	e.GET("/projects/GetAboutMe/:id", func(e echo.Context) error { return handlers.GetByABoutMe(db, e) })
	e.GET("/projects/GetProjectURL/:id", func(e echo.Context) error { return handlers.GetProjectURL(db, e) })
	e.GET("/projects/GetProjectName/:id", func(e echo.Context) error { return handlers.GetProjectName(db, e) })
	e.GET("/projects/GetProjectsByUserID/:id", func(e echo.Context) error { return handlers.GetProjectsByUserID(db, e) })
	e.GET("/projects/GetAllUserProjects/:id", func(e echo.Context) error { return handlers.GetAllUserProjects(db, e) })
	e.PUT("/projects/UpdateProject/:id", func(e echo.Context) error { return handlers.UpdateProject(db, e) })
	e.DELETE("/projects/DeleteProject/:id", func(e echo.Context) error { return handlers.DeleteProject(db, e) })
}
