package controllers

import (
	"go-api-template/services/config"
	"net/http"
)

func HandleAPIInfos() http.HandlerFunc {
	config := config.Get()

	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API name: " + config.APIName +
			"\nAPI version: " + config.APIVersion))
	}
}
