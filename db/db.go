package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB(datasource string) {
	DB, err := sql.Open("mysql", datasource)

	if err != nil {
		panic(fmt.Sprintf("Error opening database: %v", err))
	}

	// Check if the database is actually reachable
	err = DB.Ping()
	if err != nil {
		panic(fmt.Sprintf("Error connecting to the database: %v", err))
	}

	DB.SetMaxOpenConns(10)
	DB.SetConnMaxIdleTime(2)

	createTables()
}

func createTables() {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTIME DATETIME NOT NULL,
		user_id INTEGER
	);`
	_, err := DB.Exec(createEventsTable)
	if err != nil {
		panic(fmt.Sprintf("Error creating tables: %v", err))
	}
}
