package main

import (
	"go-api-template/server"
	"go-api-template/services"
	"net/http"
)

func main() {
	config := services.GetConfig()

	httpServer := &server.HTTPServer{
		ID: &http.Server{Addr: config.HTTPPort, Handler: nil},
	}

	httpServer.Start()
}
