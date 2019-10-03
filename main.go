package main

import (
	"github.com/aearly/netlify-ns1-service/openapi"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	server := echo.New()

	// Middleware
	server.Use(middleware.Logger())
	server.Use(middleware.Recover())

	handlers := Handlers{}

	// Routes
	ServiceOapi.RegisterHandlers(server, handlers)

	// Start server
	server.Logger.Info(server.Start(":5000"))
}
