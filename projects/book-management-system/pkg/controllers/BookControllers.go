package controllers

import (
	"book-management-system/pkg/models"
	// "book-management-system/pkg/utils"
	"encoding/json"
	// "fmt"
	// "github.com/gorilla/mux"
	"net/http"
	// "strconv"
)

var NewBook models.Book

func HandleBookCreate(w http.ResponseWriter, res *http.Request) {}

func GetBooks(w http.ResponseWriter, res *http.Request) {
	newBooks := models.GetAllBooks()
	book_res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(book_res)
}

func GetBookById(http.ResponseWriter, *http.Request) {}

func DeleteBook(http.ResponseWriter, *http.Request) {}

func UpdateBook(http.ResponseWriter, *http.Request) {}
