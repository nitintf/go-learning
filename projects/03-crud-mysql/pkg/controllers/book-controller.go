package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nitintf/go-learning/projects/03-crud-mysql/pkg/models"
)

var NewBook models.Book

func CreateBook(w http.ResponseWriter, _ *http.Request) {
	newBook := models.GetAllBooks()
	data, _ := json.Marshal(newBook)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(data)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["bookId"]

	ID, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		panic(err)
	}

	bookDetails, _ := models.GetBookById(ID)

	data, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {}

func UpdateBookById(w http.ResponseWriter, r *http.Request) {}

func DeleteBookById(w http.ResponseWriter, r *http.Request) {}
