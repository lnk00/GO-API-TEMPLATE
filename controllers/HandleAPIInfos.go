package controllers

import (
	"fmt"
	"go-api-template/services/config"
	"net/http"
)

func HandleAPIInfos() http.HandlerFunc {
	config := config.Get()
	infos := fmt.Sprintf("API name: %s\nAPI version: %s\n",
		config.APIName, config.APIVersion)

	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(infos))
	}
}
