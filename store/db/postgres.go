package db

import (
	"context"
	"errors"
	"os"

	"github.com/jackc/pgx/v4"
)

var (
	ErrDBQueryFailed = errors.New("DB Query Failed")
)

type DB struct {
	*pgx.Conn
}

func SetupDB() (*pgx.Conn, error) {
	dbURL := os.Getenv("DB_URL")

	ctx := context.Background()

	conn, err := pgx.Connect(ctx, dbURL)
	if err != nil {
		return nil, errors.New("Unable to connect to database!")
	}

	return conn, nil
}

//CreateUser stores a newly Registered to the database
func (db *DB) CreateUser(email, username, passwordHash string) error {
	sqlStmt := `INSERT INTO users (email, username, password)`
	_, err := db.Exec(context.Background(), sqlStmt, email, username, passwordHash)
	if err != nil {
		return ErrDBQueryFailed
	}

	return nil
}
