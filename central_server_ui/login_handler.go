package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/2huskies/structs"
)

type LoginHandler struct{}

func (h *LoginHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	log.Printf("in LoginHandler")
	if req.Method != "POST" {
		http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	dec := json.NewDecoder(req.Body)
	uc := structs.UserCheck{}
	err := dec.Decode(&uc)
	if err != nil {
		http.Error(res, fmt.Sprintf("Invalid request: %s", err), http.StatusBadRequest)
		return
	}
	result, err := api.verify_user(uc.UserName, uc.Password)
	if err != nil {
		http.Error(res, fmt.Sprintf("Internal error: %s", err), http.StatusInternalServerError)
		return
	}
	if result == nil {
		http.Error(res, "Permission denied", http.StatusUnauthorized)
		return
	}
	json.NewEncoder(res).Encode(result)
}
