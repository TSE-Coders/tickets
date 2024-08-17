package config

import "os"

type envVars struct {
	PORT          string
	POSTGRES_URL  string
	POSTGRES_PORT string
	POSTGRES_DB   string
	POSTGRES_USER string
}

func LoadEnvVars() *envVars {
	e := &envVars{
		PORT:          os.Getenv("PORT"),
		POSTGRES_URL:  os.Getenv("POSTGRES_URL"),
		POSTGRES_PORT: os.Getenv("POSTGRES_PORT"),
		POSTGRES_DB:   os.Getenv("acastro"),
		POSTGRES_USER: os.Getenv("acastro"),
	}

	return e
}
