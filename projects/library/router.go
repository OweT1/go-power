package main

import (
	"library/api/middleware"
	"library/api/v1/routes"
	"library/db"
	"net/http"
)

func GetRouter(repo *db.Repository) (*http.ServeMux){
	mux := http.NewServeMux()

	authHandler := routes.AuthHandler{Repo: repo}
	bookHandler := routes.BookHandler{Repo: repo}

	// Public Routes
	mux.HandleFunc("POST /register", authHandler.Register)
	mux.HandleFunc("POST /login", authHandler.Login)
	mux.HandleFunc("GET /books", bookHandler.GetBooks)

	// Protected Routes
	mux.Handle("POST /books", middleware.AuthMiddleware(http.HandlerFunc(bookHandler.CreateBook)))
	mux.Handle("DELETE /books/{id}", middleware.AuthMiddleware(http.HandlerFunc(bookHandler.DeleteBook)))

	return mux
}