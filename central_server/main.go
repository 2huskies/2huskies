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

	"github.com/2huskies/structs"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

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

	abiturients := make([]*structs.Abiturient, 0)
	for rows.Next() {
		abiturient := new(structs.Abiturient)
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
	log.Printf("getAbiturient params: %v", params)
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

	abiturient := new(structs.Abiturient)

	if !rows.Next() {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
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
	dec := json.NewDecoder(r.Body)
	uc := &structs.UserCheck{}
	err := dec.Decode(uc)
	if err != nil {
		http.Error(w, fmt.Sprintf("%s: %s", http.StatusText(500), err), 500)
		return
	}

	log.Printf("verifyUser data: %v", uc)

	query := fmt.Sprintf("SELECT role, abiturient_id FROM login WHERE login = '%s' AND password = '%s'", uc.UserName, uc.Password)
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, fmt.Sprintf("query %s: %s", http.StatusText(500), err), 500)
		log.Printf("query: %s", err)
		return
	}
	defer rows.Close()

	user := new(structs.UserCheckResult)
	if !rows.Next() {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	err = rows.Scan(&user.Role, &user.AbiturientID)
	if err != nil {
		http.Error(w, fmt.Sprintf("%s: %s", http.StatusText(500), err), 500)
		log.Printf("%s", err)
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
	r.HandleFunc("/verify_user", verifyUser).Methods("POST")
	r.HandleFunc("/subjects", getSubjects).Methods("GET")
	r.HandleFunc("/specialties", getSpecialties).Methods("GET")
	r.HandleFunc("/universities", getUniversities).Methods("GET")
	r.HandleFunc("/abiturient_scores/{id}", getAbiturientScores).Methods("GET")
	//r.HandleFunc("/books", createAbiturient).Methods("POST")

	log.Printf("listening: %s", conf.Bind)
	log.Printf("db: %s", conf.DBConn)
	log.Fatal(http.ListenAndServe(conf.Bind, r))
}
