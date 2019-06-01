package server

import (
	"go-api-template/services/config"
	"go-api-template/services/db"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type HTTPServer struct {
	ID *http.Server
}

func (s *HTTPServer) Start() {
	config := config.Get()

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	setupCloseHandler()
	db.StartMigration()
	s.initRoutes()

	log.Printf("- Sever is running on port %s\n", config.HTTPPort)
	err := s.ID.ListenAndServe()
	if err != http.ErrServerClosed {
		log.Fatalln(err)
	}
}

func setupCloseHandler() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		db := db.Get()
		<-c
		log.Println("- Ctrl+C pressed, Server shutdown.")
		db.Close()
		os.Exit(0)
	}()
}
