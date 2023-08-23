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
func New(db *sql.DB) *handler {
	return &handler{db}
}

// GET: list of books
func (h *handler) getAllBooks(w http.ResponseWriter, r *http.Request) {
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
func (h *handler) addBook(w http.ResponseWriter, r *http.Request) {
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

func _getBookHelper(h *handler, id string, w http.ResponseWriter) {
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
		util.CreateResponse(w, bookData, http.StatusOK)
		return
	}
	util.CreateResponse(w, "", http.StatusNoContent)
}

// GET: a book by Id
func (h *handler) getBook(w http.ResponseWriter, r *http.Request) {
	id, _ := mux.Vars(r)["id"]

	_getBookHelper(h, id, w)
}

// PUT: a book by Id
func (h *handler) updateBook(w http.ResponseWriter, r *http.Request) {
	id, _ := mux.Vars(r)["id"]

	var book model.Book
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&book)

	updateBookQuery := "UPDATE books SET title=$1, author=$2, quantity=$3 WHERE id = $4;"
	_, err = h.DB.Exec(updateBookQuery, book.Author, book.Title, book.Quantity, id)
	if err != nil {
		fmt.Println("failed to execute UPDATE query", err)
		w.WriteHeader(500)
		return
	}

	_getBookHelper(h, id, w)
}

// DELETE: a book by Id
func (h *handler) deleteBook(w http.ResponseWriter, r *http.Request) {
	id, _ := mux.Vars(r)["id"]

	deleteQuery := "DELETE FROM books WHERE id = $1"
	_, err := h.DB.Exec(deleteQuery, id)
	if err != nil {
		fmt.Println("failed to execute DELETE query", err)
		w.WriteHeader(500)
		return
	}
	jsonData := map[string]string{
		"id":      id,
		"message": "Book deleted successfully",
	}
	util.CreateResponse(w, jsonData, http.StatusOK)
}
