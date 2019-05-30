package server

import (
	"go-api-template/services"
	"log"
	"net/http"
)

type HTTPServer struct {
	ID *http.Server
}

func (s *HTTPServer) Start() {
	services.StartDBMigration()
	s.initRoutes()

	err := s.ID.ListenAndServe()
	if err != http.ErrServerClosed {
		log.Fatalln(err)
	}
}
