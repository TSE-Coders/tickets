package main

import (
	"fmt"
	"log"

	"github.com/TSE-Coders/tickets/internal/app"
	"github.com/TSE-Coders/tickets/internal/store"
)

func main() {
	storeConfig := store.NewDBConnectionConfig()
	appConfig := app.NewAppConfig(storeConfig)

	app, err := app.NewApp(appConfig)
	if err != nil {
		log.Fatalf("Failed to initialize the Application: %s", err)
	}

	fmt.Println("Yoooo")

	if err := app.Run(); err != nil {
		log.Fatalf("Application Crashed: %s", err)
	}
}
