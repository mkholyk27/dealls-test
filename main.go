package main

import (
	"dating-app/configs"
	"dating-app/routes"
)

func main() {
	// init configuration
	cfg := configs.New()

	// init server
	server := configs.NewServer(&cfg)
	srv := server.InitServer()

	// init routes
	routes := routes.New(srv)
	routes.GenerateRoutes()

	// start server
	server.Start()
}
