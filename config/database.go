package config

import (
	"database/sql"
	"fmt"
)

// Database setup
var (
	Db *sql.DB
)


func InitDatabase() error {
	var err error
	// Initialize config
	Db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", dbUser, dbPass, dbHost, dbName))
	if err != nil {
		return err
	}
	if err = Db.Ping(); err != nil {
		return err
	}
	return nil
}
