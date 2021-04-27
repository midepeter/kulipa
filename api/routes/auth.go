package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Register struct {
	ID               int
	Email            string
	AssociationName  string
	AssociationAlias string
	AssocationType   string
	Username         string
	Password         string
	ConfirmPassword  string
}

type Login struct {
	Email    string
	Password string
}

func Resgister(w http.ResponseWriter, r *http.Request) {
	var register Register
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(body, &register)
	if err != nil {
		fmt.Println("Unable to unmarshal error")
	}
	err = json.Marshal(register)
	if err != nil {
		panic(err)
	}

	token, err := Generatetoken(register["ID"])
	if err != nil {
		fmt.Println("Unable to generate token")
	}

	return token, nil
}
