package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	model "lets-go/models"
	"lets-go/util"
	"net/http"

	"github.com/gorilla/mux"
)

type handler struct {
	DB *sql.DB
}

// like constructor for handler
func New(db *sql.DB) handler {
	return handler{db}
}

// GET: list of books
func (h handler) getAllBooks(w http.ResponseWriter, r *http.Request) {
	results, err := h.DB.Query("SELECT * FROM books;")
	if err != nil {
		fmt.Println("failed to execute query", err)
		return
	}
	var allBooks = make([]model.Book, 0)
	for results.Next() {
		var bookData model.Book
		err = results.Scan(&bookData.Id, &bookData.Title, &bookData.Author, &bookData.Quantity)

		allBooks = append(allBooks, bookData)
	}

	util.CreateResponse(w, allBooks, http.StatusOK)
}

// POST: add new Book
func (h handler) addBook(w http.ResponseWriter, r *http.Request) {
	// TODO: validate the request body data
	var book model.Book
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&book)

	insertBookQuery := "INSERT INTO books (id,title,author,quantity) VALUES ($1, $2, $3, $4);"
	_, err = h.DB.Exec(insertBookQuery, book.Id, book.Title, book.Author, book.Quantity)
	if err != nil {
		w.WriteHeader(500)
		fmt.Println("failed to execute query", err)
		return
	}

	util.CreateResponse(w, book, http.StatusCreated)
}

// GET: a book by Id
func (h handler) getBook(w http.ResponseWriter, r *http.Request) {
	id, _ := mux.Vars(r)["id"]
	fmt.Println(id)
	results, err := h.DB.Query(`SELECT * FROM books WHERE id = $1 ;`, id)
	if err != nil {
		fmt.Println("failed to execute SELECT query", err)
		w.WriteHeader(500)
		return
	}
	var bookData model.Book
	for results.Next() {
		err = results.Scan(&bookData.Id, &bookData.Title, &bookData.Author, &bookData.Quantity)
		if err != nil {
			fmt.Println("failed in reading results", err)
			w.WriteHeader(500)
			return
		}
		util.CreateResponse(w, bookData, http.StatusNoContent)
		return
	}
	util.CreateResponse(w, "", http.StatusNoContent)
}
