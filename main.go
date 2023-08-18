package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type book struct {
	Id       int
	Title    string
	Author   string
	Quantity int
}

var books = []book{
	{Id: 1, Title: "A strange Loop", Author: "Alex Bradman", Quantity: 5},
	{Id: 2, Title: "Atomic Habits", Author: "James Clear", Quantity: 7},
	{Id: 3, Title: "Homo Sapiens", Author: "Yuval Noah Harari", Quantity: 2},
}

// GET: list of books
func getBooks(w http.ResponseWriter, r *http.Request) {
	jsonResponse, err := json.Marshal(books)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err != nil {
		return
	}
	w.Write(jsonResponse)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", getBooks).Methods("GET")

	http.Handle("/", router)

	fmt.Println("Server Listening at http://127.0.0.1:8080/")
	http.ListenAndServe(":8080", nil)
}
