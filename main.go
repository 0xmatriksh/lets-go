package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"lets-go/db"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func handleRequest(DB *sql.DB) {
	h := New(DB)
	router := mux.NewRouter()
	router.HandleFunc("/books", h.getAllBooks).Methods("GET")
	router.HandleFunc("/books/{id}", h.getBook).Methods("GET")
	router.HandleFunc("/books", h.addBook).Methods("POST")
	router.HandleFunc("/books/{id}", h.updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", h.deleteBook).Methods("DELETE")

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
