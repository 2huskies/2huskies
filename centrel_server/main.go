package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
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
var db *sql.DB

func getAbiturients(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(abiturients)
}

func getAbiturient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	id = id + 1

	//query := fmt.Sprintf("SELECT * FROM abiturient WHERE id = %d;", id)
	query := "SELECT * FROM abiturient"
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	defer rows.Close()

	abiturient := new(Abiturient)

	err = rows.Scan(&abiturient.ID,
		&abiturient.FirstName,
		&abiturient.LastName,
		&abiturient.BirthDate,
		&abiturient.BirthPlace,
		&abiturient.Address,
		&abiturient.PhoneNumber)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	json.NewEncoder(w).Encode(&abiturient)
}

func createAbiturient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var abiturient Abiturient
	_ = json.NewDecoder(r.Body).Decode(&abiturient)
	abiturient.ID = strconv.Itoa(rand.Intn(1000000))
	abiturients = append(abiturients, abiturient)
	json.NewEncoder(w).Encode(abiturients)
}

type config struct {
	Bind   string
	DBConn string
}

var conf = &config{
	Bind:   ":3081",
	DBConn: "postgres://postgres:admin@localhost/postgres?sslmode=verify-full",
}

var conffile string

func main() {
	flag.StringVar(&conffile, "c", "config.json", "`файл` конфигурации")
	flag.Parse()

	if _, err := os.Stat(conffile); err == nil {
		f, _ := os.Open(conffile)
		dec := json.NewDecoder(f)
		err = dec.Decode(conf)
		if err != nil {
			log.Fatalf("cannot parse config file '%s': %s", conffile, err)
		}
	}
	var err error

	connStr := conf.DBConn
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	abiturients = append(abiturients, Abiturient{ID: "1", FirstName: "Cris"})

	r := mux.NewRouter()
	r.HandleFunc("/abiturient", getAbiturients).Methods("GET")
	r.HandleFunc("/abiturient/{id}", getAbiturient).Methods("GET")
	r.HandleFunc("/books", createAbiturient).Methods("POST")

	log.Printf("listening: %s", conf.Bind)
	log.Fatal(http.ListenAndServe(conf.Bind, r))
}
