package jwt

import (
	"fmt"
	"go-api-template/services/config"
	"log"
	"net/http"

	jwtlib "github.com/dgrijalva/jwt-go"
)

func Authorize(endpoint http.HandlerFunc) http.HandlerFunc {
	config := config.Get()

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {
			token, err := jwtlib.Parse(r.Header["Token"][0], func(token *jwtlib.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwtlib.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(config.SigningKey), nil
			})

			if err != nil {
				log.Print("Jwt authorization error: ", err)
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("401 - Unauthorized"))
			}

			if token.Valid {
				endpoint(w, r)
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("401 - Unauthorized"))
			return
		}
	}
}
