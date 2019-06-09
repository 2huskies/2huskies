package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type AbiturientScoresHandler struct{}

func (h *AbiturientScoresHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
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
		http.Error(res, "Unauthorized", http.StatusUnauthorized)
		return
	}

	result, err := api.getAbiturientScores(s.getAbiturientID())
	if err != nil {
		http.Error(res, fmt.Sprintf("get abiturient scores: %s", err), http.StatusBadRequest)
		return
	}
	json.NewEncoder(res).Encode(result)
}
