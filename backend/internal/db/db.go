package db

import (
	"database/sql"
	"fmt"
	"task-manager/internal/config"

	_ "github.com/lib/pq"
)

func Connect(cfg *config.Config) (*sql.DB, error) {
	connStr := cfg.DBUrl

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("Error connect: %s", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("DB unavailable: %s", err)
	}
	return db, nil
}
