package db

import (
	"database/sql"

	"github.com/Rohan2998/the-dating-app-backend/config"
	_ "github.com/go-sql-driver/mysql"
)

// create a connection to database
func Connect(cfg *config.Config) (*sql.DB, error) {
	dsn := cfg.DBUser + ":" + cfg.DBPassword + "@tcp(" + cfg.DBHost + ":" + cfg.DBPort + ")/" + cfg.DBName
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
