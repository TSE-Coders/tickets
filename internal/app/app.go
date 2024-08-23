package app

import (
	"fmt"

	"github.com/TSE-Coders/tickets/internal/generator"
	"github.com/TSE-Coders/tickets/internal/scheduler"
	"github.com/labstack/echo/v4"
)

type App struct {
	BackgroundJobs []scheduler.Schedule
	Generator      generator.Generator
	Server         *echo.Echo
	config         AppConfig
}

func NewApp(config AppConfig) (App, error) {
	a := App{
		config: config,
	}

	// Setup App's Background Job
	for count := range 2 {
		a.addBackgroundJob(
			*scheduler.New(30, true, a.ticketGeneratorBackgroundJob(count)),
		)
	}

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
	a.Server.GET("/api/health-check", a.healthCheck)
	a.Server.GET("/api/tickets/game", a.getGameTicket)
	a.Server.POST("/api/tickets", a.submitTicket)

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

func (a *App) addBackgroundJob(job scheduler.Schedule) {
	a.BackgroundJobs = append(a.BackgroundJobs, job)
}

func (a *App) ticketGeneratorBackgroundJob(jobId int) func() error {
	return func() error {
		t, err := a.Generator.GenetateRandomTicket()
		if err != nil {
			fmt.Printf("failed to generate ticket: %s\n", err.Error())
			return err
		}
		fmt.Printf("Job %d: Ticket Created: %d %+v\n", jobId, a.Generator.TicketCount, t)
		return nil
	}
}
