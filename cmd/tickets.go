package main

import (
	"log"

	"github.com/TSE-Coders/tickets/internal/app"
	"github.com/TSE-Coders/tickets/internal/config"
)

func main() {
	env := config.LoadEnvVars()
	appConfig := app.NewAppConfig().WithPort(env.PORT)
	app := app.NewAppServer(appConfig)

	if err := app.Run(); err != nil {
		log.Fatalf("Application Crashed: %s", err)
	}
}
