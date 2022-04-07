package auth

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/midepeter/kulipa/store/db"
)

var (
	ErrAuthInvalid = errors.New("Invalid Authentication")
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	email := r.Header.Get("email")
	username := r.Header.Get("username")
	password := r.Header.Get("password")

	if email == "" || password == "" {
		fmt.Fprintf(w, "Email and Password required")
	}

	if len(password) < 8 {
		fmt.Fprintf(w, "Password not strong enough")
	}

	err := db.CreateUser(email, username, password)
	if err != nil {
		errors.New("Unable to store user in database")
	}
}

func Login(w http.ResponseWriter, r *http.Request) {

}

func Logout(w http.ResponseWriter, r *http.Request) {

}
