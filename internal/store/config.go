package store

import (
	"fmt"
	"os"
	"strconv"
)

type DBConnectionConfig struct {
	DatabaseDriver  string
	DatabaseName    string
	Host            string
	Password        string
	QueryBufferSize uint8
	User            string
}

func NewDBConnectionConfig() DBConnectionConfig {
	// Create the default configuration
	defaultConfig := DBConnectionConfig{
		DatabaseDriver:  "postgres",
		DatabaseName:    "app",
		Host:            "localhost",
		Password:        "password",
		QueryBufferSize: 10,
		User:            "admin",
	}

	// Checking for Environment Variables
	envDatabaseDriver := os.Getenv("DATABASE_DRIVER")
	if envDatabaseDriver != "" {
		defaultConfig = defaultConfig.WithDatabaseDriver(envDatabaseDriver)
	}
	envDatabaseName := os.Getenv("DATABASE_NAME")
	if envDatabaseName != "" {
		defaultConfig = defaultConfig.WithDatabaseName(envDatabaseName)
	}
	envDatabaseHost := os.Getenv("DATABASE_HOST")
	if envDatabaseHost != "" {
		defaultConfig = defaultConfig.WithHost(envDatabaseHost)
	}
	envDatabasePassword := os.Getenv("DATABASE_PASSWORD")
	if envDatabasePassword != "" {
		defaultConfig = defaultConfig.WithPassword(envDatabasePassword)
	}
	envDatabaseQueryBufferSize := os.Getenv("DATABASE_QUERY_BUFFER_SIZE")
	if envDatabaseQueryBufferSize != "" {
		sizeInt, err := strconv.Atoi(envDatabaseQueryBufferSize)
		if err != nil {
			fmt.Printf("failed to convert environment variable 'DATABASE_QUERY_BUFFER_SIZE' to an integer, using the default: %q", err)
		} else {
			defaultConfig = defaultConfig.WithQueryBufferSize(uint8(sizeInt))
		}
	}
	envDatabaseUser := os.Getenv("DATABASE_USER")
	if envDatabaseUser != "" {
		defaultConfig = defaultConfig.WithUser(envDatabaseUser)
	}

	return defaultConfig
}

func (config DBConnectionConfig) WithDatabaseDriver(driver string) DBConnectionConfig {
	config.DatabaseDriver = driver
	return config
}
func (config DBConnectionConfig) WithDatabaseName(name string) DBConnectionConfig {
	config.DatabaseName = name
	return config
}
func (config DBConnectionConfig) WithHost(host string) DBConnectionConfig {
	config.Host = host
	return config
}
func (config DBConnectionConfig) WithPassword(password string) DBConnectionConfig {
	config.Password = password
	return config
}
func (config DBConnectionConfig) WithQueryBufferSize(size uint8) DBConnectionConfig {
	config.QueryBufferSize = size
	return config
}
func (config DBConnectionConfig) WithUser(user string) DBConnectionConfig {
	config.User = user
	return config
}
