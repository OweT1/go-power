package routes

import (
	"encoding/json"
	"library/utils"
	"net/http"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(utils.LIBRARY)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newBook utils.Book
	json.NewDecoder(r.Body).Decode(&newBook)
	utils.LIBRARY = append(utils.LIBRARY, newBook)
	json.NewEncoder(w).Encode(newBook)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bookId := r.PathValue("id")
	newLibrary := make([]utils.Book, 0, len(utils.LIBRARY))

	found := false
	for _, book := range utils.LIBRARY {
		if book.ID == bookId {
			found = true
			continue
		}
		newLibrary = append(newLibrary, book)
	}

	// Return 404 Book not found
	if !found {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Book not found"})
		return
	}
	// Update Library
	utils.LIBRARY = newLibrary

	// Return Success / No change
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(utils.LIBRARY)
}