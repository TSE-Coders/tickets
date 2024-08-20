package app

import (
	"fmt"
	"os"

	"github.com/TSE-Coders/tickets/internal/store"
)

type AppConfig struct {
	Port        string
	StoreConfig store.DBConnectionConfig
}

func NewAppConfig(storeConfig store.DBConnectionConfig) AppConfig {
	defaultConfig := AppConfig{
		Port:        "3000",
		StoreConfig: storeConfig,
	}
	defaultConfig = defaultConfig.WithPort(checkEnv(defaultConfig.Port, "PORT"))

	return defaultConfig
}

func (a AppConfig) WithPort(p string) AppConfig {
	a.Port = p
	return a
}
func (a AppConfig) WithStoreConfig(config store.DBConnectionConfig) AppConfig {
	a.StoreConfig = config
	return a
}

func checkEnv(currentValue, envVariable string) string {
	value := os.Getenv(envVariable)
	if value != "" {
		fmt.Printf("found environment variable: %s\n", envVariable)
		return value
	}
	return currentValue
}
