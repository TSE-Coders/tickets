package app

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/TSE-Coders/tickets/internal/types"
	"github.com/labstack/echo/v4"
)

func (a App) healthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func (a *App) getRandomTicket(c echo.Context) error {
	ticket, err := a.Generator.GenetateGameTicket()
	if err != nil {
		return c.JSON(500, struct {
			Message string
		}{
			Message: "failed to generate ticket",
		})
	}
	return c.JSON(http.StatusOK, ticket)
}

func (a *App) submitTicket(c echo.Context) error {
	ticket := types.Ticket{}
	if err := c.Bind(&ticket); err != nil {
		return err
	}
	id, err := a.Generator.Store.InsertTicket(ticket)
	if err != nil {
		return fmt.Errorf("failed to insert game ticket: %s", err.Error())
	}

	ticket = ticket.WithId(strconv.Itoa(id))
	fmt.Printf("The ticket created via game: %+v\n", ticket)

	return c.JSON(http.StatusCreated, ticket)
}
