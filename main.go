package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"errors"

	"github.com/midepeter/kulipa/api"
	"github.com/midepeter/kulipa/store/db"
)

func main() {
	fmt.Println("This is a golang payment app using lightning network")
	err := db.Setup()
	if err != nil {
		errors.New("Unable set up db")
	}

	api.Run()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	sig := <-sigChan

	log.Println("Gracefully shutting down", sig)
}
