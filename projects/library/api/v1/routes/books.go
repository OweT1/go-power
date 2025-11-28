package routes

import (
	"encoding/json"
	"fmt"
	"library/db"
	"library/utils"
	"net/http"
)

type BookHandler struct {
	Repo *db.Repository
}

func (h *BookHandler) GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var books []utils.Book
	result := h.Repo.DB.Find(&books)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(books)
}

func (h *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Reads the book details into newBook
	var newBook utils.Book
	json.NewDecoder(r.Body).Decode(&newBook)

	// Adds the book into our Database
	result := h.Repo.DB.Create(&newBook)
	
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newBook)
}

func (h *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	bookId := r.PathValue("id")
	result := h.Repo.DB.Unscoped().Delete(&utils.Book{}, bookId)

	// Check if any of the results were deleted
	if result.RowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Book not found"})
		return
	}

	// Return Success / No change
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": fmt.Sprintf("Book(s) of ID = %s Deleted Successfully", bookId)})
}