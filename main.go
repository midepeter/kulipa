package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/midepeter/kulipa/api"
	"github.com/midepeter/kulipa/store/database"
)

func main() {
	fmt.Println("This is a golang payment app using lightning network")
	database.InitDB()
	api.Run()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	sig := <-sigChan

	log.Println("Gracefully shutting down", sig)
}
