package main

import (
	"fmt"
	"library/db"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load environmental variables
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Println("No .env file found, using defaults")
	}

	dbPath := os.Getenv("DB_PATH")
	port := os.Getenv("PORT")

	// Initialise DB
	var repo *db.Repository = db.InitDB(dbPath)

	// Start up server
	router := GetRouter(repo)
	fmt.Printf("Server listening on port %s...", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), router)
}