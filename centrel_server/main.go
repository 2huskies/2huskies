package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Abiturient struct {
	ID          string `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	BirthDate   string `json:"birth_date"`
	BirthPlace  string `json:"birth_place"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
}

var abiturients []Abiturient

func getAbiturients(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(abiturients)
}

func getAbiturient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range abiturients {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Abiturient{})
}

func createAbiturient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var abiturient Abiturient
	_ = json.NewDecoder(r.Body).Decode(&abiturient)
	abiturient.ID = strconv.Itoa(rand.Intn(1000000))
	abiturients = append(abiturients, abiturient)
	json.NewEncoder(w).Encode(abiturients)
}

func main() {
	r := mux.NewRouter()
	//books = append(books, Book{ID: "1", Title: "Война и Мир", Author: &Author{Firstname: "Лев", Lastname: "Толстой"}})
	//books = append(books, Book{ID: "2", Title: "Преступление и наказание", Author: &Author{Firstname: "Фёдор", Lastname: "Достоевский"}})

	abiturients = append(abiturients, Abiturient{ID: "1", FirstName: "Cris"})

	r.HandleFunc("/abiturient", getAbiturients).Methods("GET")
	r.HandleFunc("/abiturient/{id}", getAbiturient).Methods("GET")
	r.HandleFunc("/books", createAbiturient).Methods("POST")

	log.Fatal(http.ListenAndServe(":3081", r))
}
