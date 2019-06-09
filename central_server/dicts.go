package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/2huskies/2huskies/structs"
	"github.com/gorilla/mux"
)

func getSubjects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	query := "SELECT * FROM subject order by id"
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, fmt.Sprintf("query: %s: %s", http.StatusText(500), err), 500)
		return
	}
	defer rows.Close()

	subjects := make([]*structs.Subject, 0, 30)
	for rows.Next() {
		subj := &structs.Subject{}
		err = rows.Scan(
			&subj.ID,
			&subj.Name,
		)
		if err != nil {
			http.Error(w, fmt.Sprintf("scan: %s: %s", http.StatusText(500), err), 500)
			return
		}
		subjects = append(subjects, subj)
	}

	json.NewEncoder(w).Encode(subjects)
}

func getSpecialties(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	query := "SELECT * FROM specialty order by code"
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, fmt.Sprintf("query: %s: %s", http.StatusText(500), err), 500)
		return
	}
	defer rows.Close()

	specialties := make([]*structs.Specialty, 0, 30)
	for rows.Next() {
		spec := &structs.Specialty{}
		err = rows.Scan(
			&spec.Code,
			&spec.Name,
		)
		if err != nil {
			http.Error(w, fmt.Sprintf("scan: %s: %s", http.StatusText(500), err), 500)
			return
		}
		specialties = append(specialties, spec)
	}

	json.NewEncoder(w).Encode(specialties)
}

func getUniversities(w http.ResponseWriter, r *http.Request) {
	log.Printf("in getUniversities")
	w.Header().Set("Content-Type", "application/json")
	query := "SELECT * FROM university order by code"
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, fmt.Sprintf("query: %s: %s", http.StatusText(500), err), 500)
		return
	}
	defer rows.Close()

	arr := make([]*structs.University, 0, 30)
	for rows.Next() {
		item := &structs.University{}
		err = rows.Scan(
			&item.Code,
			&item.Name,
			&item.Active,
			&item.URL,
			&item.ShortName,
			&item.City,
			&item.Rate,
		)
		if err != nil {
			http.Error(w, fmt.Sprintf("scan: %s: %s", http.StatusText(500), err), 500)
			log.Printf("scan: %s: %s", http.StatusText(500), err)
			return
		}
		arr = append(arr, item)
	}

	log.Printf("%v", arr)
	json.NewEncoder(w).Encode(arr)
}

func getAbiturientScores(w http.ResponseWriter, r *http.Request) {
	log.Printf("in getAbiturientScores")
	w.Header().Set("Content-Type", "application/json")
	query := `
SELECT 
  j.id as subject_id,
  j.name as subject_name,
  s.score as score
FROM score s 
  join subject j on s.subject_id = j.id 
where abiturient_id = %d order by j.id
`
	params := mux.Vars(r)
	log.Printf("getAbiturientScores params: %v", params)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		http.Error(w, fmt.Sprintf("%s: %s", http.StatusText(500), err), 500)
		log.Printf("%s: %s", http.StatusText(500), err)
		return
	}

	rows, err := db.Query(fmt.Sprintf(query, id))
	if err != nil {
		http.Error(w, fmt.Sprintf("query: %s: %s", http.StatusText(500), err), 500)
		log.Printf("query %s: %s", http.StatusText(500), err)
		return
	}
	defer rows.Close()

	arr := make([]*structs.AbiturientScore, 0, 30)
	for rows.Next() {
		item := &structs.AbiturientScore{AbiturientID: id}
		err = rows.Scan(
			&item.SubjectID,
			&item.SubjectName,
			&item.Score,
		)
		if err != nil {
			http.Error(w, fmt.Sprintf("scan: %s: %s", http.StatusText(500), err), 500)
			log.Printf("scan: %s: %s", http.StatusText(500), err)
			return
		}
		arr = append(arr, item)
	}

	log.Printf("%v", arr)
	json.NewEncoder(w).Encode(arr)
}
