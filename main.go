package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/midepeter/kulipa/api"
	"github.com/midepeter/kulipa/pkg/logger"
	"github.com/midepeter/kulipa/store/db"
)

func main() {
	logger.Setuplogger()

	err := db.Setup()

	if err != nil {
		_ = fmt.Errorf("Unable set up db %v\n", err)
	}

	api.Run()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	sig := <-sigChan

	log.Println("Gracefully shutting down", sig)
}
