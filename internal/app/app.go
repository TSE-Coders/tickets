package app

import (
	"fmt"

	"github.com/TSE-Coders/tickets/internal/generator"
	"github.com/TSE-Coders/tickets/internal/scheduler"
	"github.com/TSE-Coders/tickets/internal/store"
	"github.com/labstack/echo/v4"
)

type App struct {
	config       AppConfig
	DBConnection store.DB
}

type AppConfig struct {
	Server         *echo.Echo
	Port           string
	Generator      generator.Generator
	BackgroundJobs []scheduler.Schedule
	StoreConfig    store.DBConnectionConfig
}

func NewAppConfig() AppConfig {
	storeConfig := store.NewDBConnectionConfig()

	return AppConfig{
		Server:      echo.New(),
		Port:        "3000",
		Generator:   generator.NewGenerator(),
		StoreConfig: storeConfig,
	}
}

func (a AppConfig) WithPort(p string) AppConfig {
	a.Port = p
	return a
}

func (a *App) AddBackgroundJob(job scheduler.Schedule) {
	a.config.BackgroundJobs = append(a.config.BackgroundJobs, job)
}

func NewAppServer(config AppConfig) (App, error) {
	// Setup App's Database
	dbConnection, err := store.NewDBConnection(config.StoreConfig)
	if err != nil {
		return App{}, err
	}

	a := App{
		config:       config,
		DBConnection: *dbConnection,
	}

	// Setup App's HTTP Server handlers
	a.config.Server.GET("/health-check", a.healthCheck)
	a.config.Server.GET("/tickets/random", a.getRandomTicket)

	// Setup App's background job
	job := scheduler.New(15, true, func() error {
		t := a.config.Generator.GenetateRandomTicket()
		fmt.Printf("Ticket Created: %s\n", t.TicketID)
		return nil
	})
	a.AddBackgroundJob(*job)

	return a, nil
}

func (a App) Run() error {
	fmt.Println("Starting Application...")

	for _, job := range a.config.BackgroundJobs {
		job.Run()
	}

	port := fmt.Sprintf(":%s", a.config.Port)
	return a.config.Server.Start(port)
}
