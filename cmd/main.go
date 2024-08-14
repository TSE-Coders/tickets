package main

import (
	"fmt"
	"net/http"

	"github.com/TSE-Coders/tickets/internal/env"
	"github.com/labstack/echo/v4"
)

func init() {
	env.LoadEnvVars()
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", env.Env.PORT)))
}
