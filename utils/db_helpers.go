package utils

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/josep/mockha/constants"
	"log"
)

func InitDatabase(db *sqlx.DB) error {
	_, err := db.Exec(constants.Schema)
	if err != nil {
		log.Println(fmt.Errorf("failed to execute initial schema: %w", err))
	}

	if db.DriverName() == "sqlite3" {
		_, err = db.Exec(`PRAGMA foreign_keys = ON`)
		if err != nil {
			return fmt.Errorf("failed to enable foreign keys: %w", err)
		}
	}

	return nil
}
