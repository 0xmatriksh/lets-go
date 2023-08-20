package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"lets-go/db"
	model "lets-go/models"
	"lets-go/util"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type handler struct {
	DB *sql.DB
}

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

func handleRequest(DB *sql.DB) {
	h := New(DB)
	router := mux.NewRouter()
	router.HandleFunc("/books", h.getAllBooks).Methods("GET")
	router.HandleFunc("/books/{id}", h.getBook).Methods("GET")

	http.Handle("/", router)

	fmt.Println("Server Listening at http://127.0.0.1:8080/")
	http.ListenAndServe(":8080", nil)
}

func main() {
	fmt.Println("Started--")
	var myDB *sql.DB = db.DBConnect()
	db.InitializeTable(myDB)
	handleRequest(myDB)
	db.DBCloseConnection(myDB)
}
