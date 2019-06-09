package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/2huskies/structs"
	uuid "github.com/satori/go.uuid"
)

type LoginHandler struct{}

func (h *LoginHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	//log.Printf("in LoginHandler")
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
	id, _ := uuid.NewV4()
	sessionToken := id.String()
	_, err = cache.Do("SETEX", sessionToken+"-login", "300", uc.UserName)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = cache.Do("SETEX", sessionToken+"-abiturient-id", "300", fmt.Sprintf("%d", result.AbiturientID))
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(res, &http.Cookie{
		Name:    "session_token",
		Value:   sessionToken,
		Expires: time.Now().Add(120 * time.Second),
	})

	json.NewEncoder(res).Encode(result)
}
