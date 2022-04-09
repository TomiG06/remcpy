package main

import (
	"net/http"
    "os"
    "strconv"
    "fmt"
    "github.com/joho/godotenv"
)

func main() {
    godotenv.Load(".env")

    port := os.Getenv("PORT")

    //Check if PORT is given by the user
    if port == "" {
        fmt.Fprintf(os.Stderr, "PORT not provided\n")
        os.Exit(1)
    }

    temp, err := strconv.Atoi(port)

    //Error on atoi means that PORT contains non-digit characters
    if err != nil {
        fmt.Fprintf(os.Stderr, "PORT must only contain digits\n")
        os.Exit(1)
    }

    //Check if PORT is greater than 1024
    if temp <= 1024 {
        fmt.Fprintf(os.Stderr, "PORT must be greater than 1024\nPORT given: %d\n", temp)
        os.Exit(1)
    }

    http.Handle("/login/", http.StripPrefix("/login/", http.FileServer(http.Dir("public/login"))))
    http.Handle("/main/", http.StripPrefix("/main/", http.FileServer(http.Dir("public/main"))))
    http.HandleFunc("/verify", VerificationHandler)
    http.Handle("/api", Authorize(APIHandler))
    http.ListenAndServe(":" + port, nil)

    fmt.Println("PORT already in use")
}
