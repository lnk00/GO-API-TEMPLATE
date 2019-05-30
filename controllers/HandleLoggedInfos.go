package controllers

import (
	"fmt"
	"go-api-template/services/config"
	"net/http"
)

func HandleLoggedInfos() http.HandlerFunc {
	config := config.Get()
	infos := fmt.Sprintf("API name: %s\nAPI version: %s\nHTTP port: %s\nDB type: %s\nDB host: %s\nDB port: %s\nDB name: %s\nDB user: %s\n",
		config.APIName, config.APIVersion, config.HTTPPort, config.DBType, config.DBHost, config.DBPort, config.DBName, config.DBUser)

	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(infos))
	}
}
