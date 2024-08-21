package app

import (
	"fmt"

	"github.com/TSE-Coders/tickets/internal/generator"
	"github.com/TSE-Coders/tickets/internal/scheduler"
	"github.com/labstack/echo/v4"
)

type App struct {
	config         AppConfig
	BackgroundJobs []scheduler.Schedule
	Generator      generator.Generator
	Server         *echo.Echo
}

func NewApp(config AppConfig) (App, error) {
	a := App{
		config: config,
	}

	// Setup App's Background Job
	job := scheduler.New(3, true, func() error {
		t := a.Generator.GenetateRandomTicket()
		fmt.Printf("Ticket Created: %+v\n", t)
		return nil
	})
	a.AddBackgroundJob(*job)

	// Setup App's Ticket Generator
	g, err := generator.NewGenerator(a.config.StoreConfig)
	if err != nil {
		return a, err
	}
	a.Generator = g

	// Setup App's HTTP Server handlers
	a.Server = echo.New()
	a.Server.HideBanner = true
	a.Server.HidePort = true
	a.Server.GET("/health-check", a.healthCheck)
	a.Server.GET("/tickets/random", a.getRandomTicket)

	return a, nil
}

func (a App) Run() error {
	fmt.Printf("Starting Application on port %s\n", a.config.Port)

	for _, job := range a.BackgroundJobs {
		job.Run()
	}

	port := fmt.Sprintf(":%s", a.config.Port)
	return a.Server.Start(port)
}

func (a *App) AddBackgroundJob(job scheduler.Schedule) {
	a.BackgroundJobs = append(a.BackgroundJobs, job)
}
