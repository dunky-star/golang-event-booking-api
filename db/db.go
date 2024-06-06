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
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INT PRIMARY KEY AUTO_INCREMENT,
		email VARCHAR(50) NOT NULL UNIQUE,
		password VARCHAR(255) NOT NULL
	)`

	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic(fmt.Sprintf("Error creating users table: %v", err))
	}

	createEventsTable := `
    CREATE TABLE IF NOT EXISTS events (
        id INT PRIMARY KEY AUTO_INCREMENT,
        name VARCHAR(50) NOT NULL,
        description VARCHAR(50) NOT NULL,
        location VARCHAR(50) NOT NULL,
        dateTime DATETIME NOT NULL,
        user_id INT,
		FOREIGN KEY (user_id) REFERENCES users(id)
    );`

	_, err = DB.Exec(createEventsTable)
	if err != nil {
		panic(fmt.Sprintf("Error creating events table: %v", err))
	}

	log.Println("Events table created or already exists")
}
