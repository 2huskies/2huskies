package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type UniversitiesHandler struct{}

func (h *UniversitiesHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	//	log.Printf("abiturient handler")
	if req.Method != "GET" {
		http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	result, err := api.getUniversities()
	if err != nil {
		http.Error(res, fmt.Sprintf("get universities: %s", err), http.StatusBadRequest)
		return
	}
	json.NewEncoder(res).Encode(result)
}
