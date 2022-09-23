package routes

import (
	"github.com/gorilla/mux"
	"github.com/nitintf/go-learning/projects/03-crud-mysql/pkg/controllers"
)

var RegisterBookRoutes = func(router *mux.Router) {
	router.HandleFunc("/book", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.DeleteBookById).Methods("DELETE")
	router.HandleFunc("/book/{id}", controllers.UpdateBookById).Methods("PUT")
}
