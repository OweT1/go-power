package main

import (
	"library/api/v1/routes"
	"net/http"
)

func GetRouter() (*http.ServeMux){
	mux := http.NewServeMux()

	mux.HandleFunc("GET /books", routes.GetBooks)
	mux.HandleFunc("POST /books", routes.CreateBook)
	mux.HandleFunc("DELETE /books/{id}", routes.DeleteBook)

	return mux
}