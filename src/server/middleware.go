package main

import (
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"
)

//https://gowebexamples.com/basic-middleware/
func Authorize(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//Get cookies
		JWT, err := r.Cookie("JWT")

		if err != nil {
			panic(err)
		}

		SignTime, err := r.Cookie("Time")

		if err != nil {
			panic(err)
		}

		/*
		   if no cookies exist, it means that the person
		   who did the request did not login so we cannot
		   proceed with the request
		*/
		if JWT.Value == "" || SignTime.Value == "" {

			w.WriteHeader(403)
			w.Write([]byte("Unauthorized"))
			return
		}
		claims := jwt.MapClaims{}

		_, err = jwt.ParseWithClaims(JWT.Value, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("KEY")), nil
		})

		if err != nil {
			panic(err)
		}

		if claims["signtime"] != SignTime.Value {
			w.WriteHeader(403)
			w.Write([]byte("Unauthorized"))
			return
		}

		handler(w, r)
	}
}
