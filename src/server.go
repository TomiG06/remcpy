package main

import (
	"net/http"
    "github.com/joho/godotenv"
)

func main() {
    godotenv.Load(".env")

	http.Handle("/login/", http.StripPrefix("/login/", http.FileServer(http.Dir("public/login"))))
	http.Handle("/main/", http.StripPrefix("/main/", http.FileServer(http.Dir("public/main"))))
	http.HandleFunc("/verify", VerificationHandler)
	http.Handle("/api", Authorize(APIHandler))
	http.ListenAndServe(":8080", nil)
}
