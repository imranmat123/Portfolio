package routes

import (
	"ImransProfoiloWebsite/handlers"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	e.GET("/", handlers.HomeHandler)
	e.GET("/users", handlers.UsersHandler)
}
