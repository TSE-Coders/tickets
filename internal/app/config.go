package app

import "github.com/TSE-Coders/tickets/internal/store"

type AppConfig struct {
	Port        string
	StoreConfig store.DBConnectionConfig
}

func NewAppConfig() AppConfig {
	storeConfig := store.NewDBConnectionConfig()

	return AppConfig{
		Port:        "3000",
		StoreConfig: storeConfig,
	}
}

func (a AppConfig) WithPort(p string) AppConfig {
	a.Port = p
	return a
}
func (a AppConfig) WithStoreConfig(config store.DBConnectionConfig) AppConfig {
	a.StoreConfig = config
	return a
}
