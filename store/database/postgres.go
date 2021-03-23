package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type Config struct {
	Host     string
	Port     string
	DBName   string
	Password string
	User     string
	SSLMode  string
}

type PgClient struct {
	db  *sql.DB
	ctx context.Context
}

func NewPgClient(ctx context.Context, config *Config, wait bool) (*PgClient, error) {
	fmt.Println("Starting the database connection....")

	dsn := fmt.Sprintf("")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		fmt.Println(err)
	}

	//ping the database
	err = db.Ping()
	if err != nil {
		if !wait {
			return nil, err
		}

		log.Println(err)

		//wait till the database comes up
		itr := 0
	db_wait:
		for {
			itr++
			log.Printf("waiting for the database connection..... [%d]", itr)
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			case <-time.After(1 * time.Second):
				err = db.Ping()
				if err != nil {
					break db_wait
				}
			}
		}
	}

	return &PgClient{
		db:  db,
		ctx: ctx,
	}, nil
}
