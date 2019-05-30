package server

import (
	"go-api-template/services/db"
	"log"
	"net/http"
)

type HTTPServer struct {
	ID *http.Server
}

func (s *HTTPServer) Start() {
	db.StartMigration()
	s.initRoutes()

	err := s.ID.ListenAndServe()
	if err != http.ErrServerClosed {
		log.Fatalln(err)
	}
}
