package handlers

import "github.com/labstack/echo/v4"

func UsersHandler(c echo.Context) error {
	c.Render(200, "home", nil)
	return nil
}
