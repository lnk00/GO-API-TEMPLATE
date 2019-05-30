package jwt

import (
	"go-api-template/services/config"
	"log"
	"time"

	jwtlib "github.com/dgrijalva/jwt-go"
)

func Generate() (string, error) {
	config := config.Get()
	token := jwtlib.New(jwtlib.SigningMethodHS256)

	claims := token.Claims.(jwtlib.MapClaims)
	claims["authorized"] = true
	claims["exp"] = time.Now().Add(time.Minute * 2).Unix()

	tokenString, err := token.SignedString([]byte(config.SigningKey))
	if err != nil {
		log.Print("Jwt token generation error: ", err)
		return "", err
	}

	return tokenString, nil
}
