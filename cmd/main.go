package main

import (
	"log"

	"github.com/TSE-Coders/tickets/internal/config"
	"github.com/TSE-Coders/tickets/internal/server"
)

func main() {
	env := config.LoadEnvVars()
	appConfig := server.NewHttpConfig().
		WithPort(env.PORT)

	app := server.NewHttpServer(appConfig)
	if err := app.Run(); err != nil {
		log.Fatalf("Application Crashed: %s", err)
	}
}
