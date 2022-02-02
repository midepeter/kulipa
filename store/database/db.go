package database

import (
	"errors"
	"log"
	"time"

	"github.com/boltdb/bolt"
)

func InitDB() {
	db, err := bolt.Open("kulipa.db", 0600, &bolt.Options{
		Timeout: 10 * time.Second,
	})
	if err != nil {
		errors.New("Could not not open database file on disk")
		log.Fatal(err)
	}
	defer db.Close()
}
