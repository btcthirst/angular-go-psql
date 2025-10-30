package api

import (
	"github.com/btcthirst/angular-go-psql/internal/config"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

var addr string

func Initialize() error {
	config.InitEnv()
	port := viper.GetString("SERVER_PORT")
	if port == "" {
		port = "3333" // Default port
	}
	addr = ":" + port
	// Initialize the API
	// This function can be used to set up routes, middleware, etc.
	return nil
}

func Start() {
	// Start the API server
	// This function can be used to start the HTTP server and listen for requests
	e := echo.New()

	e.Logger.Fatal(e.Start(addr))
}
func Shutdown() {
	// Shutdown the API server
	// This function can be used to gracefully shut down the server
}
