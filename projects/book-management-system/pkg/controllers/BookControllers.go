package controllers

import (
	"book-management-system/pkg/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

func HandleBookCreate(w http.ResponseWriter, req *http.Request) {
	var newBook models.Book
	err := json.NewDecoder(req.Body).Decode(&newBook)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	createdBook := newBook.CreateBook()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdBook)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	fmt.Printf("Fetched books: %+v\n", newBooks)

	book_res, err := json.Marshal(newBooks)
	if err != nil {
		http.Error(w, "Failed to encode books", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(book_res)
}

func GetBookById(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	bookId := vars["bookId"]

	// Convert bookId from string to int64
	id, err := strconv.ParseInt(bookId, 10, 64)
	if err != nil {
		http.Error(w, `{"error": "Invalid book ID"}`, http.StatusBadRequest)
		return
	}

	// Fetch book from database
	bookdetails, err := models.GetBookById(id)
	if err != nil {
		fmt.Println("Book not found in database") // Debug log
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Book not found"})
		return
	}

	// If book is found, return it
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bookdetails)
}

func DeleteBook(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	deleted_book := models.DeleteBook(id)
	res, _ := json.Marshal(deleted_book)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// ✅ UpdateBook: First delete the book, then create a new one
func UpdateBook(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	// Delete the existing book
	deletedBook := models.DeleteBook(id)
	fmt.Println("Deleted book:", deletedBook)

	// Decode new book details
	var newBook models.Book
	err = json.NewDecoder(req.Body).Decode(&newBook)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Create a new book
	createdBook := newBook.CreateBook()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(createdBook)
}
