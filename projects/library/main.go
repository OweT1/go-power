package main

import (
	"fmt"
	"library/db"
	"library/utils"
	"net/http"
)

func main() {
	// Initialise DB
	db.InitDB()

	// Start up server
	router := GetRouter()
	fmt.Printf("Server listening on port %d...", utils.PORT)
	port := fmt.Sprintf(":%d", utils.PORT)
	http.ListenAndServe(port, router)
}