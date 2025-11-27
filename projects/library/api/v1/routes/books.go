package routes

import (
	"encoding/json"
	"fmt"
	"library/db"
	"library/utils"
	"net/http"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	rows, err := db.DB.Query(db.GetBooksQuery)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer rows.Close() // close the rows on function end

	var books []utils.Book

	for rows.Next() {
		var b utils.Book
		err = rows.Scan(&b.ID, &b.Title, &b.Author)
		if err != nil {
			http.Error(w, "Error scanning book", http.StatusInternalServerError)
			return
		}
		books = append(books, b)
	}
	json.NewEncoder(w).Encode(books)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Reads the book details into newBook
	var newBook utils.Book
	json.NewDecoder(r.Body).Decode(&newBook)

	// Adds the book into our Database
	_, err := db.DB.Exec(db.AddBookQuery, newBook.ID, newBook.Title, newBook.Author)
	
	if err != nil {
		http.Error(w, "Failed to insert book: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newBook)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	bookId := r.PathValue("id")
	result, err := db.DB.Exec(db.DeleteBookQuery, bookId)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	// Check if any of the results were deleted
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Book not found"})
		return
	}

	// Return Success / No change
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": fmt.Sprintf("Book(s) of ID = %s Deleted Successfully", bookId)})
}