package app

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (a App) healthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func (a App) getRandomTicket(c echo.Context) error {
	return c.JSON(200, a.config.Generator.GenetateRandomTicket())
}
