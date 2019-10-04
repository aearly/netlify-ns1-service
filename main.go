package main

import (
	"fmt"
	"github.com/aearly/netlify-ns1-service/openapi"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	api "gopkg.in/ns1/ns1-go.v2/rest"
	"net/http"
	"os"
	"time"
)

type CustomContext struct {
	echo.Context
	Ns1Client *api.Client
}

func clientMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		httpClient := &http.Client{Timeout: time.Second * 10}
		apiKey := os.Getenv("NS1_APIKEY")
		client := api.NewClient(httpClient, api.SetAPIKey(apiKey))
		ctx := &CustomContext{
			Context:   c,
			Ns1Client: client,
		}
		return next(ctx)
	}
}

func main() {
	// Echo instance
	server := echo.New()

	// Middleware
	server.Use(middleware.Logger())
	server.Use(middleware.Recover())
	server.Use(clientMiddleware)

	apiKey := os.Getenv("NS1_APIKEY")
	if apiKey == "" {
		fmt.Println("error: NS1_APIKEY environment variable is not set")
		os.Exit(1)
	}

	handlers := Handlers{
		apiKey: apiKey,
	}

	// Routes
	ServiceOapi.RegisterHandlers(server, handlers)

	// Start server
	server.Logger.Info(server.Start(":5000"))
}
