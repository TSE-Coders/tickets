package server

import (
	"fmt"

	"github.com/TSE-Coders/tickets/internal/generator"
	"github.com/labstack/echo/v4"
)

type Http struct {
	config HttpConfig
}

type HttpConfig struct {
	Server    *echo.Echo
	Port      string
	Generator generator.Generator
}

func NewHttpConfig() HttpConfig {
	return HttpConfig{
		Server:    echo.New(),
		Port:      "3000",
		Generator: generator.NewGenerator(),
	}
}

func (h HttpConfig) WithPort(p string) HttpConfig {
	h.Port = p
	return h
}

func NewHttpServer(config HttpConfig) Http {
	h := Http{
		config,
	}

	h.config.Server.GET("/health-check", h.healthCheck)
	h.config.Server.GET("/tickets/random", h.getRandomTicket)
	return h
}

func (h Http) Run() error {
	port := fmt.Sprintf(":%s", h.config.Port)
	return h.config.Server.Start(port)
}
