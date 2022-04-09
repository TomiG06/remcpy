package main

import(
    "net/http"
    "os"
    "github.com/golang-jwt/jwt"
)

//https://gowebexamples.com/basic-middleware/
func Authorize(handler http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {

        JWT, err := r.Cookie("JWT")

        if err != nil {
            panic(err)
        }

        SignTime, err := r.Cookie("Time")

        if err != nil {
            panic(err)
        }

        if JWT.Value == "" || SignTime.Value == "" {

            w.WriteHeader(403)
            w.Write([]byte("Unauthorized"))
        } else {
            claims := jwt.MapClaims{}

            _, err := jwt.ParseWithClaims(JWT.Value, claims, func(token *jwt.Token) (interface{}, error) {
                return []byte(os.Getenv("KEY")), nil
            })

            if err != nil {
                panic(err)
            }

            if claims["signtime"] != SignTime.Value {
                w.WriteHeader(403)
                w.Write([]byte("Unauthorized"))
            } else {
                handler(w, r)
            }
        }
    }
}