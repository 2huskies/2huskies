package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type AbiturientHandler struct{}

func (h *AbiturientHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	//	log.Printf("abiturient handler")
	if req.Method != "GET" {
		http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	s, err := getSession(req)
	if err != nil {
		http.Error(res, "Bad request", http.StatusBadRequest)
		return
	}
	if s == nil {
		http.Error(res, "Unauthirized", http.StatusUnauthorized)
		return
	}

	result, err := api.getAbiturient(s.getAbiturientID())
	if err != nil {
		http.Error(res, fmt.Sprintf("get abiturient: %s", err), http.StatusBadRequest)
		return
	}
	json.NewEncoder(res).Encode(result)
}
