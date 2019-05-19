package server

import (
	"log"
	"net/http"
)

type HTTPServer struct {
	ID *http.Server
}

func (s *HTTPServer) Start() {
	s.initRoutes()

	err := s.ID.ListenAndServe()
	if err != http.ErrServerClosed {
		log.Fatalln(err)
	}
}
