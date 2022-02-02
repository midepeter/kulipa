package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/midepeter/kulipa/api/routes"
)

func Run() {
	s := routes.Server{}

	router := mux.NewRouter().PathPrefix("/api").Subrouter().StrictSlash(true)
	router.HandleFunc("/", s.Hello).Methods("GET")
	router.HandleFunc("/proceed", s.Proceed).Methods("POST")
	router.HandleFunc("/pay", s.ProceedToPay).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}
