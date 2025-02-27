package Bookroutes

import (
	"book-management-system/pkg/controllers"
	"github.com/gorilla/mux"
)

var registerBookRoutes = func(router *mux.Router) {
	router.HandleFunc("/book", controllers.HandleBookCreate).Methods("POST")
	router.HandleFunc("/book", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
