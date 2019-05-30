package controllers

import (
	"encoding/json"
	"go-api-template/models"
	"go-api-template/services/db"
	"go-api-template/services/jwt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type SigninRequest struct {
	Email    string `json:email`
	Password string `json:password`
}

type SigninResponse struct {
	Message string `json:message`
	Token   string `json:token`
}

func HandleSignin() http.HandlerFunc {
	db := db.Get()
	bodyRequest := &SigninRequest{}
	bodyResponse := &SigninResponse{}

	return func(w http.ResponseWriter, r *http.Request) {
		json.NewDecoder(r.Body).Decode(bodyRequest)

		user := &models.User{}
		db.Where(map[string]interface{}{"email": bodyRequest.Email}).First(&user)

		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(bodyRequest.Password))
		if err != nil {
			bodyResponse.Message = "User not signed in"
			bodyResponse.Token = ""
		} else {
			bodyResponse.Message = "User signed in"
			token, err := jwt.Generate()
			if err == nil {
				bodyResponse.Token = token
			} else {
				bodyResponse.Token = ""
			}
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(bodyResponse)
	}
}
