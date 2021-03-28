package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/midepeter/gopay-app/api/routes"
)

//This represents a web API server
type Server struct {
	routers
}

func (s *Server) init() {
	router := mux.NewRouter().PathPrefix("/api").Subrouter().StrictSlash(true)
	router.HandleFunc("/", s.Hello).Methods("GET")
	router.HandleFunc("/proceed", s.Proceed).Methods("POST")
	router.HandleFunc("/pay", s.ProceedToPay).Methods("POST")

	
	log.Fatal(http.ListenAndServe(":8080", router))
}
