package main

import (
	"fmt"
	"net/http"
	"strconv"

	"lets-go/util"

	"github.com/gorilla/mux"
)

// GET: list of books
func getAllBooks(w http.ResponseWriter, r *http.Request) {
	util.CreateResponse(w, books, http.StatusOK)
}

// GET: a book by Id
func getBook(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	for _, book := range books {
		if book.Id == id {
			util.CreateResponse(w, book, http.StatusOK)
			return
		}
	}
	util.CreateResponse(w, "", http.StatusNoContent)

}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/books", getAllBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")

	http.Handle("/", router)

	fmt.Println("Server Listening at http://127.0.0.1:8080/")
	http.ListenAndServe(":8080", nil)
}
