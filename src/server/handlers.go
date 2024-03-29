package main

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/atotto/clipboard"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func VerificationHandler(w http.ResponseWriter, r *http.Request) {

	//Check if request method is POST
	if r.Method != http.MethodPost {
		w.WriteHeader(405)
		w.Write([]byte("POST only"))
		return
	}
	r.ParseForm()

	pass := r.FormValue("password")

	//Check if password is given
	if pass == "" {
		w.WriteHeader(404)
		w.Write([]byte("password not provided\n"))
		return
	}
	//Open the file that contains our hashed password
	hashed, err := os.ReadFile("../password.txt")

	//Check for errors
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("server error"))
		return
	}

	//Compare hash with password given
	if bcrypt.CompareHashAndPassword(hashed, []byte(pass)) != nil {
		w.WriteHeader(401)
		w.Write([]byte("wrong password"))
		return
	}
	//Generate a JWT if password is correct
	SignTime := strconv.FormatInt(time.Now().Unix(), 10)

	claims := &jwt.MapClaims{
		"auth":     true,
		"signtime": SignTime,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenstr, err := token.SignedString([]byte(os.Getenv("KEY")))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("server error"))
		return
	}

	JWTCookie := http.Cookie{
		Name:   "JWT",
		Value:  tokenstr,
		MaxAge: 86400,
		Path:   "/",
	}

	TimeCookie := http.Cookie{
		Name:   "Time",
		Value:  SignTime,
		MaxAge: 86400,
		Path:   "/",
	}

	http.SetCookie(w, &JWTCookie)
	http.SetCookie(w, &TimeCookie)
	http.Redirect(w, r, "/main", http.StatusSeeOther)
}

func APIHandler(w http.ResponseWriter, r *http.Request) {

	//Check if method is not POST
	if r.Method != http.MethodPost {
		w.WriteHeader(405)
		w.Write([]byte("POST only"))
		return
	}
	r.ParseForm()

	//Text to be copied
	text := r.FormValue("text")

	if text == "" {
		w.WriteHeader(400)
		w.Write([]byte("text value can't be empty"))
		return
	}

	//Copying
	if err := clipboard.WriteAll(text); err == nil {
		w.Write([]byte("Copied"))
	} else {
		w.Write([]byte("Could not Copy"))
	}
}
