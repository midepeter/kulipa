package api

import (
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	config "github.com.spf13/viper"
	"github.com/urfave/negroni"
	"github.com/rs/cors"
)

//This represents a web API server
type Server struct {
	ou
}

func (s *Server) init() {
	router := mux.NewRouter().PathPrefix("/api").Subrouter().StrictSlash(true)
	router.HandleFunc("/", s.Hello).Methods("GET")
	router.HandleFunc("/proceed", s.Proceed).Methods("POST")
	router.HandleFunc("/pay", s.ProceedToPay).Methods("POST")

	n := negroni.New()
	n.UseFunc(s.ShutdownMiddleware)

	n.UseHandler(router)
	
	s.sender.HTTPRouter.PathPrefix("/api").Handler(n)
}