package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"log"
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
	MiddleName  string `json:"middle_name"`
}

type User struct {
	Role   string `json:"role"`
	AbiturientID int64 `json:"abiturient_id"`
}

var db *sql.DB

func getAbiturients(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	query := "SELECT * FROM abiturient;"
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	defer rows.Close()

	abiturients := make([]*Abiturient, 0)
	for rows.Next() {
		abiturient := new(Abiturient)
		err = rows.Scan(&abiturient.ID,
			&abiturient.FirstName,
			&abiturient.LastName,
			&abiturient.BirthDate,
			&abiturient.BirthPlace,
			&abiturient.Address,
			&abiturient.PhoneNumber,
			&abiturient.MiddleName,
		)
		if err != nil {
			log.Fatal(err)
		}
		abiturients = append(abiturients, abiturient)
	}

	json.NewEncoder(w).Encode(abiturients)
}

func getAbiturient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		http.Error(w, fmt.Sprintf("%s: %s", http.StatusText(500), err), 500)
		return
	}

	query := fmt.Sprintf("SELECT * FROM abiturient WHERE id = %d", id)
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, fmt.Sprintf("query: %s: %s", http.StatusText(500), err), 500)
		return
	}
	defer rows.Close()

	abiturient := new(Abiturient)

	rows.Next()
	err = rows.Scan(&abiturient.ID,
		&abiturient.FirstName,
		&abiturient.LastName,
		&abiturient.BirthDate,
		&abiturient.BirthPlace,
		&abiturient.Address,
		&abiturient.PhoneNumber,
		&abiturient.MiddleName,
	)

	if err != nil {
		http.Error(w, fmt.Sprintf("scan: %s: %s", http.StatusText(500), err), 500)
		return
	}

	json.NewEncoder(w).Encode(&abiturient)
}

func verifyUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	login, err := strconv.ParseInt(params["user"], 10, 64)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	password, err := strconv.ParseInt(params["password"], 10, 64)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	query := fmt.Sprint("SELECT in_role, abiturient_id FROM login WHERE login = %s AND password = %s", login, password)
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	defer rows.Close()

	user := new(User)
	rows.Next()
	err = rows.Scan(&user.Role, &user.AbiturientID)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	json.NewEncoder(w).Encode(&user)
}

type config struct {
	Bind   string
	DBConn string
}

var conf = &config{
	Bind:   ":3081",
	DBConn: "postgres://postgres:admin@localhost/postgres?sslmode=disable",
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

	r := mux.NewRouter()
	r.HandleFunc("/abiturient", getAbiturients).Methods("GET")
	r.HandleFunc("/abiturient/{id}", getAbiturient).Methods("GET")
	//r.HandleFunc("/books", createAbiturient).Methods("POST")

	log.Printf("listening: %s", conf.Bind)
	log.Printf("db: %s", conf.DBConn)
	log.Fatal(http.ListenAndServe(conf.Bind, r))
}
