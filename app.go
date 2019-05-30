package main

import (
	"go-api-template/server"
	"go-api-template/services/config"
	"net/http"
)

func main() {
	config := config.Get()

	httpServer := &server.HTTPServer{
		ID: &http.Server{Addr: config.HTTPPort, Handler: nil},
	}

	httpServer.Start()
}
