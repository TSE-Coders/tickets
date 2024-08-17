package main

import (
	"fmt"
	"log"

	"github.com/TSE-Coders/tickets/internal/app"
	"github.com/TSE-Coders/tickets/internal/config"
	"github.com/TSE-Coders/tickets/internal/scheduler"
)

func main() {
	env := config.LoadEnvVars()
	appConfig := app.NewAppConfig().WithPort(env.PORT)
	app := app.NewAppServer(appConfig)

	job1 := scheduler.New(5, true, func() error {
		fmt.Println("Hello World from job 1")
		return nil
	})
	app.AddBackgroundJob(*job1)

	if err := app.Run(); err != nil {
		log.Fatalf("Application Crashed: %s", err)
	}
}
