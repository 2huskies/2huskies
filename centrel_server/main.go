package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)


func main() {
	r := mux.NewRouter()
	//books = append(books, Book{ID: "1", Title: "Война и Мир", Author: &Author{Firstname: "Лев", Lastname: "Толстой"}})
	//books = append(books, Book{ID: "2", Title: "Преступление и наказание", Author: &Author{Firstname: "Фёдор", Lastname: "Достоевский"}})
	//r.HandleFunc("/books", getBooks).Methods("GET")
	//r.HandleFunc("/books/{id}", getBook).Methods("GET")
	//r.HandleFunc("/books", createBook).Methods("POST")
	//r.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	//r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":3081", r))
}
