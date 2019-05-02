// never underscore errors meaning _ = something is bad
// import go logger
// if err != nil log the err



package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"fmt"
	"github.com/gorilla/mux"
)

//model or struct
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// author struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastName"`
}

//Init Books var as a slice book struct
var books []Book

// get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {
	//set content type
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)

}
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// gets any params from the request
	params := mux.Vars(r)
	// loop through and find books the same id
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})

}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	// creates a mock ID
	book.ID = strconv.Itoa(rand.Intn(1000))
	books = append(books, book)

	json.NewEncoder(w).Encode(&Book{})
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			err := json.NewDecoder(r.Body).Decode(&book)
			fmt.Println(err)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	// make a new router
	r := mux.NewRouter()

	// Mock data
	books = append(books, Book{ID: "1", Isbn: "213", Title: "Lance`s Book", Author: &Author{Firstname: "Lance", Lastname: "Merrill"}})
	books = append(books, Book{ID: "2", Isbn: "124", Title: "Mack`s Book 2", Author: &Author{Firstname: "Mack", Lastname: "Merrill"}})

	// Route handlers // endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	//run server
	log.Fatal(http.ListenAndServe(":8000", r))
}
