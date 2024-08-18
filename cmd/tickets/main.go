package main

import (
	"log"

	"github.com/TSE-Coders/tickets/internal/app"
	"github.com/TSE-Coders/tickets/internal/config"
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
