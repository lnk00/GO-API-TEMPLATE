package server

import (
	"go-api-template/controllers"
	"net/http"
)

func (s *HTTPServer) initRoutes() {
	http.HandleFunc("/api/infos", controllers.HandleAPIInfos())
}
