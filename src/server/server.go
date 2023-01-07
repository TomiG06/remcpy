package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("../.env")

	port := os.Getenv("PORT")

	//Check if PORT is given by the user
	fmt.Println(port)
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
    
    //<At this particular point I have no idea how this works
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../public/static/"))))
    // />

	http.Handle("/login/", http.StripPrefix("/login/", http.FileServer(http.Dir("../public/login"))))
	http.Handle("/main/", http.StripPrefix("/main/", http.FileServer(http.Dir("../public/main"))))
	http.HandleFunc("/verify", VerificationHandler)
	http.Handle("/api", Authorize(APIHandler))
	http.ListenAndServe(":"+port, nil)

	//if Listen fails it probably means that something runs on the PORT given
	fmt.Printf("PORT %v already in use", port)
}
