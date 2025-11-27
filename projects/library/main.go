package main

import (
	"fmt"
	"library/utils"
	"net/http"
)

func main() {
	router := GetRouter()
	fmt.Printf("Server listening on port %d...", utils.PORT)
	port := fmt.Sprintf(":%d", utils.PORT)
	http.ListenAndServe(port, router)
}