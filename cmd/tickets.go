package main

import (
	"embed"
	"log"

	"github.com/TSE-Coders/tickets/internal/app"
	"github.com/TSE-Coders/tickets/internal/config"
)

var (
	// //go:embed embed/migrations/*.sql
	migrations embed.FS
)

func main() {
	env := config.LoadEnvVars()
	appConfig := app.NewAppConfig()
	appConfig = appConfig.WithPort(env.PORT)

	app, err := app.NewAppServer(appConfig)
	if err != nil {
		log.Fatalf("Failed to initialize the Application: %s", err)
	}

	if err := app.Run(); err != nil {
		log.Fatalf("Application Crashed: %s", err)
	}
}
