package config

import "os"

type envVars struct {
	PORT string
}

func LoadEnvVars() *envVars {
	e := &envVars{
		PORT: os.Getenv("PORT"),
	}

	return e
}
