package db

import (
	"database/sql"
	"fmt"
	"log"

	//_ "github.com/mattn/go-sqlite3"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB(datasource string) {
	var err error
	DB, err = sql.Open("mysql", datasource)

	if err != nil {
		panic(fmt.Sprintf("Error opening database: %v", err))
	}

	//Check if the database is actually reachable
	err = DB.Ping()
	if err != nil {
		panic(fmt.Sprintf("Error connecting to the database: %v", err))
	}

	DB.SetMaxOpenConns(10)
	DB.SetConnMaxIdleTime(5)

	createTables()
}

// createTables func creates the necessary tables
func createTables() {
	createEventsTable := `
    CREATE TABLE IF NOT EXISTS events (
        id INT AUTO_INCREMENT PRIMARY KEY,
        name TEXT NOT NULL,
        description TEXT NOT NULL,
        location TEXT NOT NULL,
        dateTime DATETIME NOT NULL,
        user_id INT
    );`

	_, err := DB.Exec(createEventsTable)
	if err != nil {
		panic(fmt.Sprintf("Error creating events table: %v", err))
	}

	log.Println("Events table created or already exists")
}
