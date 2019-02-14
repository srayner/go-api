package main

import(
	"encoding/json"
	"log"
	"net/http"
	"math/rand"
	"strconv"
)

type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func getBooks(w http.ResponseWriter, r *http.Request) {

}

func createBook(w http.ResponseWriter, r *http.Request) {

}

func getBook(w http.ResponseWriter, r *http.Request) {

}

func updateBook(w http.ResponseWriter, r *http.Request) {

}

func deleteBook(w http.ResponseWriter, r *http.Request) {

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/books", getBooks).Method("GET")
	r.HandleFunc("/api/books", createBook).Method("POST")
	r.HandleFunc("/api/books/{id}", getBook).Method("GET")
	r.HandleFunc("/api/books/{id}", updateBook).Method("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Method("DELETE")
	log.Fatal(http.ListenAndServe(":80", r))
}
