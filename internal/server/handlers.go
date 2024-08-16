package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Http) healthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func (h Http) getRandomTicket(c echo.Context) error {
	return c.JSON(200, h.config.Generator.GenetateRandomTicket())
}
