package controllers

import (
	"encoding/json"
	"go-api-template/models"
	"go-api-template/services/db"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type SignupRequest struct {
	Email    string `json:email`
	Password string `json:password`
}

type SignupResponse struct {
	Message string `json:message`
}

func HandleSignup() http.HandlerFunc {
	db := db.Get()
	bodyRequest := &SignupRequest{}
	bodyResponse := &SignupResponse{}

	return func(w http.ResponseWriter, r *http.Request) {
		json.NewDecoder(r.Body).Decode(bodyRequest)

		passByte, err := bcrypt.GenerateFromPassword([]byte(bodyRequest.Password), 10)
		if err != nil {
			log.Print("Bcrypt hash password error: ", err)
		}

		dbErr := db.Create(&models.User{Email: bodyRequest.Email, Password: string(passByte)})
		if len(dbErr.GetErrors()) != 0 {
			bodyResponse.Message = "User not created"
		} else {
			bodyResponse.Message = "User created"
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(bodyResponse)
	}
}
