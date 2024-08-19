package main

import (
	"log"

	"github.com/TSE-Coders/tickets/internal/app"
	"github.com/TSE-Coders/tickets/internal/config"
	"github.com/TSE-Coders/tickets/internal/store"
)

func main() {
	env := config.LoadEnvVars()
	storeConfig := store.NewDBConnectionConfig()

	appConfig := app.NewAppConfig(storeConfig)
	appConfig = appConfig.WithPort(env.PORT)

	app, err := app.NewApp(appConfig)
	if err != nil {
		log.Fatalf("Failed to initialize the Application: %s", err)
	}

	if err := app.Run(); err != nil {
		log.Fatalf("Application Crashed: %s", err)
	}
}
