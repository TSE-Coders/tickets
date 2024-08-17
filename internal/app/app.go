package app

import (
	"fmt"

	"github.com/TSE-Coders/tickets/internal/generator"
	"github.com/TSE-Coders/tickets/internal/scheduler"
	"github.com/labstack/echo/v4"
)

type App struct {
	config AppConfig
}

type AppConfig struct {
	Server         *echo.Echo
	Port           string
	Generator      generator.Generator
	BackgroundJobs []scheduler.Schedule
}

func NewAppConfig() AppConfig {
	return AppConfig{
		Server:    echo.New(),
		Port:      "3000",
		Generator: generator.NewGenerator(),
	}
}

func (a AppConfig) WithPort(p string) AppConfig {
	a.Port = p
	return a
}

func (a *App) AddBackgroundJob(job scheduler.Schedule) {
	a.config.BackgroundJobs = append(a.config.BackgroundJobs, job)
}

func NewAppServer(config AppConfig) App {
	a := App{
		config,
	}

	a.config.Server.GET("/health-check", a.healthCheck)
	a.config.Server.GET("/tickets/random", a.getRandomTicket)

	job := scheduler.New(15, true, func() error {
		t := a.config.Generator.GenetateRandomTicket()
		fmt.Printf("Ticket Created: %s\n", t.TicketID)

		return nil
	})

	a.AddBackgroundJob(*job)

	return a
}

func (a App) Run() error {
	for _, job := range a.config.BackgroundJobs {
		job.Run()
	}

	port := fmt.Sprintf(":%s", a.config.Port)
	return a.config.Server.Start(port)
}
