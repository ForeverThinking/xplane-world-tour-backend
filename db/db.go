package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error

	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not conect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createFlightsTable := `
	CREATE TABLE IF NOT EXISTS flights (
	flight_id INTEGER PRIMARY KEY AUTOINCREMENT,
	start_icao TEXT NOT NULL,
	end_icao TEXT NOT NULL
	)
	`

	 if _, err := DB.Exec(createFlightsTable); err != nil {
		panic("Cannot create flights table.")
	 }
}