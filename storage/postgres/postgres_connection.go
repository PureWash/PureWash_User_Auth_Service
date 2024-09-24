package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectToPostgres(postgresUrl string) (*sql.DB, error) {
	db, err := sql.Open("postgres", postgresUrl)

	if err != nil {
		return nil, fmt.Errorf("failed to open to postgres: %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping to postgres: %v", err)
	}

	return db, nil
}
