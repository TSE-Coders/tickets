package main

import (
	"fmt"
	"net/http"

	"github.com/TSE-Coders/tickets/internal/config"
	"github.com/TSE-Coders/tickets/internal/generator"
	"github.com/labstack/echo/v4"
)

func main() {
	env := config.LoadEnvVars()
	g := generator.NewGenerator()
	httpServer := echo.New()

	httpServer.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "")
	})
	httpServer.GET("/tickets/random", func(c echo.Context) error {
		return c.JSON(200, g.GenetateRandomTicket())
	})

	httpServer.Logger.Fatal(httpServer.Start(fmt.Sprintf(":%s", env.PORT)))
}
