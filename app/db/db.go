package db

import (
	"database/sql"
	"fmt"

	"app/config"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func New(config *config.Config) (*sql.DB, error) {
	db, err := sql.Open("pgx", config.Db.URL)
	if err != nil {
		return nil, fmt.Errorf("[db] open: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("[db] ping: %w", err)
	}

	return db, nil
}
