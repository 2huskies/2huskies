package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type SubjectsHandler struct{}

func (h *SubjectsHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	//	log.Printf("abiturient handler")
	if req.Method != "GET" {
		http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	result, err := api.getSubjects()
	if err != nil {
		http.Error(res, fmt.Sprintf("get subjects: %s", err), http.StatusBadRequest)
		return
	}
	json.NewEncoder(res).Encode(result)
}
