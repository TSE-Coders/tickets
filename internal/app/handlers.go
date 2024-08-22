package app

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (a App) healthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func (a *App) getRandomTicket(c echo.Context) error {
	ticket, err := a.Generator.GenetateRandomTicket()
	if err != nil {
		return c.JSON(500, struct {
			Message string
		}{
			Message: "failed to generate ticket",
		})
	}
	return c.JSON(200, ticket)
}
