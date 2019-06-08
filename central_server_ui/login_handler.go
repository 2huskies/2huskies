package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LoginHandler struct{}

func (h *LoginHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	dec := json.NewDecoder(req.Body)
	uc := UserCheck{}
	err := dec.Decode(&uc)
	if err != nil {
		http.Error(res, "Invalid request", http.StatusBadRequest)
		return
	}
	valid, err := api.check_user(uc.UserName, uc.Password)
	if err != nil {
		http.Error(res, fmt.Sprintf("Internal error: %s", err), http.StatusInternalServerError)
		return
	}
	if !valid {
		http.Error(res, "Permission denied", http.StatusUnauthorized)
		return
	}
}
