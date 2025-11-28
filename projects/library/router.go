package main

import (
	"library/api/v1/routes"
	"library/db"
	"net/http"
)

func GetRouter(repo *db.Repository) (*http.ServeMux){
	mux := http.NewServeMux()

	bookHandler := routes.BookHandler{
		Repo: repo,
	}
	mux.HandleFunc("GET /books", bookHandler.GetBooks)
	mux.HandleFunc("POST /books", bookHandler.CreateBook)
	mux.HandleFunc("DELETE /books/{id}", bookHandler.DeleteBook)

	return mux
}