package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite" // Import the driver
)

var DB *sql.DB

func InitDB() {
	var err error

	DB, err = sql.Open("sqlite", "./library.db")
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Unable to connect to database: ", err)
	}
	fmt.Println("Connected to Library Database!")

	_, err = DB.Exec(createTableQuery)
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}
	fmt.Println("Database initialized successfully.")
}