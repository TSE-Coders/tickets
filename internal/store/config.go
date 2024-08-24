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
	Port            string
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
		Port:            "5432",
		QueryBufferSize: 10,
		User:            "admin",
	}

	// Checking for Environment Variables
	defaultConfig = defaultConfig.WithDatabaseDriver(checkEnv(defaultConfig.DatabaseDriver, "DATABASE_DRIVER"))
	defaultConfig = defaultConfig.WithDatabaseName(checkEnv(defaultConfig.DatabaseName, "DATABASE_NAME"))
	defaultConfig = defaultConfig.WithHost(checkEnv(defaultConfig.Host, "DATABASE_HOST"))
	defaultConfig = defaultConfig.WithPassword(checkEnv(defaultConfig.Password, "DATABASE_PASSWORD"))
	defaultConfig = defaultConfig.WithPort(checkEnv(defaultConfig.Port, "DATABASE_PORT"))
	defaultConfig = defaultConfig.WithUser(checkEnv(defaultConfig.User, "DATABASE_USER"))
	envDatabaseQueryBufferSize := os.Getenv("DATABASE_QUERY_BUFFER_SIZE")
	if envDatabaseQueryBufferSize != "" {
		sizeInt, err := strconv.Atoi(envDatabaseQueryBufferSize)
		if err != nil {
			fmt.Printf("failed to convert environment variable 'DATABASE_QUERY_BUFFER_SIZE' to an integer, using the default: %s", err.Error())
		} else {
			fmt.Println("found environment variable: DATABASE_QUERY_BUFFER_SIZE")
			defaultConfig = defaultConfig.WithQueryBufferSize(uint8(sizeInt))
		}
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
func (config DBConnectionConfig) WithPort(port string) DBConnectionConfig {
	config.Port = port
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

func checkEnv(currentValue, envVariable string) string {
	value := os.Getenv(envVariable)
	if value != "" {
		fmt.Printf("found environment variable: %s\n", envVariable)
		return value
	}
	return currentValue
}
