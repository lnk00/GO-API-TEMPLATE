package controllers

import (
	"go-api-template/services"
	"net/http"
)

func HandleAPIInfos() http.HandlerFunc {
	config := services.GetConfig()

	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API name: " + config.APIName +
			"\nAPI version: " + config.APIVersion))
	}
}
