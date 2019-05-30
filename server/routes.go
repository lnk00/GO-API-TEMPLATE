package server

import (
	"go-api-template/controllers"
	"go-api-template/services/jwt"
	"net/http"
)

func (s *HTTPServer) initRoutes() {
	http.HandleFunc("/api/infos", controllers.HandleAPIInfos())
	http.HandleFunc("/api/loggedinfos", jwt.Authorize(controllers.HandleLoggedInfos()))
	http.HandleFunc("/api/signup", controllers.HandleSignup())
	http.HandleFunc("/api/signin", controllers.HandleSignin())
}
